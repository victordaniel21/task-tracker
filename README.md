# 🎯 Focus - Fullstack Task Tracker

A modern, full-stack task management application built from the ground up using **Go (Golang)**, **React**, and **PostgreSQL**. 

This project demonstrates clean architecture, robust API design, and a responsive, highly-polished user interface with optimistic UI updates.

![Focus UI Preview](https://via.placeholder.com/800x400.png?text=Focus+Task+Tracker+Preview) ## ✨ Features

* **Full CRUD API:** Create, read, update, and delete tasks.
* **Optimistic UI:** Instant state updates on the frontend for a native-app feel without waiting for network latency.
* **Custom Modal Architecture:** Beautiful, non-blocking custom confirmation modals (replacing native browser alerts).
* **Graceful Shutdown:** The Go server catches OS signals and waits for active connections to finish before terminating.
* **CORS & JSON Handling:** Professional middleware for Cross-Origin requests and structured JSON error reporting.
* **Clean Architecture:** Strict separation of concerns (Models vs. Handlers) and Dependency Injection.

## 🛠️ Tech Stack

**Backend**
* Language: Go 1.22+ (Standard Library routing via `net/http`)
* Database: PostgreSQL (Running in Docker)
* Driver: `github.com/lib/pq`

**Frontend**
* Framework: React + TypeScript
* Build Tool: Vite
* Styling: Tailwind CSS
* Icons: Lucide React

## 🚀 Getting Started

### Prerequisites
Make sure you have the following installed on your machine:
* [Go](https://golang.org/doc/install) (v1.22 or higher)
* [Node.js](https://nodejs.org/) (v18 or higher)
* [Docker](https://www.docker.com/) (For PostgreSQL)

### 1. Database Setup
Start a PostgreSQL container on port 5433:
```bash
docker run --name focus-postgres -e POSTGRES_PASSWORD=secret -p 5433:5432 -d postgres
