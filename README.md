# First REST API

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
- `entities`: Defines the data model for the `Student` and `Response` entities.
- `handlers`: Implements the HTTP request handlers for CRUD operations on student data.
- `database`: Manages the MongoDB database connection.
- `tests`: Responsible for writing and running test cases to ensure the correctness and reliability of application.

## Launching the Application

```shell
# Clone the repository:
git clone https://github.com/woozie-10/first-rest-api.git

# Change to the project directory:
cd first-rest-api

# Run the application:
sudo docker-compose up

```
The server should start, and you can access the API at http://localhost:8081.

## API Endpoints

- `GET /students`: Get a list of all students.
- `GET /students/:id`: Get a specific student by their ID.
- `POST /students`: Add a new student.
- `PUT /students/:id`: Update a student by ID.
- `DELETE /students/:id`: Delete a specific student by their ID.

# Running Tests in a project

To run tests in a project, follow these steps:

1. **Navigate to Project Root**: First, open your terminal and navigate to the root directory of project:

   ```shell
   cd first-rest-api
   ```
2. **Enter the Tests Directory**: Move to the tests directory within project:
  
    ```shell
    cd tests
    ```
3. **Run Tests**: Run the tests using the go test command. This command will automatically discover and execute all the test functions in project:

   ```shell
   go test
   ```

## Usage

You can use tools like [Postman](https://www.postman.com/) or `curl` to interact with the API endpoints.
