// real-time-forum/frontend/js/auth.js

import { login, register } from './api.js';
import { render } from './navigation.js';

export function setupLogin() {
    const form = document.getElementById('login-form');
    form.addEventListener('submit', function(event) {
        event.preventDefault();
        const data = new FormData(form);
        login(data)
            .then(response => {
                if (response.ok) {
                    sessionStorage.setItem('loggedIn', 'true');
                    render('home');
                } else {
                    // Handle error
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
                    // Handle error
                }
            });
    });
}
