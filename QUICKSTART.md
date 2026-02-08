# Quick Start Guide

## For Complete Beginners

### What You Need to Install

1. **Install Go** (for backend)
   - Visit: https://golang.org/dl/
   - Download for your OS (Windows/Mac/Linux)
   - Follow installation wizard
   - Verify: Open terminal and type `go version`

2. **Install Node.js** (for frontend)
   - Visit: https://nodejs.org/
   - Download LTS version
   - Follow installation wizard
   - Verify: Open terminal and type `node --version`

### Running the Application

#### Terminal 1 - Backend
```bash
# Navigate to backend folder
cd shopping-cart-project/backend

# Download dependencies (first time only)
go mod download

# Start the server
go run .
```

Keep this terminal running! You should see the server start on port 8080.

#### Terminal 2 - Frontend
```bash
# Navigate to frontend folder
cd shopping-cart-project/frontend

# Install dependencies (first time only)
npm install

# Start the app
npm start
```

The app will open in your browser automatically at http://localhost:3000

### First Time Usage

1. **Sign Up**
   - Click "Sign Up" button
   - Enter username: `demo`
   - Enter password: `demo123`
   - Click "Sign Up"
   - You'll see success message

2. **Login**
   - Enter username: `demo`
   - Enter password: `demo123`
   - Click "Login"

3. **Start Shopping!**
   - You'll see 5 sample items
   - Click "Add to Cart" on any item
   - Click "View Cart" to see your items
   - Click "Checkout" to place order
   - Click "Order History" to see your orders

### Common Issues

**"Command not found: go"**
- Go is not installed or not in PATH
- Reinstall Go and restart terminal

**"Command not found: npm"**
- Node.js is not installed
- Install Node.js from nodejs.org

**"Port 8080 already in use"**
- Another program is using port 8080
- Stop that program or change port in backend/main.go

**"Port 3000 already in use"**
- Another React app is running
- Stop it or the terminal will ask if you want to use a different port

**Cannot connect to backend**
- Make sure backend terminal is still running
- Check that backend is on port 8080
- Check browser console for errors

### Understanding the Flow

```
User Flow:
1. Sign Up → Creates account
2. Login → Gets authentication token
3. Browse Items → See all products
4. Add to Cart → Items go into your cart
5. View Cart → See what you've selected
6. Checkout → Cart becomes an order
7. Order History → See all past orders
```

### File Locations

**Backend files:**
- `backend/main.go` - Main server file
- `backend/models.go` - Database models
- `backend/handlers.go` - API logic
- `shopping_cart.db` - Database (created automatically)

**Frontend files:**
- `frontend/src/App.js` - Main React component
- `frontend/src/components/Login.js` - Login screen
- `frontend/src/components/ItemsList.js` - Shopping screen

### Need Help?

1. Check both terminals for error messages
2. Make sure both servers are running
3. Try restarting both servers
4. Delete `shopping_cart.db` for fresh start
5. Delete `node_modules` and run `npm install` again

### What Each Terminal Should Show

**Backend Terminal (Terminal 1):**
```
[GIN-debug] POST   /users
[GIN-debug] GET    /users
[GIN-debug] POST   /users/login
[GIN-debug] POST   /items
[GIN-debug] GET    /items
[GIN-debug] POST   /carts
[GIN-debug] GET    /carts
[GIN-debug] POST   /orders
[GIN-debug] GET    /orders
[GIN-debug] Listening and serving HTTP on :8080
```

**Frontend Terminal (Terminal 2):**
```
Compiled successfully!

You can now view shopping-cart-frontend in the browser.

  Local:            http://localhost:3000
```

## Done! 🎉

You now have a working shopping cart application!
