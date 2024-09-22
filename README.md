# Golang env validator example

This project is a small application that reads values from a .env file and validates them against a schema specified in the file located at <code>./config/.env</code>.

## Installation

To run this project, please make sure to install the dependencies first. You can do this by running the following command:
~~~
go build -o build && ./build
~~~

## Usage

After installing the dependencies, you can execute the script to read values from the .env file and validate them against the schema. The script will output any validation errors encountered.

## Configuration

The schema for validating the values from the .env file is specified in the file located at ./config/.env. Make sure to update this file with the correct schema before running the script. Here is an example of the updated schema:

~~~
type Env struct {
  PORT   string validate:"required,numeric"
  TOKEN  string validate:"required,alphanum"
  SECRET string validate:"required,alphanum"
}
~~~

Ensure that the schema matches the structure of your .env file. The validation functionality is provided by the github.com/go-playground/validator/v10 package.

## Dependencies

This project relies on Go for building and running the application, as well as the github.com/go-playground/validator/v10 package for property validation. Please ensure that you have Go installed on your machine before proceeding.

## License

This project is licensed under the MIT License. Feel free to use, modify, and distribute it as needed.
