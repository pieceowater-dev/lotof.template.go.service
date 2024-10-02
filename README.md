# lotof.template.go.service

## Description
🚀 This project is a microservice built with Go, leveraging GORM for database interactions and integrating with RabbitMQ for message brokering. It serves as a central point for managing requests and routing them to the appropriate services.

🔌 With a modular architecture, this microservice enables seamless communication across different services, facilitating efficient and scalable application development.

🔒 The use of GORM simplifies database operations, while RabbitMQ ensures reliable and asynchronous communication, allowing the microservice to effectively handle high volumes of requests.

🏗️ Whether you're developing a monolithic application or a distributed system, this microservice provides a robust foundation for your project.

## Installation

### Docker Installation
To pull the Docker image from the GitHub Container Registry:

1. Login to GitHub Container Registry:
   ```bash
   echo "YOUR_GITHUB_TOKEN" | docker login ghcr.io -u YOUR_GITHUB_USERNAME --password-stdin
   ```
2. Go to [GitHub Packages](https://github.com/orgs/pieceowater-dev/packages) and copy the command to pull the image.

### Manual Installation
To install the dependencies, run:
```bash
go mod tidy
```

## Running the Application

### Development Mode
To start the application in development mode:
```bash
go run main.go
```

### Production Mode
To start the application in production mode:
```bash
go build -o app
./app
```

### Debug Mode
To start the application in debug mode, use your preferred debugger (e.g., Delve):
```bash
dlv debug
```

## Environment Variables
The application uses the following environment variables, which should be defined in a [.env](.env) file at the root of the project:

```properties
RABBITMQ_DSN='amqp://guest:guest@localhost:5672/'
DATABASE_DSN='postgres://postgres:dima3raza@localhost:5432/go-template?sslmode=disable'
```

## Scripts
The following scripts are available in the `Makefile`:

#### Update dependencies
```bash
make update
```

#### Build the project
```bash
make build
```

#### Start the application
```bash
make start
```

#### Run unit tests
```bash
make test
```

#### Run unit tests with coverage
```bash
make test-coverage
```

## Testing
To run the unit tests:
```bash
make test
```

To run the unit tests with coverage:
```bash
make test-coverage
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author
![PCWT Dev Logo](https://avatars.githubusercontent.com/u/168465239?s=50)
### [PCWT Dev](https://github.com/pieceowater-dev)
