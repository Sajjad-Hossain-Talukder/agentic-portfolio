# Agentic Portfolio

> A personal chatbot application built with a no-framework, agentic approach, leveraging React for frontend and Go for backend, integrated with Gemini Flash 2.5 as the backend agent.

Explore **Sajjad's career highlights, expertise, and experiences** through an interactive chat interface.


---

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Frontend](#frontend)
- [Backend](#backend)
- [Getting Started](#getting-started)
- [Custom Data](#custom-data)
- [Usage](#usage)
- [Screenshots](#screenshots)
- [License](#license)

---

## Overview
**Agentic Portfolio** is a personal chatbot application designed to answer questions based on LinkedIn and CV PDF data. The application is built using a minimal framework approach to practice the agentic methodology.

- **Frontend:** React
- **Backend:** Go
- **Agent:** Gemini Flash 2.5
- **Notification:** Pushover API for unknown questions or user contact requests
- **Hosting:** Docker (live at [Let's Chat](http://160.191.162.38:5173/))


---

## Features
- Answer questions using LinkedIn and CV data.
- Notify the owner in case of unknown questions via Pushover.
- No database required — lightweight and easy to deploy.
- Fully Dockerized for easy deployment.
- Allows users to use their own PDFs for customization.

---

## Architecture

The chatbot follows a lightweight agentic design, connecting the user interface, backend, and AI agent seamlessly:


- **Frontend**: Built with React, it provides a responsive chat interface served via Nginx in Docker.  
- **Backend**: Go server handles API requests, communicates with the agent, and triggers notifications.  
- **Agent**: Gemini Flash 2.5 processes queries using Sajjad’s professional information and provides intelligent responses.  
- **Notifications**: Pushover API alerts the user or admin in case of unknown questions or when user engagement is required.  

This architecture demonstrates a minimal **agentic system** without relying on heavy frameworks, focusing on modular, tool-driven responses.

## Frontend

The frontend of this project is a **React application** built with **Vite** and styled using Tailwind CSS. It provides a responsive chat interface for interacting with the agent.

### Project Structure

```bash
chatbot/
├─ public/
│ └─ favicon.svg
├─ src/
│ ├─ assets/ # Images, icons, and static assets
│ ├─ components/ # Reusable React components
│ ├─ pages/ # Individual pages (e.g., chat page)
│ ├─ root/ # Application entry point and root layout
│ ├─ services/ # API services and utilities
│ ├─ styles/ # Global styles and Tailwind setup
│ ├─ vite-env.d.ts # Vite TypeScript environment definitions
├─ .dockerignore
├─ .gitignore
├─ Dockerfile
├─ Makefile
├─ README.md
├─ eslint.config.js
├─ index.html
├─ package-lock.json

``` 
**Explanation of key folders:**

- `src/components/` → Contains reusable UI elements such as buttons, chat bubbles, and forms.  
- `src/pages/` → Defines individual pages, like the main chat interface.  
- `src/services/` → Handles API calls to the backend server.  
- `src/root/` → Entry point of the React app (`App.tsx` or `main.tsx`).  
- `src/assets/` → Stores images, icons, and static media.  
- `src/styles/` → Contains Tailwind CSS setup and additional global styles.  
- `Dockerfile` → Defines the frontend Docker image for deployment.

This structure ensures **modular development** and separates concerns between components, pages, and services, making the frontend maintainable and scalable.


## Backend

The backend of this chatbot is built with **Go** and uses **Gemini Flash 2.5** as the agent engine. It is designed with a **framework-free agentic approach**, focusing on simplicity and direct control over task execution.

### Key Features

- **Agentic Processing:**  
  Handles queries using a lightweight agentic approach without any heavy frameworks. Processes user input, decides actions, and generates responses dynamically.

- **Data Source:**  
  Uses **LinkedIn profile data** and **CV PDFs** fed via the initial prompt. No database is used; all data resides in memory.

- **Notifications:**  
  Sends notifications via the **Pushover API** when the agent encounters an unknown question or when the user requests human intervention.

- **API Endpoints:**  
  - `/chat` – send messages and receive responses  
  - `/status` – optional health/status check

- **Docker Deployment:**  
  The backend runs in a Docker container and is accessible at [Let's Chat](http://160.191.162.38:5173/).

### Backend Flow

```bash 
    User --> |Message| API[Go Backend API]
    API --> |Process| Agent[Gemini Flash 2.5 Agent]
    Agent --> |Response| API
    API --> |Send| User
    Agent --> |Unknown Question| Pushover[Pushover Notification]
```
## Getting Started

You can easily build and run both the backend and frontend using the provided **Makefiles**.

### Backend

Navigate to the backend folder and use the following commands:

| Command           | Description                                  |
|------------------|----------------------------------------------|
| `make build`      | Build the backend Docker image               |
| `make run`        | Run the backend container on port 8082       |
| `make stop`       | Stop the backend container                   |
| `make logs`       | View backend logs                            |
| `make remove`     | Remove the backend container                 |
| `make run-server` | Run the backend Go server locally without Docker |

### Frontend (UI)

Navigate to the frontend folder and use the following commands:

| Command         | Description                                |
|-----------------|--------------------------------------------|
| `make build`    | Build the frontend Docker image            |
| `make run`      | Run the frontend container on port 5173    |
| `make stop`     | Stop the frontend container                |
| `make logs`     | View frontend logs                          |
| `make remove`   | Remove the frontend container              |
| `make run-local`| Run the frontend locally with `yarn dev`  |

## Custom Data

The chatbot is initialized with **LinkedIn profile data** and **CV PDFs**. You can customize the data by:

1. Replacing the existing PDFs or LinkedIn data in the backend memory feed. Check them in `resources` folder  
2. Adjusting the **initial prompt** in the backend agent code to include new data or context.

> ⚠️ Note: No database is used; all custom data is loaded in memory during runtime.

---


## Usage

### 1. Clone the Repository

Begin by cloning the repository to your local machine:

```bash
git clone https://github.com/Sajjad-Hossain-Talukder/agentic-portfolio.git
cd agentic-portfolio
```

### 2. Using Docker and Makefiles

#### Backend

Navigate to the backend directory and use the provided Makefile to build and run the backend:

```bash
cd chatbot-backend

# Build the backend Docker image
make build

# Run the backend container on port 8082
make run

# To stop the backend container
make stop

# View backend logs
make logs

# Remove the backend container
make remove

# Alternatively, run the backend server locally without Docker
make run-server

```

#### Frontend (UI)

Navigate to the frontend directory and use the provided Makefile to build and run the frontend:
```bash
cd chatbot

# Build the frontend Docker image
make build

# Run the frontend container on port 5173
make run

# To stop the frontend container
make stop

# View frontend logs
make logs

# Remove the frontend container
make remove

# Run the frontend locally with Vite
make run-local

```

## Screenshots 

<img width="1461" height="796" alt="Screenshot 2025-09-06 at 11 21 21 PM" src="https://github.com/user-attachments/assets/aba54451-fc4b-40fb-910f-5b5c21e78feb" />
<img width="1461" height="796" alt="Screenshot 2025-09-06 at 11 21 49 PM" src="https://github.com/user-attachments/assets/aee2fbf3-abf8-4b0e-a862-69ede560b6fa" />

## License

This project is licensed under the **MIT License**.  

See the full license text in the [LICENSE](./LICENSE) file.  

---

### Summary

- You are free to **use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies** of this project.  



