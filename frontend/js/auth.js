// real-time-forum/frontend/js/auth.js

import { login, register } from './api.js';
import { render } from './navigation.js';

export function errorLog() {
    const form = document.getElementById('login-form');
    let errorMessage = document.getElementById('error-message');
    if (!errorMessage) {
        errorMessage = document.createElement('div');
        errorMessage.setAttribute('id', 'error-message');
        errorMessage.style.color = 'red'; // Style the error message in red
        form.appendChild(errorMessage); // Append the error message element to the form
    }
    // Update the text content of the error message
    errorMessage.textContent = "Check your email or password, please try again.";
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
                        let errorMessage = document.getElementById('error-message');
                        if (!errorMessage) {
                            errorMessage = document.createElement('div');
                            errorMessage.setAttribute('id', 'error-message');
                            errorMessage.style.color = 'red';
                            form.insertBefore(errorMessage, form.firstChild); // Insert before the form elements
                        }
                        // Update the text content of the error message based on the error received
                        errorMessage.textContent = errorData.message || "An error occurred during registration. Please try again.";
                    }).catch(() => {
                        // Fallback error message if the response cannot be parsed
                        let errorMessage = document.getElementById('error-message');
                        if (!errorMessage) {
                            errorMessage = document.createElement('div');
                            errorMessage.setAttribute('id', 'error-message');
                            errorMessage.style.color = 'red';
                            form.insertBefore(errorMessage, form.firstChild); // Insert before the form elements
                        }
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
                    // Handle error
                    //errorLog();
                }
            });
    });
}
