# Hackathon Fiber Starter

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)

This project presents a boilerplate/starter kit for rapidly developing RESTful APIs using Go, Fiber, and PostgreSQL.

The application is built on [Go v1.22.3](https://tip.golang.org/doc/go1.22) and [PostgreSQL](https://www.postgresql.org/). Ensure you have installed the necessary dependencies before proceeding.

## Quick Start

To initiate a project, execute the following command:

```bash
go mod init <project-name>
```

## Manual Installation

For those who prefer a manual setup, follow these steps:

Clone the repository:

```bash
git clone --depth 1 https://github.com/ahargunyllib/hackathon-fiber-starter.git
cd hackathon-fiber-starter
rm -rf ./.git
```

Install the required dependencies:

```bash
go mod tidy
```

Configure the environment variables:

```bash
cp ./config/.env.example ./config/.env

# open .env and adjust the environment variables as needed
```

## Table of Contents

- [Features](#features)
- [Commands](#commands)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Error Handling](#error-handling)
- [Validation](#validation)
- [Authentication](#authentication)
- [Authorization](#authorization)
- [Logging](#logging)
- [Linting](#linting)
- [Convention](#convention)

## Features

- **SQL database**: [PostgreSQL](https://www.postgresql.org) Object Relational Mapping via [Gorm](https://gorm.io)
- **Validation**: request data validation utilizing [Package validator](https://github.com/go-playground/validator)
- **Logging**: implemented with [Logrus](https://github.com/sirupsen/logrus) and [Fiber-Logger](https://docs.gofiber.io/api/middleware/logger)
- **Testing**: unit and integration tests powered by [Testify](https://github.com/stretchr/testify) with formatted output using [gotestsum](https://github.com/gotestyourself/gotestsum)
- **Error handling**: centralized error management system
- **API documentation**: generated with [Swag](https://github.com/swaggo/swag) and [Swagger](https://github.com/gofiber/swagger)
- **Email functionality**: implemented using [Gomail](https://github.com/go-gomail/gomail)
- **Environment variables**: managed with [Godotenv](https://github.com/joho/godotenv)
- **Security**: HTTP headers secured by [Fiber-Helmet](https://docs.gofiber.io/api/middleware/helmet)
- **CORS**: Cross-Origin Resource-Sharing enabled through [Fiber-CORS](https://docs.gofiber.io/api/middleware/cors)
- **Compression**: gzip compression provided by [Fiber-Compress](https://docs.gofiber.io/api/middleware/compress)
- **Linting**: code quality ensured with [golangci-lint](https://golangci-lint.run)
- **Docker support**
- **Vercel support**

## Commands

For local execution:

```bash
make start
```

Alternatively, run with live reload:

```bash
air
```

> [!NOTE]
> Ensure that `Air` is installed.\
> Refer to 👉 [How to install Air](https://github.com/air-verse/air)

Docker:

```bash
# run docker container
make docker
```

Linting:

```bash
# run lint
make lint
```

Swagger:

```bash
# generate the swagger documentation
make swagger
```

Environment Variables

The environment variables are located in the `./config/.env` file and can be modified. They come with these default values:

```bash
# server configuration
# Env value : production || staging || development
APP_ENV=development
APP_PORT=8080
API_KEY=API_KEY

# database configuration
DB_HOST=postgresdb
DB_PORT=5432
DB_USER=postgres
DB_PASS=123456
DB_NAME=hackathon_fiber_starter

# OAuth2 configuration
GOOGLE_CLIENT_ID=<yourapps.googleusercontent.com>
GOOGLE_CLIENT_SECRET=<thisissamplesecret>

# GOMAIL configuration options for the email service
GOMAIL_HOST=smtp.gmail.com
GOMAIL_PORT=465
GOMAIL_USERNAME=<email-server-username>
GOMAIL_PASSWORD=<email-server-password>

# Firebase configuration
FIREBASE_BUCKET=<hackathon-fiber-starter.appspot.com>
FIREBASE_CREDENTIALS_PATH=config/firebase-credentials.json

# Redis configuration
REDIS_ADDR=redis:6379
REDIS_PASS=password

# JWT
JWT_SECRET_KEY=thisisasamplesecret
JWT_EXP_TIME=8h
```

## Project Structure

```zsh
├── cmd                 # executeables file
│  ├── app              # application entry
├── config              # config files (firebase admin, etc)
├── data
│  ├── seeders          # seeders data for db
│  │   ├── dev          # for development purposes
│  │   ├── prod         # for production also
├── deploy              # for deployment purposes
├── docs                # application documentation
├── domain              # entities structure, dtos, contracts
├── internal            # private application and libraries
│  ├── app              # application functionality
│  ├── infra            # application external systems
│  ├── middlewares      # application middlewares
├── pkg                 # external reusable packges and libraries
├── tests               # integration testing files
├── web                 # static html files
```

## API Documentation

To view the available APIs and their specifications, launch the server and navigate to `http://localhost:8080/v1/docs` in your web browser.

The documentation page is automatically generated using [Swag](https://github.com/swaggo/swag) definitions written as comments in the controller files.

Refer to 👉 [Declarative Comments Format.](https://github.com/swaggo/swag#declarative-comments-format)

## API Endpoints

Available routes:

**Essential routes**:
`GET /` - health check
`GET /api/v1` - health check

**Auth routes**:
COMING SOON

**User routes**:
COMING SOON

## Error Handling
The application incorporates a custom error handling mechanism, located in the `src/pkg/helpers/http/error_handler/error.go` file.

It also employs the `Fiber-Recover` middleware to gracefully manage any panics that may occur in the handler stack, preventing unexpected application crashes.

The error handling process returns an error response in the following format:

```json
{
  "error": "something not found"
}
```

Fiber provides a custom error struct via `fiber.NewError()`, allowing you to specify a response code and message. This error can be returned from any part of your code, and Fiber's `ErrorHandler` will automatically capture it.

For instance, if you're attempting to retrieve a user from the database but the user is not found, and you wish to return a 404 error, the code might resemble:

```go
func (s *userService) GetUserByID(c *fiber.Ctx, id string) {
	user := new(model.User)

	err := s.DB.WithContext(c.Context()).First(user, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}
}
```

## Validation

Request data validation is performed using [Package validator](https://github.com/go-playground/validator). Consult the [documentation](https://pkg.go.dev/github.com/go-playground/validator/v10) for more details on writing validations.

## Authentication

COMING SOON

## Authorization

COMING SOON

## Logging

Import the logger from `src/pkg/log`. It utilizes the [Logrus](https://github.com/sirupsen/logrus) logging library.

Logging should adhere to the following severity levels (in ascending order from most to least important):

```go
log.Fatal(log.LogInfo{
		"error": error
}, "message")
log.Error(log.LogInfo{
		"error": error
}, "message")
log.Warn(log.LogInfo{
		"warning": warning
}, "message")
log.Info(log.LogInfo{
		"data": data
}, "message")
```

> [!NOTE]
> API request information (request url, response code, timestamp, etc.) is automatically logged (using [Fiber-Logger](https://docs.gofiber.io/api/middleware/logger)).

## Linting

Linting is performed using [golangci-lint](https://golangci-lint.run)

See 👉 [How to install golangci-lint](https://golangci-lint.run/welcome/install)

To modify the golangci-lint configuration, update the `.golangci.yml` file.

## Convention

Please review and adhere to the conventions outlined [here](./CONVENTION.md)

## Inspirations

- [hagopj13/node-express-boilerplate](https://github.com/hagopj13/node-express-boilerplate)
- [khannedy/golang-clean-architecture](https://github.com/khannedy/golang-clean-architecture)
- [zexoverz/express-prisma-template](https://github.com/zexoverz/express-prisma-template)
- [indrayyana/go-fiber-boilerplate](https://github.com/indrayyana/go-fiber-boilerplate)
- [devanfer02/nosudes-be](https://github.com/devanfer02/nosudes-be)
- [kmdavidds/abdimasa-backend](https://github.com/kmdavidds/abdimasa-backend)
- [nathakusuma/sea-salon-be](https://github.com/nathakusuma/sea-salon-be)

## License

[MIT](LICENSE)
