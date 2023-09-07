import React, { useState } from 'react';
import { Route, Navigate, useNavigate, Routes } from 'react-router-dom';
import Login from './Login';
import Lobby from './Lobby';
import Chatroom from './Chatroom';

function LogoutButton({loggedIn, logout}) {
  if (!loggedIn) {
    return;
  }
  return <button onClick={logout}> logout </button>
}

const App = () => {
  const [loggedIn, setLoggedIn] = useState(localStorage.getItem('token') !== null);

  const navigate = useNavigate();

  const logout = () => {
    localStorage.removeItem('token');
    setLoggedIn(false);
    navigate('/')
  }

  return (
    <div className="App">
      <div>
        <LogoutButton loggedIn={loggedIn} logout={logout}/>
      <Routes>
        <Route path="/" element={<Lobby loggedIn={loggedIn} />} />
        <Route path="/Login" element={<Login setLoggedIn={setLoggedIn} />} />
        <Route path="/ChatRoom/:id" element={<Chatroom loggedIn={loggedIn} />} />
      </Routes>
    </div>
    </div>
  );
};

export default App;
