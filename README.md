# Shopping Cart Full Stack Application

A complete e-commerce shopping cart application built with Go (Gin + GORM) backend and React frontend.

## Project Structure

```
shopping-cart-project/
├── backend/           # Go backend API
│   ├── main.go
│   ├── models.go
│   ├── handlers.go
│   ├── middleware.go
│   └── go.mod
└── frontend/          # React frontend
    ├── public/
    │   └── index.html
    ├── src/
    │   ├── components/
    │   │   ├── Login.js
    │   │   ├── Login.css
    │   │   ├── ItemsList.js
    │   │   └── ItemsList.css
    │   ├── App.js
    │   ├── App.css
    │   ├── index.js
    │   └── index.css
    └── package.json
```

## Features

### Backend Features
- User authentication with token-based system
- Single active cart per user
- Cart to order conversion
- RESTful API endpoints
- SQLite database with GORM
- CORS enabled for frontend communication

### Frontend Features
- User login/signup
- Items listing
- Add items to cart
- View cart contents
- View order history
- Checkout functionality
- Responsive design

## How to Use the Application

### 1. Create an Account
- Click "Sign Up" on the login page
- Enter a username and password
- Click "Sign Up" button
- You'll see "User created successfully! Please login."

### 2. Login
- Enter your username and password
- Click "Login" button
- On successful login, you'll see the items list

### 3. Shop for Items
- You'll see a list of available items (Laptop, Mouse, Keyboard, etc.)
- Click "Add to Cart" on any item
- You'll get a confirmation alert

### 4. View Your Cart
- Click the "View Cart" button at the top
- An alert will show all items in your cart with their IDs

### 5. View Order History
- Click "Order History" button
- See all your previous orders (if any)

### 6. Checkout
- Click "Checkout" button
- Your cart will be converted to an order
- You'll see "Order successful!" message
- Your cart is now empty and ready for new items

### 7. Logout
- Click "Logout" button to return to login screen

## API Endpoints

### User Management
- `POST /users` - Create new user
  ```json
  {
    "username": "john",
    "password": "password123"
  }
  ```

- `GET /users` - List all users

- `POST /users/login` - Login user
  ```json
  {
    "username": "john",
    "password": "password123"
  }
  
## Technologies Used

### Backend
- **Go** - Programming language
- **Gin** - Web framework
- **GORM** - ORM for database operations
- **SQLite** - Database

### Frontend
- **React** - UI library
- **CSS3** - Styling
- **Fetch API** - HTTP requests



