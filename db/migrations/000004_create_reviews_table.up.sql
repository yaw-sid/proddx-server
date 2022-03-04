CREATE TABLE IF NOT EXISTS reviews(
    id UUID PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id),
    product_id UUID NOT NULL REFERENCES products(id),
    comment TEXT,
    rating INT,
    created_at TIMESTAMP
);