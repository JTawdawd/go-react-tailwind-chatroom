SELECT 
    u.username
FROM 
    user_account u
WHERE 
    u.id = $1 
LIMIT 1