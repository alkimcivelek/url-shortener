# URL Shortener Service

A production-ready URL shortening service built with Go, following Domain Driven Design principles.

## Features

- **URL Shortening**: Convert long URLs to short, manageable links
- **URL Redirection**: HTTP 301 redirects to original URLs
- **Access Statistics**: Track click counts for shortened URLs
- **Health Monitoring**: Built-in health check endpoints
- **Kubernetes Ready**: Deployable with multiple replicas


## Architecture

The service follows Clean Architecture and DDD patterns:

```
domain/          # Business entities and rules
application/     # Use cases and orchestration
infrastructure/  # Data persistence and external services
adapter/         # HTTP handlers and DTOs
```


## API Endpoints

POST | `/api/v1/shorten` | Create short URL |
GET | `/{shortCode}` | Redirect to original URL |
GET | `/api/v1/stats/{shortCode}` | Get URL statistics |
GET | `/health` | Health check |


## Test the API

curl -X POST http://localhost:8080/api/v1/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://test.com/very/long/path"}'

curl -s http://localhost:8080/api/v1/stats/shortcode


## Configuration

Set environment variables:

- `PORT` - Server port (default: 8080)
- `BASE_URL` - Base URL for shortened links (default: http://localhost:8080)


## Testing

### Run tests
go test ./tests/...


## Deployment

### Local Docker
```bash
docker build -t url-shortener .
docker run -p 8080:8080 url-shortener
```

### Kubernetes
```bash
kubectl apply -f k8s/
kubectl get pods -l app=url-shortener
```

The service runs with 2+ replicas for high availability and includes health checks for reliability.


## Development

The codebase uses only Go standard library and follows these principles:

- **Clean Architecture**: Clear separation of concerns
- **Domain Driven Design**: Business logic isolation
- **SOLID Principles**: Maintainable and extensible code
- **Thread Safety**: Concurrent request handling
- **Graceful Shutdown**: Proper resource cleanup


## Technical Decisions

- **In-Memory Storage**: Fast access with thread-safe operations
- **Deterministic Hashing**: Consistent short codes across replicas
- **Value Objects**: Immutable domain concepts
- **Use Cases**: Single responsibility business operations
- **Dependency Injection**: Testable and flexible design