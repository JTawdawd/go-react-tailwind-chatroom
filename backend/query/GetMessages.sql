SELECT 
    id, content, createdby, createdat 
FROM 
    message 
WHERE 
    chatroomid = $1