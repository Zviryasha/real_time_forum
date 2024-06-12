// main.js
import { createFooter } from './footer.js';

// Create a WebSocket connection to the server
var ws = new WebSocket('ws://localhost:8080/ws');

// Function to run when the WebSocket connection is opened
ws.onopen = function(event) {
    console.log('WebSocket connection opened');
    createFooter();
    ws.send('Hello, server!');
};

// Function to run when a message is received from the server
ws.onmessage = function(event) {
    console.log('Message from server: ' + event.data);
};



ws.onerror = function(event) {
    console.log('WebSocket error observed:'+ event.data);
};

// Function to run when the WebSocket connection is closed
ws.onclose = function(event) {
    console.log('WebSocket connection closed');
};