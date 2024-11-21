# **Movie Ticket Booking Application**
 
This repository contains a Go-based web server that serves a simple, static movie ticket booking application. The server serves static files such as HTML, CSS, and JavaScript, and includes unit tests to ensure proper functionality.
 
---
 
## **Features**
- Serves static HTML files with a pre-designed ticket booking form.
- Supports serving JavaScript and CSS files from structured directories.
- Includes unit tests for verifying file server functionality.
- Responsive design for a user-friendly booking experience.
 
---
## **Commands**
Start the server:
``` bash
go run main.go
```
Access the application in your browser at:
``` bash
http://localhost:80
```
Run the test suite:
``` bash
go test
go test -v
```
Testing includes:
- Validating HTTP status codes for existing and non-existent files.
- Verifying MIME types for HTML, CSS, and JavaScript files.
- Ensuring the default file (index.html) serves correctly.
 
# **Test Scenarios**
#### Serving Existing Files:
- Verifies that files present in the server's directories are served correctly.
- Ensures the server responds with:
- HTTP Status Code: 200 OK.
Handling Non-Existent Files:
- Tests requests for files that do not exist on the server.
#### Confirms the server returns:
- HTTP Status Code: 404 Not Found.
#### MIME Type Validation:
- Ensures that files are served with the correct MIME type based on their extensions.
#### MIME types validated:
- HTML → templates/index.html
- CSS → templates/style.css
- JavaScript → templates/script.js
#### Default File Handling:
- Validates that the server serves the default index.html file when the root path (/) is accessed.
#### Confirms the server responds with:
- HTTP Status Code: 200 OK.
