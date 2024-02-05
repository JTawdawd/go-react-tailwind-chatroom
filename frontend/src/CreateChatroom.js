import React, { useState, useEffect } from 'react';
import { Navigate, Link } from 'react-router-dom';

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
        return <Navigate to="/login" />;
    }

    return (
        <div>
            <Link to={'/'}>
                <button> Back </button>
            </Link>
            <form>

                <label> Title </label>
                <input type ="text" id ="title" value={title} onChange={(e) => setTitle(e.target.value)} />

                <button onClick={createChatroom}> Create Chatroom</button>
            </form>
        </div>
    );
};

export default CreateChatroom;