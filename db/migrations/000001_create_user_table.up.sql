CREATE TABLE users (
    id serial PRIMARY KEY, 
    username varchar UNIQUE,
    "name" varchar NOT NULL,
    "password" varchar NOT NULL,
    created_at timestamp,
    updated_at timestamp
); 
