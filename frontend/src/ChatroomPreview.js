import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './App.css'

const ChatroomPreview = ({ id, title }) => {

  return (
    <div className="chatroomPreview">
      <h3>Chatroom ID: {id}</h3>
      <p>Title: {title}</p>
      <Link to={`/chatroom/${id}`}>
        <button >Join</button>
      </Link>
    </div>
  );
};

export default ChatroomPreview;