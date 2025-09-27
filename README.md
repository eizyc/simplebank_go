# GoFinanceServer

**GoFinanceServer** is a backend banking service written in Go.  
It provides secure APIs and background processing for user management, account operations, and money transfers.  
The project demonstrates best practices in Go for API development, database management, and asynchronous job handling.

---

## ✨ Features

- 🔐 **User Authentication**: Secure login using Paseto tokens, with role-based access (Depositor, Banker).  
- 🏦 **Account Management**: Create, retrieve, and list bank accounts.  
- 💸 **Money Transfer**: Transfer funds between accounts safely.  
- 🌐 **REST & gRPC APIs**: HTTP endpoints (Gin) and gRPC services, with grpc-gateway for HTTP/JSON translation.  
- 🗄️ **Database Migrations**: Automated schema updates via golang-migrate.  
- ✅ **Type-Safe DB Access**: SQL queries compiled to Go code with sqlc.  
- ⚙️ **Background Jobs**: Asynchronous task processing powered by Redis and Asynq.  
- 🧪 **Testing & Mocking**: Comprehensive unit tests and mock generation.  
- 📖 **API Documentation**: Swagger/OpenAPI docs generated from protobuf definitions.  
- 🐳 **Containerized Deployment**: Docker configuration for easy local development and deployment.  

---

## 🛠 Tech Stack

- Go (Golang)  
- Gin (REST API)  
- gRPC / grpc-gateway  
- PostgreSQL  
- Redis  
- sqlc  
- golang-migrate  
- Paseto (token authentication)  
- Docker  
- dbdiagram/dbml  
- ZeroLog  
- Evans (gRPC client/testing)  

---

## 🚀 Getting Started

### Prerequisites
- Go **1.22+**  
- Docker  
- Make  
- PostgreSQL  
- Redis  

### Local Setup

Clone the repository:
```sh
git clone https://github.com/eizyc/GoFinanceServer.git
cd GoFinanceServer
