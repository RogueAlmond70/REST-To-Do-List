Problem:
You are building a simple RESTful API for a to-do list application. The API should allow users to create, read, update, and delete to-do items. Each to-do item has the following properties:

ID (integer)
Title (string)
Description (string)
Completed (boolean)
Created At (timestamp)
Your task is to implement the following endpoints:

GET /todos
This endpoint should return a list of all to-do items in the database. The list should be sorted by the Created At field in ascending order.

GET /todos/{id}
This endpoint should return a single to-do item with the given ID. If no to-do item with the given ID exists, it should return a 404 error.

POST /todos
This endpoint should create a new to-do item in the database. It should accept a JSON request body with the following format:


{
"title": "string",
"description": "string",
"completed": boolean
}

The ID, Created At, and Completed fields should be generated automatically by the server.

PUT /todos/{id}
This endpoint should update an existing to-do item with the given ID. It should accept a JSON request body with the same format as the POST endpoint above. If no to-do item with the given ID exists, it should return a 404 error.

DELETE /todos/{id}
This endpoint should delete an existing to-do item with the given ID. If no to-do item with the given ID exists, it should return a 404 error.

Requirements:
Your code should be written in Go.
You should use a SQL database to store the to-do items (you can use any database engine you like).
You should use a third-party library for handling the HTTP server and routing (e.g. net/http, gorilla/mux).
You should use a third-party library for handling the database (e.g. database/sql, gorm).
You should write tests for your code.
Deliverables:
The source code for your solution.
A script for setting up the database and running the server (e.g. a Makefile).
Instructions for running the tests.
Any other supporting documentation (

Evaluation Criteria:
Your solution will be evaluated based on the following criteria:

Correctness: Does the code implement the to-do list API correctly?
Code quality: Is the code well-organized, readable, and maintainable?
Testing: Are the tests comprehensive and do they cover the important cases?
Documentation: Are the instructions for running the code and the supporting documentation clear and sufficient?