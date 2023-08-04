import React, { useState } from 'react';
import { Link } from 'react-router-dom';

const ChatroomPreview = ({ id, title }) => {

  return (
    <div>
      <h3>Chatroom ID: {id}</h3>
      <p>Title: {title}</p>
      <Link to={`/chatroom/${id}`}>
        <button >Join</button>
      </Link>
    </div>
  );
};

export default ChatroomPreview;