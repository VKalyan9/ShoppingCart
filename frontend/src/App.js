import React, { useState } from 'react';
import Login from './components/Login';
import ItemsList from './components/ItemsList';
import './App.css';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [token, setToken] = useState('');
  const [userId, setUserId] = useState('');

  const handleLogin = (authToken, id) => {
    setToken(authToken);
    setUserId(id);
    setIsLoggedIn(true);
  };

  const handleLogout = () => {
    setToken('');
    setUserId('');
    setIsLoggedIn(false);
  };

  return (
    <div className="App">
      {!isLoggedIn ? (
        <Login onLogin={handleLogin} />
      ) : (
        <ItemsList token={token} userId={userId} onLogout={handleLogout} />
      )}
    </div>
  );
}

export default App;
