# Event Management API
### About
This Event Management API is a fast, scalable, and robust system designed with the Go programming language and the Gin web framework. It allows users to create, retrieve, and manage events efficiently. This project is ideal for developers looking to implement event management features in their applications or for anyone interested in learning how to build RESTful APIs with Go and Gin.

### Features
Create Events: Add new events with details including name, description, location, and time.
Fetch Events: Retrieve a list of all events or specific events by ID.
Robust Error Handling: Detailed error responses for better debugging.
Easy to Set Up: Simple setup process with minimal dependencies.

### Getting Started
Prerequisites
Go version 1.22
SQLite3

## Installation
Clone the repository
`git clone https://github.com/kortby/go-gin.git`


`cd event-management-api`
Install dependencies

This project uses the Gin web framework and the SQLite3 driver, which can be installed using:


`go get -u github.com/gin-gonic/gin`
`go get -u github.com/mattn/go-sqlite3`

### Initialize the database

Run the db.InitDB() function included in the project to set up the database with the required tables.

Start the server


`go run main.go`
The API will be available at http://localhost:8080.

Usage
Create an Event

`curl -X POST http://localhost:8080/events \
-H 'Content-Type: application/json' \
-d '{"name": "Tech Conference", "description": "Annual Tech Conference", "location": "New York", "datetime": "2024-05-23T10:00:00Z"}'`
Get All Events

`curl http://localhost:8080/events`
