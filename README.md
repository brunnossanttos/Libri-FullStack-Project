# Skoob Clone

A clone of the Skoob app built to learn Go for the backend and modern frontend development.

## Table of Contents

- [Description](#description)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running Locally](#running-locally)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Description

This project is a learning exercise to recreate the core features of Skoob—a Brazilian social network for book enthusiasts. The backend is written in Go, and the frontend uses React.

## Tech Stack

- **Backend**: Go, Gin (or net/http), GORM
- **Database**: SQLite (development), PostgreSQL (production)
- **Frontend**: React, Vite
- **Authentication**: JWT
- **API Documentation**: Swagger

## Architecture

The application follows a RESTful API design. The backend serves JSON endpoints consumed by the React frontend.

## Features

- User registration and login (JWT-based)
- Personal book collections: “Want to Read”, “Reading”, “Read”
- Book search by title, author, or ISBN
- Rate and review books
- User profile management

## Getting Started

### Prerequisites

- Go 1.20+
- Node.js 16+
- npm or yarn

### Installation

```bash
# Clone the repository
git clone https://github.com/your-username/skoob-clone.git
cd libri-api
```

### Running the Backend

```bash
cd libri
go mod download
go run main.go
```

By default, the backend listens on port `8080`.

### Running the Frontend

```bash
cd frontend
npm install
npm run dev
```

The frontend will run on `http://localhost:8080`.

## Usage

- Register a new user via the frontend UI.
- Browse and manage your book collections.
- Search for books and add them to your lists.
- Rate and review books you’ve read.

## Project Structure

```
/libri-api        # Go API
  /cmd
  /docs
  /internal
  /migrations

/frontend       # React application
  /src
    /components
    /pages
    /services
  package.json
```

## Contributing

Contributions are welcome! Please fork the repository and open a pull request with your changes.

## License

Distributed under the MIT License. See `LICENSE` for more information.
