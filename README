# XM Golang Exercise - v22.0.0

## Project setup
In order to install the application you need Docker and Docker Compose, so please make sure that these are installed on your computer.

#### Run
```bash
$ docker-compose up --build
```
This will install the following microservices (containers): MySQL, the Go application and PMA (phpmyadmin) for DB management.

### How to test it?
If the installation went successfully you could access the application under the following link: `127.0.0.1:8080`. The PMA is accessible on `127.0.0.1:8081`.

This is an API/JSON based application, so it does not have any kind of visual parts. This means that you have to use Postman or CURL for example to test the API endpoints. 

### Endpoints:

```go
    POST - /users/add // create a user
    PATCH - /users/:id
    GET - /users // get all the registered users 
```

The body of the post should be in the following form:

```json
{
    "email": "john.doe@example.com",
    "password": "password"
}
```
I could have hardcoded a user without exposing an API endpoint, but this approach was more elegant. We could to restrict this endpoint to not be publicly accessible. This will create an user together with a JWT token. We can check if the token has been generated correctly accessing the following endpoint.

```go
    POST - /login // login a user
```
This will return a JWT token like the following: `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2ODM4NzI1MjgsInVzZXJfaWQiOjF9.JkmBMnw7G2EbFgGs80dHy0Gdv7WKpREQmSwddJ-c3Dk"`.

For the upcoming operations we need to use this token as a Bearer Token in order to access the API endpoints. If you use Postman you have to put it in the `Authorization` section.

```go
    POST - /company/create
    GET - /company/:id
    PATCH - /company/:id
    DELETE - /company/:id
```

Body:
```json
{
    "name": "Company name",
    "description": "Company description",
    "employees": 10,
    "registered": true,
    "type": "NonProfit"
}
```

