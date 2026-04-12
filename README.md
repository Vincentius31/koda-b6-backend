# Coffee Shop RESTful API

A RESTful API backend for a coffee shop e-commerce platform, built with **Go**, **Gin**, and **PostgreSQL**. It supports product catalog browsing, cart management, order checkout, user authentication, and an admin dashboard.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.25 |
| Framework | Gin v1.12 |
| Database | PostgreSQL (pgx/v5 with connection pool) |
| Auth | JWT (golang-jwt/v5) |
| Password Hashing | Argon2 (matthewhartstonge/argon2) |
| Configuration | godotenv |
| Containerization | Docker (multi-stage build) |

---

## Project Structure

```
koda-b6-backend/
├── cmd/
│   └── main.go                  # Entry point, DB pool init, CORS, server start
├── internal/
│   ├── di/
│   │   └── container.go         # Dependency injection container
│   ├── handlers/                # HTTP request handlers (one file per domain)
│   ├── middleware/
│   │   └── auth.go              # JWT authentication middleware
│   ├── models/                  # Request/response structs and DB models
│   ├── repository/              # Database query layer
│   ├── routes/
│   │   └── routes.go            # Route registration
│   └── service/                 # Business logic layer
├── migrations/                  # SQL migration files (up/down)
├── uploads/
│   └── users/                   # User profile picture uploads
├── .env                         # Environment variables
├── Dockerfile                   # Multi-stage Docker build
└── go.mod
```

---

## Getting Started

### Prerequisites

- Go 1.25+
- PostgreSQL 13+
- (Optional) Docker

### 1. Clone & Configure

```bash
git clone <repo-url>
cd koda-b6-backend
```

Copy and edit the environment file:

```bash
cp .env.example .env
```

### 2. Environment Variables

```env
# Server
PORT=8888

# Database
PGHOST=localhost
PGPORT=5432
PGUSER=postgres
PGPASSWORD=your_password
PGDATABASE=postgres
PGSSLMODE=disable

# JWT Secret
APP_SECRET=your_secret_key

# CORS
FRONTEND_URL=http://localhost:5173
```

### 3. Run Database Migrations

Apply the SQL files in the `migrations/` directory in order:

```bash
psql -U postgres -d postgres -f migrations/000001_init_db.up.sql
psql -U postgres -d postgres -f migrations/000002_forgot_password.up.sql
psql -U postgres -d postgres -f migrations/20260318042053_adding_generic_column.up.sql
psql -U postgres -d postgres -f migrations/20260412063554_init_schema.up.sql
```

### 4. Run the Server

```bash
go mod tidy
go run cmd/main.go
```

The server will start on `http://localhost:8888`.

### 5. Run with Docker

```bash
docker build -t koda-b6-backend .
docker run -p 8888:8888 --env-file .env koda-b6-backend
```

---

## API Reference

### Base URL

```
http://localhost:8888
```

### Authentication

Protected routes require a JWT token in the `Authorization` header:

```
Authorization: Bearer <token>
```

---

### Auth

| Method | Endpoint | Description | Auth |
|---|---|---|---|
| POST | `/auth/register` | Register a new user | No |
| POST | `/auth/login` | Login and receive a JWT token | No |
| POST | `/auth/forgot-password` | Request an OTP for password reset | No |
| PATCH | `/auth/forgot-password` | Reset password using OTP | No |

**Login Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "<jwt_token>",
    "role_id": 1
  }
}
```

---

### Public — Landing & Catalog

| Method | Endpoint | Description | Auth |
|---|---|---|---|
| GET | `/landing/recommended-products` | Fetch recommended products for homepage | No |
| GET | `/landing/reviews` | Fetch latest customer reviews | No |
| GET | `/products` | Get paginated product catalog (with filters) | No |
| GET | `/products/promos` | Get products with active discounts/flash sales | No |
| GET | `/detail-product/:id` | Get full product detail by ID | No |

---

### User (Authenticated)

| Method | Endpoint | Description |
|---|---|---|
| GET | `/profile` | Get logged-in user's profile |
| PATCH | `/profile` | Update profile (name, address, phone, etc.) |
| POST | `/profile/upload` | Upload profile picture |
| GET | `/cart` | Get all items in cart |
| POST | `/cart` | Add item to cart |
| PATCH | `/cart/:id` | Update cart item quantity |
| DELETE | `/cart/:id` | Remove item from cart |
| POST | `/checkout` | Place an order from cart items |
| GET | `/transactions` | Get user's order history |
| GET | `/transactions/:id` | Get detailed view of a specific order |

**Checkout Request Body:**
```json
{
  "delivery_method": "delivery",
  "subtotal": 50000,
  "total": 55000,
  "payment_method": "transfer",
  "items": [
    {
      "product_id": 1,
      "quantity": 2,
      "size": "Large",
      "variant": "Hot",
      "price": 25000
    }
  ]
}
```

---

### Admin (`/admin/*`)

All admin routes require authentication (`Authorization: Bearer <token>`).

#### Dashboard

| Method | Endpoint | Description |
|---|---|---|
| GET | `/admin/dashboard/sales-category` | Sales breakdown by category |
| GET | `/admin/dashboard/best-sellers` | Top-selling products |
| GET | `/admin/dashboard/order-stats` | Order statistics overview |

#### Resource CRUD

Each of the following resources supports full CRUD at `/admin/<resource>`:

| Resource | Base Path |
|---|---|
| Users | `/admin/users` |
| Roles | `/admin/roles` |
| Categories | `/admin/categories` |
| Products | `/admin/product` |
| Product Categories | `/admin/productcategory` |
| Product Images | `/admin/productimage` |
| Product Variants | `/admin/productvariant` |
| Product Sizes | `/admin/productsize` |
| Discounts | `/admin/discount` |
| Transactions | `/admin/transaction` |
| Transaction Products | `/admin/transactionproduct` |
| Reviews | `/admin/review` |

Standard CRUD pattern for each:

| Method | Path | Action |
|---|---|---|
| GET | `/admin/<resource>` | Get all records |
| GET | `/admin/<resource>/:id` | Get single record |
| POST | `/admin/<resource>` | Create new record |
| PATCH | `/admin/<resource>/:id` | Update record |
| DELETE | `/admin/<resource>/:id` | Delete record |

---

## Database Schema

The main tables and their relationships:

```
roles
  └── users (roles_id → roles.id_roles)

products
  ├── products_category (product_id → products.id_product)
  │     └── category (category_id → category.id_category)
  ├── product_images (product_id → products.id_product)
  ├── product_variant (product_id → products.id_product)
  ├── product_size (product_id → products.id_product)
  ├── discount (product_id → products.id_product)
  └── review (product_id → products.id_product)

users
  ├── cart (user_id → users.id_user)
  ├── transaction (user_id → users.id_user)
  └── review (user_id → users.id_user)

transaction
  └── transaction_product (transaction_id → transaction.id_transaction)
```

---

## Standard Response Format

All endpoints return responses in this envelope:

```json
{
  "success": true,
  "message": "Description of the result",
  "data": { ... }
}
```

On error:

```json
{
  "success": false,
  "message": "Error description",
  "data": null
}
```

---

## Static Files

Uploaded user profile pictures are served as static files:

```
GET /uploads/users/<filename>
```

Files are stored on disk under the `uploads/users/` directory.

---

## Connection Pool Configuration

The database connection pool is configured in `cmd/main.go`:

| Setting | Value |
|---|---|
| Max Connections | 20 |
| Min Connections | 5 |
| Max Idle Time | 30 minutes |
| Ping Timeout | 5 seconds |