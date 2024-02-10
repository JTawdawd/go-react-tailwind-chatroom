import React, { useState, useEffect } from 'react';
import { Navigate, useParams, Link } from 'react-router-dom';
import './../App.css';

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
          body: JSON.stringify({ 'id': Number(id) })
      });
      const messagesData = await response.json();
      if (messagesData.status === 'Success') {
        setMessages(messagesData.messages || []);
        setTimeout(scrollToBottom, 0);
      }
    }

    function scrollToBottom() {
      let container = document.getElementById('chatroom');
      container.scrollTop = container.scrollHeight;
    }

    function newMessage(newMessage) {
      if (!newMessage) {
        return;
      }
      messages.push({...newMessage, id: newMessage.createdby})
      setMessages(messages)
    };

    useEffect(() => {
      if (!loggedIn) {
        return;
      }
      fetchMessages();

      const websocket = new WebSocket(`ws://localhost:8080/websocket/connect?chatroomID=${id}`);

      websocket.onmessage = (event) => {
       newMessage(JSON.parse(event.data))
      };

      return () => {
        websocket.close();
      };
    }, [loggedIn]);
    
    if (!loggedIn) {
        return <Navigate to='/login' />;
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
    <div className='container mx-auto mt-8 p-4'>
      <Link to={'/'}>
        <button className='mb-4 bg-blue-500 text-white p-2 rounded hover:bg-blue-600 focus:outline-none focus:shadow-outline'>
          Back
        </button>
      </Link>
      <div id='chatroom' className='bg-white p-4 rounded shadow-md mb-4 flex' style={{ flexDirection: 'column', height: '80vh', overflowY: 'scroll'}}>
        <p className='text-sm mb-1'>This is the start of your chat</p>
        {messages.map((message) => (
          <div
            key={message.id}
            className={`p-3 mb-2 rounded-lg ${
              isOwnMessage(message.id) ? 'bg-blue-300 self-start' : 'bg-green-300 self-end'
            }`}
          >
            <div className='font-semibold text-sm mb-1'>{message.username}</div>
            <div className='text-sm mb-1'>{message.content}</div>
            <div className='text-xs text-gray-600'>
              {new Date(message.createdat).toUTCString()}
            </div>
          </div>
        ))}
      </div>
      <div className='flex'>
        <input
          onKeyDown={keyDownHandler}
          type='textarea'
          id='messageInput'
          name='message'
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className='w-full border p-2 rounded mr-2'
        />
        <button
          onClick={createMessage}
          id='sendButton'
          className='bg-blue-500 text-white p-2 rounded hover:bg-blue-600 focus:outline-none focus:shadow-outline'
        >
          Send
        </button>
      </div>
    </div>
  );
};

export default ChatRoom;