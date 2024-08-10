# Go Web Server Example

This is a simple Go web server that serves static files and handles basic form submissions.

## Getting Started

1. Clone the repository:
  ```bash
  git clone https://github.com/Verna0214/Go_basic_server.git
  cd go-server
  ```
2. Build and run the server:
  ```
  go build
  ./go-server
  ```
3. Open your browser and navigate to http://localhost:8080.

## Endpoints

	•	/: Serves static files from the static directory.
	•	/hello: Responds with “hello!” for GET requests.
	•	/form: Handles POST requests and returns submitted form data.

## Example Form Submission

You can submit a form at http://localhost:8080/form.html. The server will parse the form and return the submitted data.