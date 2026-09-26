# Delivery board

Disposable test fixture for [Observed](https://github.com/esau-morais/observed) onboarding trials. Not a real project.

A Go server with a static page. Choosing a route loads its deliveries from `GET /api/deliveries?route=<name>`.

```bash
go run .
```

The server listens on `127.0.0.1:$PORT` (default 8080).
