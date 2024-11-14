# Gowire

Gowire is a simple web server application written in Go. It uses the `net/http` package to handle HTTP requests and serves HTML templates. The application is structured with a focus on modularity and separation of concerns.

## Project structure

* `.air.toml`: Configuration file for the `air` live reloading tool.
* `.gitignore`: Specifies files and directories to be ignored by Git.
* `.vscode/settings.json`: Configuration settings for Visual Studio Code.
* `cmd/server/main.go`: Entry point of the application. It sets up the server and starts listening on port 8080.
* `go.mod`: Go module file that defines the module path and dependencies.
* `internal/handlers/handlers.go`: Contains the HTTP handlers for different routes and functions to render HTML templates.
* `internal/handlers/templates/about.html`: HTML template for the "About" page.
* `internal/handlers/templates/contact.html`: HTML template for the "Contact" page.
* `internal/handlers/templates/home.html`: HTML template for the "Home" page.
* `internal/handlers/templates/private.html`: HTML template for the "Private" page.
* `internal/handlers/templates/templates.go`: Functions to parse and manage HTML templates.
* `internal/middleware/middleware.go`: Middleware functions for logging and authentication.
* `internal/router/router.go`: Sets up the routes and applies middleware to the handlers.
* `web/static/js/light.js`: JavaScript file for handling client-side routing and dynamic content loading.

## Getting started

### Prerequisites

* Go 1.22 or later
* `air` live reloading tool (optional)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/edimoldovan/gowire.git
   cd gowire
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. (Optional) Install `air` for live reloading:
   ```sh
   go install github.com/cosmtrek/air@latest
   ```

### Running the application

1. Start the server:
   ```sh
   go run cmd/server/main.go
   ```

2. Open your browser and navigate to `http://localhost:8080`.

### Using `air` for live reloading

1. Start `air`:
   ```sh
   air
   ```

2. Open your browser and navigate to `http://localhost:8080`. The server will automatically reload when you make changes to the code.

## Routes

The application has the following routes:

* `/`: Home page
* `/about`: About page
* `/contact`: Contact page
* `/private`: Private page (requires authentication)
* `/light.js`: JavaScript file for client-side routing

## Middleware

The application uses middleware for logging and authentication:

* `Logger`: Logs the request method, URI, and duration.
* `Auth`: Checks for the presence of an `Authorization` header and returns a 401 Unauthorized response if it's missing.

## Templates

HTML templates are located in the `internal/handlers/templates` directory. The templates are parsed and managed using the `html/template` package. Custom template functions are defined in `internal/handlers/templates/templates.go`.

## JavaScript

The `web/static/js/light.js` file contains the client-side logic for handling routing and dynamic content loading. It uses the Fetch API to load content and update the DOM without reloading the page.
