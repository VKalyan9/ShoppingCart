import React, { useState, useEffect } from 'react';
import './ItemsList.css';

const API_URL = 'http://localhost:8080';

function ItemsList({ token, userId, onLogout }) {
  const [items, setItems] = useState([]);
  const [cartItems, setCartItems] = useState([]);

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      const response = await fetch(`${API_URL}/items`);
      const data = await response.json();
      setItems(data);
    } catch (error) {
      alert('Error fetching items: ' + error.message);
    }
  };

  const addToCart = async (itemId) => {
    try {
      const response = await fetch(`${API_URL}/carts`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Token': token,
        },
        body: JSON.stringify({ item_id: itemId }),
      });

      if (response.ok) {
        const data = await response.json();
        alert(`Item added to cart!`);
        // Keep track of cart items locally
        const item = items.find(i => i.id === itemId);
        setCartItems([...cartItems, { cart_id: data.id, item_id: itemId, item_name: item.name }]);
      } else {
        alert('Failed to add item to cart');
      }
    } catch (error) {
      alert('Error: ' + error.message);
    }
  };

  const viewCart = async () => {
    try {
      const response = await fetch(`${API_URL}/carts`, {
        headers: {
          'Token': token,
        },
      });

      if (response.ok) {
        const data = await response.json();
        if (data.length === 0) {
          window.alert('Your cart is empty');
          return;
        }

        let cartInfo = 'Cart Items:\n\n';
        data.forEach((cart) => {
          if (cart.items && cart.items.length > 0) {
            cart.items.forEach((item) => {
              cartInfo += `Cart ID: ${cart.id}, Item ID: ${item.id}, Item Name: ${item.name}\n`;
            });
          }
        });

        if (cartInfo === 'Cart Items:\n\n') {
          window.alert('Your cart is empty');
        } else {
          window.alert(cartInfo);
        }
      }
    } catch (error) {
      alert('Error: ' + error.message);
    }
  };

  const viewOrderHistory = async () => {
    try {
      const response = await fetch(`${API_URL}/orders`, {
        headers: {
          'Token': token,
        },
      });

      if (response.ok) {
        const data = await response.json();
        if (data.length === 0) {
          window.alert('No orders yet');
          return;
        }

        let orderInfo = 'Order History:\n\n';
        data.forEach((order) => {
          orderInfo += `Order ID: ${order.id}, Cart ID: ${order.cart_id}, Date: ${new Date(order.created_at).toLocaleDateString()}\n`;
        });

        window.alert(orderInfo);
      }
    } catch (error) {
      alert('Error: ' + error.message);
    }
  };

  const checkout = async () => {
    try {
      const response = await fetch(`${API_URL}/orders`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Token': token,
        },
      });

      if (response.ok) {
        alert('Order successful!');
        setCartItems([]);
      } else {
        const data = await response.json();
        alert(data.error || 'Checkout failed');
      }
    } catch (error) {
      alert('Error: ' + error.message);
    }
  };

  return (
    <div className="items-container">
      <div className="header">
        <h1>Shopping Portal</h1>
        <div className="header-buttons">
          <button onClick={viewCart} className="btn-secondary">
            View Cart
          </button>
          <button onClick={viewOrderHistory} className="btn-secondary">
            Order History
          </button>
          <button onClick={checkout} className="btn-primary">
            Checkout
          </button>
          <button onClick={onLogout} className="btn-logout">
            Logout
          </button>
        </div>
      </div>

      <div className="items-grid">
        <h2>Available Items</h2>
        <div className="items-list">
          {items.map((item) => (
            <div key={item.id} className="item-card">
              <h3>{item.name}</h3>
              <p>ID: {item.id}</p>
              <button
                onClick={() => addToCart(item.id)}
                className="btn-add-cart"
              >
                Add to Cart
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default ItemsList;
