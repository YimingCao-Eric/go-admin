# E-Commerce Admin Backend API (Golang)

A robust Go-based RESTful API backend service built with Fiber framework for managing an e-commerce admin dashboard. This backend provides comprehensive functionality including user authentication, role-based access control (RBAC), product management, order processing, and analytics. The codebase follows (and extends) lessons from the course [React and Golang: A Practical Guide](https://www.udemy.com/course/react-go-admin/).

**Frontend Repository:** [react-admin](https://github.com/YimingCao-Eric/react-admin)

## 🚀 Features

- **Authentication & Authorization**
  - JWT-based authentication with HTTP-only cookies
  - Role-Based Access Control (RBAC)
  - User registration and login
  - Password hashing with bcrypt

- **User Management**
  - User CRUD operations
  - Profile management
  - Password updates
  - Paginated user listing

- **Role & Permission Management**
  - Role CRUD operations
  - Permission management
  - Many-to-many role-permission relationships
  - Fine-grained access control

- **Product Management**
  - Product CRUD operations
  - Paginated product listing
  - Image upload support

- **Order Management**
  - Order listing with pagination
  - Order item tracking
  - Order total calculation
  - CSV export functionality

- **Analytics**
  - Daily sales aggregation
  - Chart data endpoints
  - Sales trend analysis

- **File Management**
  - Image upload endpoint
  - Static file serving

## 📋 Prerequisites

Before running this project, ensure you have the following installed:

- **Go** (version 1.25.0 or higher)
- **MySQL** (version 5.7 or higher)
- **Git** (for cloning the repository)

## 🔧 Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-admin
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Database Setup**
   
   Create a MySQL database:
   ```sql
   CREATE DATABASE go_admin;
   ```

4. **Configure Database Connection**
   
   Update the database connection string in `database/connect.go`:
   ```go
   // Currently configured for: root:fb112358@/go_admin
   // In production, use environment variables instead
   ```
   
   **⚠️ Important:** For production, move database credentials to environment variables.

5. **Configure JWT Secret**
   
   Update the JWT secret in `util/jwt.go`:
   ```go
   // Currently hardcoded as "JWT_SECRET"
   // In production, use environment variables
   ```

## 🏃 Running the Application

1. **Start the server**
   ```bash
   go run main.go
   ```

2. **Server will start on**
   ```
   http://localhost:8000
   ```

3. **API endpoints are available at**
   ```
   http://localhost:8000/api/*
   ```

4. **Start the frontend** (optional)
   
   To use the full admin dashboard, start the React frontend. See the [react-admin repository](https://github.com/YimingCao-Eric/react-admin) for frontend setup instructions.

## 📁 Project Structure

```
go-admin/
├── controllers/          # Request handlers
│   ├── authController.go      # Authentication endpoints
│   ├── userController.go      # User management
│   ├── roleController.go      # Role management
│   ├── permissionController.go # Permission management
│   ├── productController.go   # Product CRUD
│   ├── orderController.go     # Order management & analytics
│   └── imageController.go     # File upload handling
├── database/
│   └── connect.go        # Database connection & migration
├── middlewares/
│   ├── authMiddleware.go      # JWT authentication middleware
│   └── permissionMiddleware.go # RBAC authorization middleware
├── models/              # Data models
│   ├── user.go
│   ├── role.go
│   ├── permission.go
│   ├── product.go
│   ├── order.go
│   ├── entity.go        # Pagination interface
│   └── paginate.go      # Generic pagination utility
├── routes/
│   └── routes.go        # Route definitions
├── util/
│   └── jwt.go          # JWT token utilities
├── uploads/            # Uploaded files directory
├── csv/               # CSV export directory
├── main.go            # Application entry point
└── go.mod             # Go module dependencies
```

## 🔌 API Endpoints

### Authentication (Public)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/register` | Register a new user account |
| POST | `/api/login` | Authenticate user and receive JWT token |

### User Management (Authenticated)

| Method | Endpoint | Description | Permission Required |
|--------|----------|-------------|---------------------|
| GET | `/api/user` | Get current user profile | - |
| PUT | `/api/users/info` | Update current user's info | - |
| PUT | `/api/users/password` | Change current user's password | - |
| POST | `/api/logout` | Logout current user | - |
| GET | `/api/users` | Get paginated user list | `view_users` or `edit_users` |
| POST | `/api/users` | Create a new user | `edit_users` |
| GET | `/api/users/:id` | Get user by ID | `view_users` or `edit_users` |
| PUT | `/api/users/:id` | Update user by ID | `edit_users` |
| DELETE | `/api/users/:id` | Delete user by ID | `edit_users` |

### Role Management (Authenticated)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/roles` | Get all roles with permissions |
| POST | `/api/roles` | Create a new role |
| GET | `/api/roles/:id` | Get role by ID |
| PUT | `/api/roles/:id` | Update role |
| DELETE | `/api/roles/:id` | Delete role |

### Permission Management (Authenticated)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/permissions` | Get all permissions |
| POST | `/api/permissions` | Create a new permission |

### Product Management (Authenticated)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/products` | Get paginated product list |
| POST | `/api/products` | Create a new product |
| GET | `/api/products/:id` | Get product by ID |
| PUT | `/api/products/:id` | Update product |
| DELETE | `/api/products/:id` | Delete product |

### Order Management (Authenticated)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/orders` | Get paginated order list with items |
| POST | `/api/export` | Export orders to CSV |
| GET | `/api/chart` | Get daily sales data for charts |

### File Management (Authenticated)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/upload` | Upload file (multipart form, field: "image") |
| GET | `/api/uploads/*` | Serve uploaded files |

## 🔐 Authentication

The API uses JWT (JSON Web Tokens) for authentication:

1. **Login**: Send credentials to `/api/login`
   - On success, receives JWT token in HTTP-only cookie named `jwt`
   - Token expires after 24 hours
   - Token contains user ID in the issuer claim

2. **Authenticated Requests**: Include JWT token in cookie
   - Middleware validates token automatically
   - Invalid/expired tokens return 401 Unauthorized
   - All routes after `/api/register` and `/api/login` require authentication

3. **Logout**: Send POST request to `/api/logout`
   - Clears authentication cookie
   - Returns success message

## 🛡️ Authorization (RBAC)

The API implements Role-Based Access Control:

- **Permissions**: Granular permissions (e.g., `view_users`, `edit_users`)
- **Roles**: Collections of permissions
- **Users**: Assigned to roles

### Permission Naming Convention

- **GET requests**: Require `view_<resource>` or `edit_<resource>` permission
- **POST/PUT/DELETE requests**: Require `edit_<resource>` permission

Example permissions:
- `view_users`, `edit_users`
- `view_products`, `edit_products`
- `view_orders`, `edit_orders`

## 📊 Pagination

List endpoints support pagination using query parameter `page`:

```
GET /api/products?page=1
GET /api/users?page=2
```

Response format:
```json
{
  "data": [...],
  "meta": {
    "total": 100,
    "page": 1,
    "last_page": 20
  }
}
```

Default: 5 records per page

## 🗄️ Database Schema

### Tables

- **users**: User accounts with authentication
- **roles**: Role definitions
- **permissions**: Permission definitions
- **role_permissions**: Join table (many-to-many)
- **products**: Product catalog
- **orders**: Customer orders
- **order_items**: Order line items

### Auto-Migration

The application automatically migrates schema on startup. Tables are created/updated based on model definitions in the `models/` directory.

## 🔄 Frontend Integration

This backend is designed to work with the React admin frontend:

- **Frontend URL**: `http://localhost:3000` (development)
- **Backend URL**: `http://localhost:8000` (development)
- **CORS**: Configured to allow requests from frontend origin
- **Cookies**: Credentials enabled for JWT token transmission (HTTP-only, 24-hour expiration)
- **API Base Path**: All endpoints prefixed with `/api/`
- **Frontend Repository**: See [react-admin](https://github.com/YimingCao-Eric/react-admin) for frontend setup

## 📦 Dependencies

- **Fiber v3**: Web framework
- **GORM**: ORM for database operations
- **MySQL Driver**: Database driver
- **JWT-Go**: JWT token handling
- **bcrypt**: Password hashing

## 🚀 Deployment

### Production Considerations

1. **Environment Variables**
   - Move database credentials to environment variables
   - Use secure JWT secret from environment
   - Configure CORS origins for production domain

2. **Database**
   - Use production-grade MySQL instance
   - Configure connection pooling
   - Set up database backups

3. **Security**
   - Use HTTPS in production
   - Configure secure cookie settings
   - Implement rate limiting
   - Set up proper logging and monitoring

4. **Build and Run**
   ```bash
   # Build the application
   go build -o admin-server main.go
   
   # Run the binary
   ./admin-server
   ```

## 📝 License

This project is part of an e-commerce admin system.

## 🤝 Contributing

This project follows best practices for Go development. When contributing:

1. Follow the existing code style and comment conventions
2. Ensure proper error handling
3. Test your changes thoroughly
4. Update documentation as needed

## 📚 Learn More

- [Go Documentation](https://go.dev/doc/)
- [Fiber Documentation](https://docs.gofiber.io/)
- [GORM Documentation](https://gorm.io/docs/)
- [JWT Documentation](https://jwt.io/)

---

**Note**: This is the backend service. For the frontend React admin dashboard, visit the [react-admin repository](https://github.com/YimingCao-Eric/react-admin).
