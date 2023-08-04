import React, { useState, useEffect } from 'react';
import { Navigate } from 'react-router-dom';
import ChatroomPreview from './ChatroomPreview';

const Lobby = ({ loggedIn }) => {
  const [username, setUsername] = useState(localStorage.getItem('token'));
  const [chatrooms, setChatrooms] = useState([]);

  useEffect(() => {
    async function fetchChatrooms() {
      const response = await fetch('http://localhost:8080/chatrooms', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      const chatroomsData = await response.json();
      setChatrooms(chatroomsData);
    }

    if (loggedIn) {
      fetchChatrooms();
    }
  }, [loggedIn]);

  if (!loggedIn) {
    return <Navigate to="/login" />;
  }

  return (
    <div>
      <h2>Welcome {username}</h2>
      <div>
        {chatrooms.map((chatroom) => (
          <ChatroomPreview key={chatroom.id} id={chatroom.id} title={chatroom.title} />
        ))}
      </div>
    </div>
  );
};

export default Lobby;