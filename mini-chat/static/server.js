const WebSocket = require('ws');

const server = new WebSocket.Server({ port: 8080 });

const clients = new Set();
const messages = [];

server.on('connection', (ws) => {
    clients.add(ws);

    messages.forEach((msg) => ws.send(msg));

    ws.on('message', (message) => {
        messages.push(message); 

        clients.forEach((client) => {
            if (client.readyState === WebSocket.OPEN) {
                client.send(message);
            }
        });
    });

    ws.on('close', () => {
        clients.delete(ws);
    });
});

console.log('WebSocket server is running on ws://localhost:8080/ws');