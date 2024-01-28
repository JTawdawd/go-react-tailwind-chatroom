SELECT 
    u.id, 
    u.username 
FROM 
    user_account u
WHERE 
    u.username = $1 
    AND (
        u.password IS NOT NULL
        AND u.password = crypt($2, u.password)
    )
LIMIT 1