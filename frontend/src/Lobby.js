import React, { useState } from 'react';
import { Navigate } from 'react-router-dom';

const Lobby = ({loggedIn}) => {
  const [username, setUsername] = useState(localStorage.getItem('token'));
  
  if (!loggedIn) {
    return <Navigate to="/login" />;
  }

  return (
    <div>
      <h2>Welcome {username}</h2>
        
    </div>
  );
};

export default Lobby;