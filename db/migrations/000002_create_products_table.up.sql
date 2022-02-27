CREATE TABLE IF NOT EXISTS products(
    id UUID PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id),
    product_name VARCHAR(100) NOT NULL,
    feedback_url TEXT,
    rating INT,
    created_at TIMESTAMP
);