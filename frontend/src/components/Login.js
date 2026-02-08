import React, { useState } from 'react';
import './Login.css';

const API_URL = ''http://localhost:8080'';

function Login({ onLogin }) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isSignup, setIsSignup] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (isSignup) {
      // Sign up
      try {
        const response = await fetch(`${API_URL}/users`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
          alert('User created successfully! Please login.');
          setIsSignup(false);
          setPassword('');
        } else {
          const data = await response.json();
          alert(data.error || 'Signup failed');
        }
      } catch (error) {
        alert('Error: ' + error.message);
      }
    } else {
      // Login
      try {
        const response = await fetch(`${API_URL}/users/login`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
          const data = await response.json();
          onLogin(data.token, data.user_id);
        } else {
          window.alert('Invalid username/password');
        }
      } catch (error) {
        window.alert('Error: ' + error.message);
      }
    }
  };

  return (
    <div className="login-container">
      <div className="login-box">
        <h2>{isSignup ? 'Sign Up' : 'Login'}</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>Username:</label>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label>Password:</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit" className="btn-primary">
            {isSignup ? 'Sign Up' : 'Login'}
          </button>
        </form>
        <p className="toggle-mode">
          {isSignup ? 'Already have an account?' : "Don't have an account?"}
          <button
            className="btn-link"
            onClick={() => {
              setIsSignup(!isSignup);
              setPassword('');
            }}
          >
            {isSignup ? 'Login' : 'Sign Up'}
          </button>
        </p>
      </div>
    </div>
  );
}

export default Login;
