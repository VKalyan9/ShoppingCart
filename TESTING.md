# Testing Guide

## Manual Testing Checklist

### 1. User Sign Up
- [ ] Open http://localhost:3000
- [ ] Click "Sign Up" button
- [ ] Enter username: `testuser1`
- [ ] Enter password: `password123`
- [ ] Click "Sign Up"
- [ ] **Expected:** Alert "User created successfully! Please login."
- [ ] **Expected:** Form switches to Login mode

### 2. User Login - Success
- [ ] Enter username: `testuser1`
- [ ] Enter password: `password123`
- [ ] Click "Login"
- [ ] **Expected:** Redirects to Items List page
- [ ] **Expected:** See 5 sample items (Laptop, Mouse, etc.)

### 3. User Login - Failure
- [ ] Logout and return to login
- [ ] Enter username: `wronguser`
- [ ] Enter password: `wrongpass`
- [ ] Click "Login"
- [ ] **Expected:** Alert "Invalid username/password"
- [ ] **Expected:** Stays on login page

### 4. Add Items to Cart
- [ ] Login successfully
- [ ] Click "Add to Cart" on Laptop
- [ ] **Expected:** Alert "Item added to cart!"
- [ ] Click "Add to Cart" on Mouse
- [ ] **Expected:** Alert "Item added to cart!"
- [ ] Click "Add to Cart" on Keyboard
- [ ] **Expected:** Alert "Item added to cart!"

### 5. View Cart
- [ ] Click "View Cart" button
- [ ] **Expected:** Alert showing:
  ```
  Cart Items:
  
  Cart ID: 1, Item ID: 1, Item Name: Laptop
  Cart ID: 1, Item ID: 2, Item Name: Mouse
  Cart ID: 1, Item ID: 3, Item Name: Keyboard
  ```

### 6. View Empty Cart
- [ ] Logout and login as new user
- [ ] Click "View Cart" before adding items
- [ ] **Expected:** Alert "Your cart is empty"

### 7. View Order History - Empty
- [ ] Click "Order History" button (as new user)
- [ ] **Expected:** Alert "No orders yet"

### 8. Checkout
- [ ] Add 2-3 items to cart
- [ ] Click "View Cart" to confirm items
- [ ] Click "Checkout" button
- [ ] **Expected:** Alert "Order successful!"
- [ ] Click "View Cart" again
- [ ] **Expected:** Alert "Your cart is empty" (cart cleared)

### 9. View Order History - With Orders
- [ ] Click "Order History" button
- [ ] **Expected:** Alert showing order details:
  ```
  Order History:
  
  Order ID: 1, Cart ID: 1, Date: 2/7/2026
  ```

### 10. Multiple Shopping Sessions
- [ ] Add items to cart
- [ ] Checkout
- [ ] Add different items to cart
- [ ] Checkout again
- [ ] Click "Order History"
- [ ] **Expected:** See multiple orders listed

### 11. Logout and Re-login
- [ ] Add items to cart
- [ ] Click "Logout"
- [ ] Login with same credentials
- [ ] Click "View Cart"
- [ ] **Expected:** Cart persists with same items
- [ ] **Note:** This tests cart persistence

### 12. Multiple Users
- [ ] Logout
- [ ] Create user: `user2` / `pass2`
- [ ] Login as `user2`
- [ ] Add items to cart
- [ ] Logout
- [ ] Login as first user
- [ ] Click "View Cart"
- [ ] **Expected:** Only see first user's cart items
- [ ] **Note:** This tests user isolation

## API Testing with cURL

### Test User Creation
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "apitest",
    "password": "test123"
  }'
```
**Expected Response:**
```json
{
  "id": 1,
  "username": "apitest",
  "cart_id": null,
  "created_at": "2026-02-07T10:30:00Z"
}
```

### Test Login
```bash
curl -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "apitest",
    "password": "test123"
  }'
```
**Expected Response:**
```json
{
  "token": "a1b2c3d4e5f6...",
  "user_id": 1,
  "username": "apitest"
}
```
**Save the token for next requests!**

### Test Get Items
```bash
curl http://localhost:8080/items
```
**Expected Response:**
```json
[
  {
    "id": 1,
    "name": "Laptop",
    "status": "active",
    "created_at": "2026-02-07T10:00:00Z"
  },
  ...
]
```

### Test Add to Cart
```bash
# Replace YOUR_TOKEN with the token from login
curl -X POST http://localhost:8080/carts \
  -H "Content-Type: application/json" \
  -H "Token: YOUR_TOKEN" \
  -d '{
    "item_id": 1
  }'
```
**Expected Response:**
```json
{
  "id": 1,
  "user_id": 1,
  "name": "Shopping Cart",
  "status": "active",
  "items": [
    {
      "id": 1,
      "name": "Laptop",
      "status": "active"
    }
  ],
  "created_at": "2026-02-07T10:35:00Z"
}
```

### Test Get Carts
```bash
curl http://localhost:8080/carts \
  -H "Token: YOUR_TOKEN"
```

### Test Create Order
```bash
curl -X POST http://localhost:8080/orders \
  -H "Token: YOUR_TOKEN"
```
**Expected Response:**
```json
{
  "id": 1,
  "cart_id": 1,
  "user_id": 1,
  "created_at": "2026-02-07T10:40:00Z"
}
```

### Test Get Orders
```bash
curl http://localhost:8080/orders \
  -H "Token: YOUR_TOKEN"
```

### Test Auth Required (Should Fail)
```bash
curl -X POST http://localhost:8080/carts \
  -H "Content-Type: application/json" \
  -d '{"item_id": 1}'
```
**Expected Response:**
```json
{
  "error": "Token required"
}
```

## Edge Cases to Test

### 1. Duplicate Username
- [ ] Create user with username `duplicate`
- [ ] Try to create another user with username `duplicate`
- [ ] **Expected:** Error message

### 2. Invalid Token
- [ ] Try to access `/carts` with invalid token
- [ ] **Expected:** 401 Unauthorized

### 3. Add Non-existent Item
```bash
curl -X POST http://localhost:8080/carts \
  -H "Content-Type: application/json" \
  -H "Token: YOUR_TOKEN" \
  -d '{"item_id": 9999}'
```
- [ ] **Expected:** Error "Item not found"

### 4. Checkout Empty Cart
- [ ] Login as new user (no items added)
- [ ] Click "Checkout"
- [ ] **Expected:** Error "No active cart found"

### 5. Multiple Same Items
- [ ] Add same item multiple times
- [ ] Check if duplicates appear in cart
- [ ] **Note:** Current implementation allows duplicates

## Browser Console Testing

### Check Network Requests
1. Open browser DevTools (F12)
2. Go to "Network" tab
3. Perform any action (login, add to cart, etc.)
4. Verify:
   - [ ] Request method (GET/POST)
   - [ ] Request headers (Token)
   - [ ] Request payload
   - [ ] Response status (200, 201, 401, etc.)
   - [ ] Response data

### Check Console Errors
1. Open "Console" tab
2. Perform all actions
3. Verify:
   - [ ] No JavaScript errors
   - [ ] No CORS errors
   - [ ] No network errors

## Database Verification

### View Database Contents
```bash
# Install SQLite browser or use command line
sqlite3 shopping_cart.db

# List all users
SELECT * FROM users;

# List all carts
SELECT * FROM carts;

# List all cart items
SELECT * FROM cart_items;

# List all orders
SELECT * FROM orders;

# Exit
.exit
```

### Verify Data Integrity
```sql
-- Check if user has cart
SELECT u.username, c.id as cart_id 
FROM users u 
LEFT JOIN carts c ON u.cart_id = c.id;

-- Check cart items
SELECT c.id as cart_id, i.name as item_name 
FROM carts c 
JOIN cart_items ci ON c.id = ci.cart_id 
JOIN items i ON ci.item_id = i.id;

-- Check orders
SELECT o.id as order_id, u.username, c.status as cart_status 
FROM orders o 
JOIN users u ON o.user_id = u.id 
JOIN carts c ON o.cart_id = c.id;
```

## Performance Testing

### Load Test with Multiple Requests
```bash
# Add 10 items quickly
for i in {1..10}; do
  curl -X POST http://localhost:8080/carts \
    -H "Content-Type: application/json" \
    -H "Token: YOUR_TOKEN" \
    -d "{\"item_id\": 1}"
done
```

### Concurrent Users Test
```bash
# Create 5 users simultaneously
for i in {1..5}; do
  curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d "{\"username\": \"user$i\", \"password\": \"pass$i\"}" &
done
wait
```

## Automated Testing (Optional)

### Backend Tests with Ginkgo
```bash
cd backend

# Install ginkgo
go install github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...

# Run tests
ginkgo -r
```

### Frontend Tests with Jest
```bash
cd frontend

# Run tests
npm test

# Run with coverage
npm test -- --coverage
```

## Bug Report Template

If you find a bug, report it with:

```
Title: [Brief description]

Steps to Reproduce:
1. 
2. 
3. 

Expected Behavior:
[What should happen]

Actual Behavior:
[What actually happened]

Environment:
- OS: [Windows/Mac/Linux]
- Browser: [Chrome/Firefox/Safari]
- Go Version: [run: go version]
- Node Version: [run: node --version]

Screenshots/Logs:
[If applicable]
```

## Test Results Checklist

After completing all tests:

- [ ] All user flows work correctly
- [ ] Authentication works properly
- [ ] Cart operations function as expected
- [ ] Order creation successful
- [ ] Data persists correctly
- [ ] Error handling works
- [ ] No console errors
- [ ] API responses are correct
- [ ] Database integrity maintained
- [ ] Multi-user isolation works

## Common Issues and Solutions

| Issue | Solution |
|-------|----------|
| Token not working | Check if token is correctly sent in header |
| Cart not persisting | Verify database cart_id in users table |
| Items not showing | Check if backend seeded items on startup |
| Login fails | Verify username/password match in database |
| CORS errors | Ensure backend CORS middleware is enabled |
| Port already in use | Kill process or use different port |

## Testing Frequency

- **Before deployment:** Run all tests
- **After code changes:** Run affected tests
- **Weekly:** Full regression testing
- **After bug fixes:** Test the specific scenario

## Success Criteria

✅ All manual tests pass
✅ All API endpoints return correct responses
✅ No errors in browser console
✅ Database data is consistent
✅ Multiple users can use system simultaneously
✅ Data persists across sessions
✅ Error messages are clear and helpful
