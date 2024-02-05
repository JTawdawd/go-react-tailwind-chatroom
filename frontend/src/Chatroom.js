import React, { useState, useEffect } from 'react';
import { Navigate, useParams, Link } from 'react-router-dom';

const ChatRoom = (loggedIn) => {

    const [content, setContent] = useState('');
    const [messages, setMessages] = useState([]);       
    let { id } = useParams();

    const isOwnMessage = (userId) => {
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
        setMessages(messagesData.messages || []);
        setTimeout(scrollToBottom, 0);
      }
    }

    function scrollToBottom() {
      let container = document.getElementById("chatroom");
      container.scrollTop = container.scrollHeight;
    }

    useEffect(() => {
      if (!loggedIn) {
        return;
      }
      fetchMessages();

      const websocket = new WebSocket(`ws://localhost:8080/websocket/connect?chatroomID=${id}`);

      websocket.onmessage = function(event) {
        console.log('Received message:', event.data);
        if (event.data === 'New message') {
          fetchMessages();
        }
      };

      return () => {
        websocket.close();
      };
    }, [loggedIn]);
    
    if (!loggedIn) {
        return <Navigate to="/login" />;
    }

    function keyDownHandler(event) {
      if (event.key !== 'Enter') {
        return;
      }
      createMessage();
    }

    async function createMessage() {
      const response = await (await fetch(`/api/createMessage`, {
        method: 'POST',
        headers: {
        'Content-Type': 'application/json',
        },
        body: JSON.stringify({ 
          chatroomid: id,
          createdby: localStorage.getItem('userId'),
          content: content,
          createdat: new Date().toISOString()
        })
      })).json();
      if (response.status === 'Success') {
        setContent('');
      }
    }

  return (
    <div>
      <Link to={'/'}>
        <button >Back</button>
      </Link>
      <div id='chatroom'>
        {messages.map((message) => (
            <div className='message' style={isOwnMessage(message.id) ? {alignSelf:  'flex-start', backgroundColor: 'lightblue'} : {alignSelf:  'flex-end', backgroundColor: 'lightgreen'}}>
              <div className='username'>{ message.username }</div>
              <div className='content'>{ message.content }</div>
              <div className='createdat'>{ new Date(message.createdat).toUTCString() }</div>
          </div>
        ))}
      </div>
      <div style={{display: 'flex'}}>
          <input onKeyDown={keyDownHandler} type="textarea" id="messageInput" name="message" value={content} onChange={(e) => setContent(e.target.value)}/>
          <button onClick={createMessage} id="sendButton">Send</button>
        </div>
    </div>
  );
};

export default ChatRoom;