import React, { useState, useEffect } from 'react';
import { Navigate, useParams, Link } from 'react-router-dom';

const ChatRoom = (loggedIn) => {

    const [messages, setMessages] = useState([]);       
    let { id } = useParams();

    const isOwnMessage = (userId) => {
      console.log(`${userId} | ${localStorage.getItem('userId')} | ${userId === localStorage.getItem('userId')}`)
      return userId === localStorage.getItem('userId');
    }

    async function fetchMessages() {
      const response = await fetch(`/api/getChatroom`, {
          method: 'POST',
          headers: {
          'Content-Type': 'application/json',
          },
          body: JSON.stringify({ "id": Number(id) })
      });
      const messagesData = await response.json();
      if (messagesData.status === 'Success') {
        setMessages(messagesData.messages);
      }
    }

    useEffect(() => {
        
        if (loggedIn) {
            fetchMessages();
        }
    }, [loggedIn]);

    if (!loggedIn) {
        return <Navigate to="/login" />;
    }

  return (
    <div>
      <Link to={'/'}>
        <button >Back</button>
      </Link>
      <div className='chatroom'>
        {messages.map((message) => (
            <div className='message' style={isOwnMessage(message.id) ? {alignSelf:  'flex-start', backgroundColor: 'lightblue'} : {alignSelf:  'flex-end', backgroundColor: 'lightgreen'}}>
              <div className='username'>{ message.username }</div>
              <div className='content'>{ message.content }</div>
              <div className='createdat'>{ new Date(message.createdat).toUTCString() }</div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ChatRoom;