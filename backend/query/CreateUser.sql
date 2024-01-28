INSERT INTO 
    user_account(username, password) 
    VALUES($1, crypt($2, gen_salt('bf')))