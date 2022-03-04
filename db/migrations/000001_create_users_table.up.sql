CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY,
    email VARCHAR(150),
    user_password VARCHAR(150),
    created_at TIMESTAMP
);