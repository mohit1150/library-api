# AI Usage Transparency Document

This document outlines the high-level conceptual and architectural assistance received from Gemini (AI) to support the independent development of this project.

## 1. Architectural Guidance
- **Inquiry:** "Which project topic is best suited for demonstrating Go's core features like concurrency and data handling?"
- **Guidance Received:** Recommendation to build the Library Management System to showcase state management (book availability) and concurrent reservations.

## 2. Technical Stack Selection
- **Inquiry:** "What are the industry-standard Go packages for building a lightweight REST API with a database?"
- **Guidance Received:** Suggestions to use the Gin framework for routing and GORM with SQLite for efficient data persistence.

## 3. Core Logic & Syntax Assistance
- **Inquiry:** "How do I structure a Go struct to represent a many-to-one relationship between checkouts and books?"
- **Guidance Received:** Provided syntax examples for GORM models and the basic structure of a Gin POST handler. 

## 4. Problem Solving & Troubleshooting
- **Inquiry:** "How should a REST API logically handle a scenario where a resource (a book) is temporarily out of stock?"
- **Guidance Received:** Conceptual explanation of a reservation queue and how to return specific HTTP status codes (like 409 Conflict) when a checkout fails.

## 5. Deployment & Documentation
- **Inquiry:** "What are the standard documentation requirements for a professional Go project?"
- **Guidance Received:** Best practices for writing a README.md and documenting concurrency strategies for external reviewers.