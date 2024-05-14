CREATE TABLE users(
    id uuid DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL,
    login VARCHAR(255) NOT NULL UNIQUE,
    birthday_date TIMESTAMP,
    email VARCHAR(255) NOT NULL UNIQUE,
    hash_password VARCHAR(255) NOT NULL,
    status VARCHAR(255),
    img VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)