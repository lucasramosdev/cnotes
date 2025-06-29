CREATE TABLE users (
    id serial PRIMARY KEY,
    email varchar(50) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);
