import React, { useState, useEffect } from 'react';
import { Navigate, Link } from 'react-router-dom';
import './../App.css';

const CreateChatroom = (loggedIn) => {

    const [title, setTitle] = useState('');

    async function createChatroom() {
        const response = await (await fetch(`/api/createChatroom`, {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json',
            },
            body: JSON.stringify({ 
                title: title
            })
        })).json();
        if (response.status === 'Success') {
            setTitle('');
        }
    }

    useEffect(() => {
      return () => {
      };
    });
    
    if (!loggedIn) {
        return <Navigate to='/login' />;
    }

    return (
        <div className='container mx-auto mt-8 p-4'>
            <form className='bg-white p-4 rounded shadow-md'>
                <label className='block mb-2 font-semibold'>Title</label>
                <input
                    type='text'
                    id='title'
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    className='w-full border p-2 rounded mb-4'
                />

                <button
                    onClick={createChatroom}
                    className='bg-blue-500 text-white p-2 rounded hover:bg-blue-600 focus:outline-none focus:shadow-outline'
                >
                    Create Chatroom
                </button>
            </form>
        </div>
    );
};

export default CreateChatroom;