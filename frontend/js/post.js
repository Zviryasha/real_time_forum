// frontend/js/post.js

import { submitPost } from './api.js';

export function setupPost() {
    const form = document.getElementById('post-form');
    form.addEventListener('submit', function(event) {
        event.preventDefault();
        const data = new FormData(form);
        submitPost(data)
            .then(response => {
                if (response.ok) {
                    render('home');
                } else {
                    console.error('Failed to submit post');
                }
            })
            .catch(error => {
                console.error('Error submitting post:', error);
            });
    });
}

export function displayPosts(posts) {
    const postsDiv = document.getElementById('posts');
    postsDiv.innerHTML = '';
    if (!posts || posts.length === 0) {
        postsDiv.innerHTML = '<p>No posts available.</p>';
    } else {
        posts.forEach(post => {
            const postDiv = document.createElement('div');
            postDiv.className = 'post';
            postDiv.innerHTML = `
                <h2>${post.title}</h2>
                <p>${post.content}</p>
                <small>Posted by User ${post.userId} on ${new Date(post.createdAt).toLocaleString()}</small>
            `;
            postsDiv.appendChild(postDiv);
        });
    }
}
