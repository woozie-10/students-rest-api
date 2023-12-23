# STUDENTS REST API

This is a simple REST API project that manages student information in a MongoDB database using the Go programming language and the Gin web framework.

## Technologies

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [MongoDB](https://www.mongodb.com/)
- [Docker](https://www.docker.com/)
- [Resty](https://github.com/go-resty/resty)
- [Viper](https://github.com/spf13/viper)

## Project Structure

The project is organized into the following main components:

- `cmd`: Contains the main application entry point.
- `config`: Contains configuration settings and management for the web application.
- `router`: Manages Gin router settings for efficient HTTP routing and request handling in the web application.
- `entities`: Defines the data model for the `Student` entitie.
- `handlers`: Implements the HTTP request handlers for CRUD operations on student data.
- `database`: Manages the MongoDB database connection.
- `tests`: Responsible for writing and running test cases to ensure the correctness and reliability of application.

## Launching the Application

```shell
# Clone the repository:
git clone https://github.com/woozie-10/students-rest-api.git

# Change to the project directory:
cd students-rest-api

# Run the application:
make build && make run

```
The server should start, and you can access the API at http://localhost:8081.

## API Endpoints
- GET /students: Retrieve a list of all students.
- GET /students/:username: Retrieve information about a specific student identified by their username.
- GET /students/group/:group: Retrieve a list of students belonging to a specific group.
- GET /students/course/:course: Retrieve a list of students enrolled in a specific course.
- POST /students: Add a new student to the system.
- PUT /students/:username: Update information for a specific student identified by their username.
- DELETE /students/:username: Delete a student record based on their username.

# Running Tests in a project

To run tests in a project, follow these steps:

1. **Navigate to Project Root**: First, open your terminal and navigate to the root directory of project:

   ```shell
   cd students-rest-api
   ```
2. **Run Tests**: Run the tests using the command. This command will automatically discover and execute all the test functions in project:

   ```shell
   make test
   ```

## Usage

You can use tools like [Postman](https://www.postman.com/) or `curl` to interact with the API endpoints.
