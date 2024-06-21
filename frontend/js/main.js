// frontend/js/main.js

import { initWebSocket } from './websocket.js';
import { render } from './navigation.js';

document.addEventListener("DOMContentLoaded", () => {
    initWebSocket();
    render('home');  // Default page to render

    document.addEventListener('click', (e) => {
        if (e.target.matches('[data-nav]')) {
            e.preventDefault();
            const page = e.target.getAttribute('data-nav');
            console.log(`Navigating to ${page}`);
            navigate(page);
        }
    });

    window.addEventListener('popstate', () => {
        render('home');  // Always render home on popstate to keep URL clean
    });
});

window.navigate = function(page) {
    history.pushState({}, '', '/');
    console.log(`Navigating to ${page}`);
    render(page);
};
