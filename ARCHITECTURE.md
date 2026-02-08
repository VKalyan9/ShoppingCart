# Architecture Documentation

## System Overview

This is a full-stack e-commerce shopping cart application following a client-server architecture.

```
┌─────────────────┐         HTTP/JSON          ┌─────────────────┐
│                 │ ◄────────────────────────► │                 │
│  React Frontend │         REST API           │   Go Backend    │
│  (Port 3000)    │                            │   (Port 8080)   │
│                 │                            │                 │
└─────────────────┘                            └────────┬────────┘
                                                        │
                                                        │ GORM
                                                        ▼
                                                ┌─────────────────┐
                                                │  SQLite Database│
                                                │ shopping_cart.db│
                                                └─────────────────┘
```

## Technology Stack

### Backend
- **Language:** Go 1.21+
- **Web Framework:** Gin (fast HTTP router)
- **ORM:** GORM (database abstraction)
- **Database:** SQLite (embedded database)
- **Authentication:** Token-based (custom implementation)

### Frontend
- **Library:** React 18
- **Styling:** CSS3
- **HTTP Client:** Fetch API
- **Build Tool:** Create React App

## Database Schema

### Entity Relationship Diagram

```
┌──────────────┐         ┌──────────────┐
│    Users     │         │    Items     │
├──────────────┤         ├──────────────┤
│ id (PK)      │         │ id (PK)      │
│ username     │         │ name         │
│ password     │         │ status       │
│ token        │         │ created_at   │
│ cart_id (FK) │         └──────────────┘
│ created_at   │                │
└──────┬───────┘                │
       │                        │
       │ 1:1                    │
       │                        │
       ▼                        │
┌──────────────┐         ┌──────▼───────┐
│    Carts     │ M:N     │  Cart_Items  │
├──────────────┤◄────────┤──────────────┤
│ id (PK)      │         │ cart_id (FK) │
│ user_id (FK) │         │ item_id (FK) │
│ name         │         └──────────────┘
│ status       │
│ created_at   │
└──────┬───────┘
       │
       │ 1:M
       │
       ▼
┌──────────────┐
│   Orders     │
├──────────────┤
│ id (PK)      │
│ cart_id (FK) │
│ user_id (FK) │
│ created_at   │
└──────────────┘
```

### Relationships
- **User ↔ Cart:** One-to-One (active cart)
- **Cart ↔ Items:** Many-to-Many (through cart_items)
- **User ↔ Orders:** One-to-Many
- **Cart ↔ Orders:** One-to-Many

## API Architecture

### Authentication Flow

```
1. Sign Up
   Client → POST /users {username, password}
   Server → Creates user → Returns user object

2. Login
   Client → POST /users/login {username, password}
   Server → Validates credentials → Generates token → Returns token

3. Authenticated Requests
   Client → Includes "Token" header in all cart/order requests
   Server → Validates token → Identifies user → Processes request
```

### Request/Response Cycle

```
┌────────────┐                          ┌────────────┐
│   Client   │                          │   Server   │
└──────┬─────┘                          └──────┬─────┘
       │                                       │
       │  1. HTTP Request (JSON)               │
       ├──────────────────────────────────────►│
       │     Headers: Token, Content-Type      │
       │     Body: Request Data                │
       │                                       │
       │                              2. Middleware
       │                              ┌────────┴────────┐
       │                              │ CORS            │
       │                              │ Authentication  │
       │                              └────────┬────────┘
       │                                       │
       │                              3. Handler Logic
       │                              ┌────────┴────────┐
       │                              │ Validate Data   │
       │                              │ Database Query  │
       │                              │ Business Logic  │
       │                              └────────┬────────┘
       │                                       │
       │  4. HTTP Response (JSON)              │
       │◄──────────────────────────────────────┤
       │     Status Code                       │
       │     Body: Response Data/Error         │
       │                                       │
```

## Component Architecture

### Backend Components

```
main.go
├── Database Initialization
├── Auto Migration (Create Tables)
├── Seed Data (Sample Items)
├── Router Setup
│   ├── CORS Middleware
│   ├── Public Routes
│   │   ├── POST /users
│   │   ├── GET /users
│   │   ├── POST /users/login
│   │   ├── POST /items
│   │   └── GET /items
│   └── Protected Routes (with authMiddleware)
│       ├── POST /carts
│       ├── GET /carts
│       ├── POST /orders
│       └── GET /orders
└── Start Server (Port 8080)

models.go
├── User Model
├── Item Model
├── Cart Model
├── Order Model
└── CartItem Model (Junction)

handlers.go
├── User Handlers
│   ├── createUser()
│   ├── getUsers()
│   └── loginUser()
├── Item Handlers
│   ├── createItem()
│   └── getItems()
├── Cart Handlers
│   ├── addToCart()
│   └── getCarts()
└── Order Handlers
    ├── createOrder()
    └── getOrders()

middleware.go
└── authMiddleware()
    ├── Extract Token from Header
    ├── Validate Token in Database
    ├── Attach User ID to Context
    └── Allow/Deny Request
```

### Frontend Components

```
App.js (Root Component)
├── State Management
│   ├── isLoggedIn
│   ├── token
│   └── userId
├── Conditional Rendering
│   ├── Login Component (when logged out)
│   └── ItemsList Component (when logged in)
└── Handler Functions
    ├── handleLogin()
    └── handleLogout()

Login.js
├── State
│   ├── username
│   ├── password
│   └── isSignup
├── UI
│   ├── Username Input
│   ├── Password Input
│   └── Submit Button
└── Functions
    ├── handleSubmit()
    │   ├── Sign Up Flow
    │   └── Login Flow
    └── Toggle Sign Up/Login

ItemsList.js
├── State
│   ├── items[]
│   └── cartItems[]
├── Effects
│   └── useEffect() → fetchItems()
├── UI
│   ├── Header
│   │   ├── View Cart Button
│   │   ├── Order History Button
│   │   ├── Checkout Button
│   │   └── Logout Button
│   └── Items Grid
│       └── Item Cards
│           ├── Item Name
│           ├── Item ID
│           └── Add to Cart Button
└── Functions
    ├── fetchItems()
    ├── addToCart(itemId)
    ├── viewCart()
    ├── viewOrderHistory()
    └── checkout()
```

## Data Flow

### Adding Item to Cart

```
1. User Action
   User clicks "Add to Cart" button
   ↓
2. Frontend
   ItemsList.js → addToCart(itemId)
   ↓
3. HTTP Request
   POST /carts
   Headers: { Token: user-token }
   Body: { item_id: 1 }
   ↓
4. Backend Middleware
   authMiddleware() validates token
   Extracts user_id from database
   Attaches user_id to context
   ↓
5. Backend Handler
   addToCart(c *gin.Context)
   ├── Get user_id from context
   ├── Find or create active cart for user
   ├── Verify item exists
   ├── Add item to cart (cart_items table)
   └── Return updated cart
   ↓
6. Database
   ┌─────────────────────────────────┐
   │ Transaction:                    │
   │ 1. SELECT cart WHERE user_id    │
   │ 2. INSERT INTO cart (if needed) │
   │ 3. INSERT INTO cart_items       │
   │ 4. SELECT cart WITH items       │
   └─────────────────────────────────┘
   ↓
7. HTTP Response
   { id: 1, user_id: 1, items: [...] }
   ↓
8. Frontend Update
   Show success alert
   Update local cartItems state
```

### Checkout Flow

```
1. User Action
   User clicks "Checkout" button
   ↓
2. Frontend
   ItemsList.js → checkout()
   ↓
3. HTTP Request
   POST /orders
   Headers: { Token: user-token }
   ↓
4. Backend Handler
   createOrder(c *gin.Context)
   ├── Get user_id from context
   ├── Find active cart for user
   ├── Create order record
   │   └── Links cart_id and user_id
   ├── Update cart status to "ordered"
   └── Clear user's cart_id reference
   ↓
5. Database
   ┌─────────────────────────────────┐
   │ Transaction:                    │
   │ 1. SELECT cart WHERE active     │
   │ 2. INSERT INTO orders           │
   │ 3. UPDATE cart SET status       │
   │ 4. UPDATE user SET cart_id NULL │
   └─────────────────────────────────┘
   ↓
6. HTTP Response
   { id: 1, cart_id: 1, user_id: 1 }
   ↓
7. Frontend Update
   Show "Order successful!" alert
   Clear local cart state
   User can start shopping again
```

## Security Considerations

### Current Implementation
- Token-based authentication
- Single session per user (one token at a time)
- CORS enabled for localhost development
- User ownership verified for cart/order operations

### Production Improvements Needed
- **Password Hashing:** Use bcrypt instead of plain text
- **HTTPS:** Encrypt data in transit
- **Token Expiration:** Add TTL to tokens
- **JWT:** Use industry-standard JSON Web Tokens
- **Rate Limiting:** Prevent brute force attacks
- **Input Validation:** Sanitize all user inputs
- **SQL Injection:** GORM provides protection, but always validate
- **Environment Variables:** Store config outside code

## Performance Considerations

### Current Optimizations
- SQLite for fast local queries
- Gin framework (fast HTTP routing)
- GORM query optimization
- React component-level state

### Scalability Improvements
- **Database:** Switch to PostgreSQL/MySQL for production
- **Caching:** Add Redis for session/cart caching
- **Load Balancing:** Multiple backend instances
- **CDN:** Serve static frontend assets
- **Database Indexing:** Add indexes on foreign keys
- **Pagination:** Limit results for large datasets

## Testing Strategy

### Backend Testing (with Ginkgo)
```
handlers_test.go
├── User Tests
│   ├── Create User Success
│   ├── Create User Duplicate
│   ├── Login Success
│   └── Login Failure
├── Cart Tests
│   ├── Add Item to Cart
│   ├── View Cart
│   └── Auth Required
└── Order Tests
    ├── Create Order
    └── Get Orders
```

### Frontend Testing
```
- Component rendering tests
- User interaction tests
- API integration tests
- Error handling tests
```

## Deployment

### Backend Deployment
```
1. Build: go build -o shopping-cart-api
2. Run: ./shopping-cart-api
3. Ensure: Port 8080 accessible
```

### Frontend Deployment
```
1. Build: npm run build
2. Serve: Static files in build/
3. Configure: API_URL environment variable
```

### Full Deployment
```
Option 1: Same Server
Frontend (Nginx) ↔ Backend (Go) ↔ Database

Option 2: Separate Servers
Frontend (Vercel/Netlify) → Backend (Heroku/Railway) → Database (Managed)

Option 3: Containers
Docker → Frontend + Backend + Database
```

## Future Enhancements

1. **Features**
   - Product images and prices
   - Quantity management
   - Wishlist
   - Product search/filter
   - Payment integration
   - Email notifications

2. **Architecture**
   - Microservices (User, Product, Order services)
   - Message queue (RabbitMQ/Kafka)
   - Event-driven architecture
   - GraphQL API

3. **UI/UX**
   - Mobile app (React Native)
   - Admin dashboard
   - Real-time updates (WebSockets)
   - Progressive Web App (PWA)

## Conclusion

This architecture provides a solid foundation for an e-commerce application with:
- Clean separation of concerns
- RESTful API design
- Scalable database schema
- Extensible component structure
- Clear data flow patterns
