import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './../App.css'

const ChatroomPreview = ({ id, title }) => {

  return (
    <div className='bg-white p-4 rounded shadow-md'>
      <h3 className='text-lg font-semibold mb-2'>Chatroom ID: {id}</h3>
      <p className='text-gray-600 mb-4'>Title: {title}</p>
      <Link to={`/chatroom/${id}`}>
        <button className='bg-blue-500 text-white p-2 rounded hover:bg-blue-600 focus:outline-none focus:shadow-outline'>
          Join
        </button>
      </Link>
    </div>
  );
};

export default ChatroomPreview;