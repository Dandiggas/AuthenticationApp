
# Authentication Application

This project is an authentication application with a React frontend and a Go (Golang) backend. Currently, the backend is in development, and it handles user registration. 

## Backend Setup

The backend is built using the Go programming language with the Fiber framework. To set up and run the backend, follow these steps:

### Prerequisites

- Install Go (version 1.x): Follow the instructions on the [official Go website](https://golang.org/doc/install).
- A SQL database (e.g., PostgreSQL, MySQL) running and accessible.
- Set up your database credentials.

### Installation

1. **Clone the Repository**: First, clone the repository to your local machine.

   ```bash
   git clone https://github.com/Dandiggas/AuthenticationApp
   cd /Authenticationapp
   ```

2. **Install Dependencies**: Navigate to the backend directory and install the required Go dependencies.

   ```bash
   cd /Authenticationapp
   go mod tidy
   ```

3. **Database Configuration**: Open the database configuration file (usually located in a config file or directly in the main Go files). Insert your database credentials (host, user, password, dbname).

4. **Run the Application**: Start the server using the Go command.

   ```bash
   go run main.go
   ```

   Your backend server should now be running and listening for requests.

### Environment Variables

- Consider using environment variables for sensitive information like database credentials and secret keys.
- Use a `.env` file or export environment variables directly in your shell.

### API Endpoints

Document your API endpoints here as you develop them.

## Frontend Setup

Instructions for setting up the React frontend will go here once developed.


