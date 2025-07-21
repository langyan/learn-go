
# 代码结构
```
/go-service/
├── cmd/                → Main entrypoints (http, cli)
├── internal/
│   ├── handler/        → HTTP handlers (controllers)
│   ├── service/        → Business logic layer
│   └── repository/     → DB access layer
├── model/              → Shared domain structs
├── pkg/                → Reusable helpers (logger, config)
├── go.mod

```