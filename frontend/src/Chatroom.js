import React, { useState, useEffect } from 'react';
import { Navigate, useParams } from 'react-router-dom';

const ChatRoom = (loggedIn) => {

    const [messages, setMessages] = useState([]);       
    let { id } = useParams();

    useEffect(() => {
        async function fetchMessages() {
        const response = await fetch(`http://localhost:8080/chatroom`, {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json',
            },
            body: JSON.stringify({ "id": Number(id) })
        });
        const messagesData = await response.json();
        setMessages(messagesData);
        }

        if (loggedIn) {
            fetchMessages();
        }
    }, [loggedIn]);

    if (!loggedIn) {
        return <Navigate to="/login" />;
    }

  return (
    <div>
      <div>
        {messages.map((message) => (
            <div>
                <p>{message.content}</p>
                <p>{message.createdby}</p>
                <p>{message.createdat}</p>
            </div>
        ))}
      </div>
    </div>
  );
};

export default ChatRoom;