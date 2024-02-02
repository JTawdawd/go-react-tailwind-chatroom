SELECT
	u.username,
    u.id,
	m.content,
	m.createdat
FROM 
	message m
	JOIN user_account u on (u.id = m.createdby)
WHERE
	m.chatroomid = $1
ORDER BY m.createdat ASC