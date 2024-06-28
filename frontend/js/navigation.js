// frontend/js/navigation.js

import { fetchPosts } from './api.js';
import { setupLogin, setupRegister } from './auth.js';
import { setupPost, displayPosts } from './post.js';

export function render(page) {
    console.log(`Rendering ${page}`);
    loadPageContent(page);
}

function checkCookie(cookieName) {
    // Split document.cookie at '; ' to get an array of cookies
    const cookies = document.cookie.split('; ');
    // Find the specific cookie by name
    const targetCookie = cookies.find(cookie => cookie.startsWith(cookieName + '='));
    return targetCookie ? true : false;
}

function loadPageContent(page) {
    const app = document.getElementById('app');
    const isLoggedIn = checkCookie(`session-name`);
    
    console.log('esli true to zalogineni:', isLoggedIn);

    if (page === 'home') {
        if (isLoggedIn) {
            fetchPosts()
                .then(posts => {
                    app.innerHTML = `
                        <h1>Welcome to the Real-Time Forum</h1>
                        <p>This is the home page.</p>
                        <a href="#" data-nav="post">Create Post</a> | <a href="#" data-nav="logout">Logout</a>
                                
                        <div id="posts"></div>
                    `;
                    displayPosts(posts);
                })
                .catch(error => {
                    console.error('Error fetching posts:', error);
                    app.innerHTML = '<p>Error loading posts. Please try again later.</p>';
                });
        } else {
            app.innerHTML = `
                <h1>Welcome to the Real-Time Forum</h1>
                <p>This is the home page.</p>
                <a href="#" data-nav="login">Login</a> | <a href="#" data-nav="register">Register</a>
            `;
        }
    } else if (page === 'login') {
        app.innerHTML = `
            <h1>Login</h1>
            <form id="login-form">
                <input type="email" name="email" placeholder="Email" required><br>
                <input type="password" name="password" placeholder="Password" required><br>
                <button type="submit">Login</button></br></br>
                <button data-nav="home">Home</button>
            </form>
            
            
        `;
        setupLogin();
    } else if (page === 'register') {
        app.innerHTML = `
            <h1>Register</h1>
            <form id="register-form">
                <input type="text" name="nickname" placeholder="Nickname" required><br>
                <input type="number" name="age" placeholder="Age" required><br>
                <input type="text" name="gender" placeholder="Gender" required><br>
                <input type="text" name="firstName" placeholder="First Name" required><br>
                <input type="text" name="lastName" placeholder="Last Name" required><br>
                <input type="email" name="email" placeholder="Email" required><br>
                <input type="password" name="password" placeholder="Password" required><br>
                <button type="submit">Register</button></br></br>
                <button data-nav="home">Home</button>
            </form>
            
        `;
        setupRegister();
    } else if (page === 'post') {
        if (!isLoggedIn) {
            render('login');
            return;
        }
        app.innerHTML = `
            <h1>Create Post</h1>
            <form id="post-form">
                <input type="text" name="title" placeholder="Title" required><br>
                <textarea name="content" placeholder="Content" required></textarea><br>
                <button type="submit">Submit</button>
            </form>
            <a href="#" data-nav="home">Home</a>
        `;
        setupPost();
    } else if (page === 'logout') {

        fetch('/api/logout')
            .then(() => {
                clearCookies(); // Clear cookies
                render('login'); // Redirect to login page
            });
    } else {
        app.innerHTML = `<p>Page not found.</p>`;
    }
}

function clearCookies() {
    const cookies = document.cookie.split(";");

    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i];
        const eqPos = cookie.indexOf("=");
        const name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
        document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/";
    }
}