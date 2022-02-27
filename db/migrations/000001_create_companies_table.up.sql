CREATE TABLE IF NOT EXISTS companies(
    id UUID PRIMARY KEY,
    company_user_id VARCHAR(60) UNIQUE NOT NULL,
    company_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    logo TEXT,
    created_at TIMESTAMP
);