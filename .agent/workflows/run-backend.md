---
description: how to run the Go backend
---

To run the backend, follow these steps:

1. Open your terminal.
2. Navigate to the backend directory:
   ```powershell
   cd e:\wasooli\backend
   ```
3. Run the backend in development mode:
   ```powershell
   go run cmd/api/main.go
   ```

Alternatively, to build and run the production binary:

1. Build the binary:
   ```powershell
   go build -o api.exe ./cmd/api
   ```
2. Run the binary:
   ```powershell
   .\api.exe
   ```

The server will be available at `http://localhost:8080`.
You can check if it's running by visiting `http://localhost:8080/health`.
