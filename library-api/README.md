# Library Management System API (Go)

## Project Overview
This is a REST API built for the Infosys Go Capstone Project. [cite_start]It manages library inventory, checkouts, and handles reservations when books are out of stock[cite: 71, 74].

## Setup Instructions
1. Install Go (1.21+)
2. Clone the repository
3. Run `go mod tidy` to install dependencies
4. Run the server: `go run main.go`

## API Endpoints
- **POST /books**: Add a new book (JSON body: title, author, available_copies)
- [cite_start]**POST /checkout**: Borrow a book (JSON body: user_id, book_id) [cite: 73]