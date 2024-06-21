// frontend/js/api.js

export function fetchPosts() {
    return fetch('/api/posts')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        });
}

export function submitPost(data) {
    return fetch('/api/post', {
        method: 'POST',
        body: data
    });
}

export function login(data) {
    return fetch('/api/login', {
        method: 'POST',
        body: data
    });
}

export function register(data) {
    return fetch('/api/register', {
        method: 'POST',
        body: data
    });
}

export function logout() {
    return fetch('/api/logout');
}
