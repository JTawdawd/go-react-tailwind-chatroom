import React, { useState, useEffect } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';
import ChatroomPreview from './ChatroomPreview';
import CreateChatroom from './CreateChatroom'
import './../App.css';

const Lobby = ({ loggedIn, setLoggedIn }) => {
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

  const navigate = useNavigate();

  function LogoutButton({loggedIn, logout}) {
    if (!loggedIn) {
      return;
    }
    return (
      <button onClick={logout} className='mt-4 bg-blue-500 text-white p-2 rounded hover:bg-blue-600 focus:outline-none focus:shadow-outline' style={{float: 'right'}} >
        Logout
      </button>
    )
  }

  const logout = () => {
    localStorage.removeItem('token');
    setLoggedIn(false);
    navigate('/')
  }

  useEffect(() => {
    if (loggedIn) {
      fetchChatrooms();
    }
  }, [loggedIn]);

  if (!loggedIn) {
    return <Navigate to='/login' />;
  }

  return (
    <div className='container mx-auto mt-8 p-4'>
      <h2 className='text-2xl font-semibold mb-4'>
        Welcome {username}
        <LogoutButton loggedIn={loggedIn} logout={logout} />
      </h2>
      <CreateChatroom loggedIn={loggedIn} />
      <div className='grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-4'>
        {chatrooms.map((chatroom) => (
          <ChatroomPreview
            key={chatroom.id}
            id={chatroom.id}
            title={chatroom.title}
            className='mb-4'
          />
        ))}
      </div>
    </div>
  );
};

export default Lobby;