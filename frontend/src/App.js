import React, { useState } from 'react';
import { Route, Navigate, Routes } from 'react-router-dom';
import Login from './Login';
import Lobby from './Lobby'

function LogoutButton({loggedIn, logout}) {
  if (!loggedIn) {
    return;
  }
  return <button onClick={logout}> logout </button>
}

const App = () => {
  const [loggedIn, setLoggedIn] = useState(localStorage.getItem('token') !== null);

  const logout = () => {
    localStorage.removeItem('token');
    setLoggedIn(false);
  }

  return (
    <div className="App">
      <div>
        <LogoutButton loggedIn={loggedIn} logout={logout}/>
      <Routes>
        <Route path="/" element={<Lobby loggedIn={loggedIn} />} />
        <Route path="/Login" element={<Login setLoggedIn={setLoggedIn} />} />
      </Routes>
    </div>
    </div>
  );
};

export default App;
