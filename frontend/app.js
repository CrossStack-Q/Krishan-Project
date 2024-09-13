


document.addEventListener('DOMContentLoaded', () => {
    // Check if the elements exist before assigning them
    const loginButton = document.getElementById('login-btn');
    const userMenu = document.getElementById('user-menu');
    const userGreeting = document.getElementById('user-greeting');
    const logoutButton = document.getElementById('logout-btn');
    const signupForm = document.getElementById('signup-form');

    // Handle user session and update UI
    const user = JSON.parse(localStorage.getItem('user'));

    if (user) {
        if (userGreeting) {
            userGreeting.textContent = `Hi User`;
        }
        if (userMenu) {
            userMenu.classList.remove('hidden');
            if (loginButton) {
                loginButton.classList.add('hidden');
            }
        }
    } else {
        if (userMenu) {
            userMenu.classList.add('hidden');
        }
    }

    // Handle login form submission
    document.getElementById('login-form')?.addEventListener('submit', async (event) => {
        event.preventDefault();
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;

        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ email, password })
            });

            if (response.ok) {
                localStorage.setItem('user', JSON.stringify({ email }));
                window.location.href = '/';
            } else {
                throw new Error('Login failed');
            }
        } catch (error) {
            alert('An error occurred. Please try again.');
        }
    });

    // Handle signup form submission
    if (signupForm) {
        signupForm.addEventListener('submit', async (event) => {
            event.preventDefault();
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;

            try {
                const response = await fetch('http://localhost:8080/signup', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ email, password })
                });

                if (response.ok) {
                    localStorage.setItem('user', JSON.stringify({ email }));
                    window.location.href = '/';
                } else {
                    throw new Error('Signup failed');
                }
            } catch (error) {
                alert('An error occurred. Please try again.');
            }
        });
    }

    // Handle logout button click
    if (logoutButton) {
        logoutButton.addEventListener('click', () => {
            fetch('http://localhost:8080/logout', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
            })
                .then(response => {
                    if (response.ok) {
                        localStorage.removeItem('user');
                        window.location.href = '/';
                    } else {
                        alert('Error logging out. Please try again.');
                    }
                })
                .catch(error => {
                    console.error('Error:', error);
                    alert('Error logging out. Please try again.');
                });
        });
    }


    // data 

    function loadHotels(location) {
        fetch(`http://localhost:8080/hotels?location=${location}`)
            .then(response => response.json())
            .then(data => {
                const hotelsContainer = document.getElementById('hotels-container');
                hotelsContainer.innerHTML = ''; // Clear any existing content
    
                data.forEach(hotel => {
                    const hotelDiv = `
                    <div class="bg-gray-100 p-6 rounded-lg hover:cursor-pointer" onclick="window.location.href='/hotel.html?hotelID=${hotel.hotel_id}';">
                        <img class="h-40 rounded w-full object-cover object-center mb-6" src="${hotel.images[0]}" alt="hotel image">
                        <div class="flex justify-between">
                            <h3 class="tracking-widest text-indigo-500 text-xs font-medium title-font">
                                $${hotel.price}
                            </h3>
                            <h3 class="tracking-widest text-indigo-500 text-xs font-medium title-font">
                                ${hotel.rating} ★
                            </h3>
                        </div>
                        <h2 class="text-lg text-gray-900 font-medium title-font mb-4">${hotel.name}</h2>
                        <p class="leading-relaxed text-base">${hotel.landmark}</p>
                    </div>
                `;
                    hotelsContainer.innerHTML += hotelDiv;
                });
            })
            .catch(error => {
                console.error('Error fetching hotel data:', error);
            });
    }
    
    loadHotels('gaza');


    
    




});
