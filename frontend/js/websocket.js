// real-time-forum/frontend/js/websocket.js

export function initWebSocket() {
    const socket = new WebSocket('ws://localhost:8080/ws');
    
    socket.onopen = () => console.log('WebSocket is connected.');
    socket.onmessage = (event) => handleMessage(event);
    socket.onclose = () => console.log('WebSocket is disconnected.');
    socket.onerror = (error) => console.log('WebSocket error:', error);
    
    return socket;
}

function handleMessage(event) {
    const data = JSON.parse(event.data);
    console.log('WebSocket message received:', data);
    // Handle the parsed data
}
