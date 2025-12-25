# 💬 Go & React Real-Time Chat Application

![Go](https://img.shields.io/badge/Backend-Go-00ADD8?logo=go&logoColor=white)
![React](https://img.shields.io/badge/Frontend-React-61DAFB?logo=react&logoColor=black)
![WebSocket](https://img.shields.io/badge/Protocol-WebSocket-orange)
![License](https://img.shields.io/badge/License-MIT-green)

A high-performance, real-time messaging application built with a **Go (Golang)** backend and a **React** frontend. This project utilizes a hybrid architecture: **REST API** for authentication and user management, and **WebSockets** for bi-directional, low-latency live communication.

## 🚀 Features

* **Real-time Messaging:** Instant message delivery using WebSockets.
* **Hybrid Architecture:** REST for stateless operations (Auth) and WebSockets for stateful operations (Chat).
* **Authentication:** Secure user login and registration using JWT (JSON Web Tokens).
* **Message History:** Persistent chat history stored in the database.
* **Scalable Backend:** Built on Go's powerful concurrency model (Goroutines).
* **Modern Frontend:** Responsive UI built with React and modern hooks.

## 🛠 Tech Stack

### Backend (Server)
* **Language:** Go (Golang)
* **Framework:** Gin / Fiber / Standard `net/http` (Update this based on your choice)
* **WebSocket Library:** `gorilla/websocket` or `nhooyr.io/websocket`
* **Database:** PostgreSQL / MongoDB
* **Auth:** JWT (JSON Web Tokens)

### Frontend (Client)
* **Framework:** React.js (Vite)
* **Styling:** Tailwind CSS / CSS Modules
* **State Management:** Context API / Redux
* **HTTP Client:** Axios / Fetch API

## 📂 Project Structure

```text
root/
├── backend/            # Go Backend source code
│   ├── cmd/            # Application entry point
│   ├── internal/       # Business logic (handlers, models, websockets)
│   ├── pkg/            # Utility packages (database, helpers)
│   ├── go.mod          # Go module definitions
│   └── main.go
├── frontend/           # React Frontend source code
│   ├── src/
│   │   ├── components/ # Chat, Login, Register components
│   │   ├── context/    # WebSocket and Auth context
│   │   └── App.jsx
│   ├── package.json
│   └── vite.config.js
└── README.md