# Hackathon Fiber Starter

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)

## Description

This project presents a boilerplate/starter kit for rapidly developing RESTful APIs using Go, Fiber, and PostgreSQL.

The application is built on [Go v1.23.4](https://tip.golang.org/doc/go1.22) and [PostgreSQL](https://www.postgresql.org/). It uses [Fiber](https://docs.gofiber.io/) as the HTTP framework and [pgx](https://github.com/jackc/pgx) as the driver and [sqlx](github.com/jmoiron/sqlx) as the query builder. It also utilizes [Redis](https://redis.io/) as the caching layer with [go-redis](https://github.com/redis/go-redis/) as the client.

## Getting started

1. Ensure you have [Go](https://go.dev/dl/) 1.23 or higher and [Task](https://taskfile.dev/installation/) installed on your machine:

   ```bash
   go version && task --version
   ```

2. Create a copy of the `.env.example` file and rename it to `.env`:

   ```bash
   cp ./config/.env.example ./config/.env
   ```

   Update configuration values as needed.

3. Install all dependencies, run docker compose, create database schema, and run database migrations:

   ```bash
   task
   ```

4. Run the project in development mode:

   ```bash
   task dev
   ```

## Documentation

For database schema documentation, see [here](https://dbdocs.io/ahargunyllib/hackathon-fiber-starter), powered by [dbdocs.io](https://dbdocs.io/).

For API documentation, see [here](n1nxox08nh.apidog.io), powered by [Apidog](https://apidog.com/).

## Features

- **Migration**: database schema migration using [golang-migrate](https://github.com/golang-migrate/migrate)
- **Validation**: request data validation utilizing [Package validator](https://github.com/go-playground/validator)
- **Logging**: implemented with [zerolog](https://github.com/rs/zerolog)
- **Testing**: unit and integration tests powered by [Testify](https://github.com/stretchr/testify) with formatted output using [gotestsum](https://github.com/gotestyourself/gotestsum)
- **Error handling**: centralized error management system
- **Email functionality**: implemented using [Gomail](https://github.com/go-gomail/gomail)
- **Environment variables**: managed with [Viper](https://github.com/spf13/viper)
- **Security**: HTTP headers secured by [Fiber-Helmet](https://docs.gofiber.io/api/middleware/helmet)
- **CORS**: Cross-Origin Resource-Sharing enabled through [Fiber-CORS](https://docs.gofiber.io/api/middleware/cors)
- **Compression**: gzip compression provided by [Fiber-Compress](https://docs.gofiber.io/api/middleware/compress)
- **Linting**: code quality ensured with [golangci-lint](https://golangci-lint.run)
- **Docker support**
- **Vercel support**

## Convention

Please review and adhere to the conventions outlined [CONVENTION](./CONVENTION.md)

## Contributing

Developers interested in contributing can refer to the [CONTRIBUTING](CONTRIBUTING.md) file for detailed guidelines and instructions on how to contribute.

## Inspirations

- [hagopj13/node-express-boilerplate](https://github.com/hagopj13/node-express-boilerplate)
- [khannedy/golang-clean-architecture](https://github.com/khannedy/golang-clean-architecture)
- [zexoverz/express-prisma-template](https://github.com/zexoverz/express-prisma-template)
- [indrayyana/go-fiber-boilerplate](https://github.com/indrayyana/go-fiber-boilerplate)
- [devanfer02/nosudes-be](https://github.com/devanfer02/nosudes-be)
- [kmdavidds/abdimasa-backend](https://github.com/kmdavidds/abdimasa-backend)
- [nathakusuma/sea-salon-be](https://github.com/nathakusuma/sea-salon-be)
- [bagashiz/go-pos](https://github.com/bagashiz/go-pos)

## License

[MIT](LICENSE)
