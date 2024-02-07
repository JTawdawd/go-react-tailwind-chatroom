import React, { useState } from 'react';
import { Route, Routes } from 'react-router-dom';
import Login from './components/Login';
import Lobby from './components/Lobby';
import Chatroom from './components/Chatroom';

const App = () => {
  const [loggedIn, setLoggedIn] = useState(localStorage.getItem('token') !== null);

  return (
    <div className="App">
      <div>
        <Routes>
          <Route path="/" element={<Lobby loggedIn={loggedIn} setLoggedIn={setLoggedIn} />} />
          <Route path="/Login" element={<Login setLoggedIn={setLoggedIn} />} />
          <Route path="/ChatRoom/:id" element={<Chatroom loggedIn={loggedIn} />} />
        </Routes>
      </div>
    </div>
  );
};

export default App;
