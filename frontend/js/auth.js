import { login, register } from './api.js';
import { render } from './navigation.js';

export function errorLog() {
    const form = document.getElementById('login-form');
    if (!form) {
        console.error('Login form not found.');
        return; // Exit the function if the form is not found
    }

    let errorMessage = document.getElementById('login-error');
    if (!errorMessage) {
        errorMessage = document.createElement('div');
        errorMessage.setAttribute('id', 'login-error');
        errorMessage.classList.add('error-message'); // Use a class for styling
        errorMessage.setAttribute('role', 'alert'); // Improve accessibility
        errorMessage.style.color = 'red'; // Set the error message color to red
        errorMessage.textContent = "An error occurred. Please try again."; // Set the error message text

        const icon = document.createElement('i');
        icon.classList.add('fas', 'fa-exclamation-triangle');
        errorMessage.appendChild(icon);
        
        form.appendChild(errorMessage);
    }

    errorMessage.style.display = 'flex'; // Ensure it's visible
}


export function errorRegister() {
    const form = document.getElementById('register-form');
    form.addEventListener('submit', function(event) {
        event.preventDefault();
        const data = new FormData(form);
        register(data)
            .then(response => {
                if (response.ok) {
                    render('login');
                } else {
                    response.json().then(errorData => {
                        // Assuming errorData contains a message property with the error description
                        let errorMessage = document.getElementById('register-error');
                        if (!errorMessage) {
                            errorMessage = document.createElement('div');
                            errorMessage.setAttribute('id', 'register-error');
                            errorMessage.classList.add('error-message'); // Use a class for styling
                            errorMessage.setAttribute('role', 'alert'); // Improve accessibility
                            form.appendChild(errorMessage);
                        }
                        // Update the text content of the error message based on the error received
                        errorMessage.style.display = 'flex';
                        errorMessage.textContent = errorData.message || "An error occurred during registration. Please try again.";
                    }).catch(() => {
                        // Fallback error message if the response cannot be parsed
                        let errorMessage = document.getElementById('register-error');
                        if (!errorMessage) {
                            errorMessage = document.createElement('div');
                            errorMessage.setAttribute('id', 'register-error');
                            errorMessage.classList.add('error-message'); // Use a class for styling
                            errorMessage.setAttribute('role', 'alert'); // Improve accessibility
                            form.appendChild(errorMessage);
                        }
                        errorMessage.style.display = 'flex';
                        errorMessage.textContent = "An error occurred, but we couldn't get the details. Please try again.";
                    });
                }
            });
    });
}

export function setupLogin() {
    const form = document.getElementById('login-form');
    form.addEventListener('submit', function(event) {
        event.preventDefault();
        const data = new FormData(form);
        login(data)
            .then(response => {
                if (response.ok) {
                    render('home');
                } else {
                    errorLog();
                }
            });
    });
}

export function setupRegister() {
    const form = document.getElementById('register-form');
    form.addEventListener('submit', function(event) {
        event.preventDefault();
        const data = new FormData(form);
        register(data)
            .then(response => {
                if (response.ok) {
                    render('login');
                } else {
                    errorRegister();
                }
            });
    });
}
