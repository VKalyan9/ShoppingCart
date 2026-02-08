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

## Prerequisites

### Backend
- Go 1.21 or higher
- Install from: https://golang.org/dl/

### Frontend
- Node.js 16+ and npm
- Install from: https://nodejs.org/

## Installation & Setup

### Step 1: Setup Backend

1. Open a terminal and navigate to the backend directory:
```bash
cd shopping-cart-project/backend
```

2. Install Go dependencies:
```bash
go mod download
```

3. Start the backend server:
```bash
go run .
```

The backend will start on `http://localhost:8080`

You should see sample items created automatically (Laptop, Mouse, Keyboard, Monitor, Headphones).

### Step 2: Setup Frontend

1. Open a NEW terminal (keep backend running) and navigate to frontend:
```bash
cd shopping-cart-project/frontend
```

2. Install npm dependencies:
```bash
npm install
```

3. Start the React development server:
```bash
npm start
```

The frontend will open automatically at `http://localhost:3000`

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
  ```
  Returns: `{ "token": "...", "user_id": 1, "username": "john" }`

### Items
- `POST /items` - Create item
- `GET /items` - List all items

### Cart (Requires Token Header)
- `POST /carts` - Add item to cart
  ```json
  {
    "item_id": 1
  }
  ```
  Headers: `Token: <your-token>`

- `GET /carts` - Get user's carts
  Headers: `Token: <your-token>`

### Orders (Requires Token Header)
- `POST /orders` - Create order from cart
  Headers: `Token: <your-token>`

- `GET /orders` - Get user's orders
  Headers: `Token: <your-token>`

## Database Schema

### Users Table
- id (primary key)
- username (unique)
- password
- token
- cart_id
- created_at

### Items Table
- id (primary key)
- name
- status
- created_at

### Carts Table
- id (primary key)
- user_id
- name
- status
- created_at

### Orders Table
- id (primary key)
- cart_id
- user_id
- created_at

### Cart_Items Table (Junction)
- cart_id
- item_id

## Testing the API with cURL

### Create a user:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"test123"}'
```

### Login:
```bash
curl -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"test123"}'
```

### Get items:
```bash
curl http://localhost:8080/items
```

### Add to cart (replace YOUR_TOKEN):
```bash
curl -X POST http://localhost:8080/carts \
  -H "Content-Type: application/json" \
  -H "Token: YOUR_TOKEN" \
  -d '{"item_id":1}'
```

### Create order (replace YOUR_TOKEN):
```bash
curl -X POST http://localhost:8080/orders \
  -H "Token: YOUR_TOKEN"
```

## Troubleshooting

### Backend won't start
- Make sure Go is installed: `go version`
- Check if port 8080 is available
- Run `go mod tidy` to fix dependencies

### Frontend won't start
- Make sure Node.js is installed: `node --version`
- Delete `node_modules` and run `npm install` again
- Check if port 3000 is available

### Cannot connect to backend
- Ensure backend is running on port 8080
- Check browser console for CORS errors
- Verify API_URL in Login.js and ItemsList.js is `http://localhost:8080`

### Login not working
- Check backend terminal for errors
- Verify you created a user first (sign up)
- Make sure username and password match

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

## Development Notes

- The backend uses SQLite, so a `shopping_cart.db` file will be created automatically
- Tokens are generated randomly and stored in the database
- One user can only have one active cart at a time
- When you checkout, the cart is converted to an order and a new cart is created for future shopping
- Passwords are stored in plain text (for simplicity - in production, use hashing!)

## Future Enhancements

- Password hashing (bcrypt)
- Item images and prices
- Quantity management
- Payment integration
- User profile management
- Product categories
- Search and filter
- Admin panel

## License

This is a learning project for educational purposes.
