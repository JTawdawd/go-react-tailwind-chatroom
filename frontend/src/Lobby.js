import React, { useState, useEffect } from 'react';
import { Navigate } from 'react-router-dom';
import ChatroomPreview from './ChatroomPreview';
import './App.css'

const Lobby = ({ loggedIn }) => {
  const [username, setUsername] = useState(localStorage.getItem('token'));
  const [chatrooms, setChatrooms] = useState([]);

  async function fetchChatrooms() {
    const response = await fetch('/api/getChatrooms', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });
    const data = await response.json();
    if (Array.isArray(data.chatrooms) && data.chatrooms.length) {
      setChatrooms(data.chatrooms);
    }
  }

  useEffect(() => {
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
      <div className="chatroomContainer">
        {chatrooms.map((chatroom) => (
          <ChatroomPreview key={chatroom.id} id={chatroom.id} title={chatroom.title} />
        ))}
      </div>
    </div>
  );
};

export default Lobby;