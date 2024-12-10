CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE reviews
(
    id SERIAL PRIMARY KEY,
    name_of_reviewer VARCHAR(255),
    text TEXT NOT NULL,
    score INT NOT NULL CHECK (score BETWEEN 1 AND 5),
    category VARCHAR(50) NOT NULL,
    date DATE NOT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    item_id INT REFERENCES items(id) ON DELETE SET NULL
);

CREATE TABLE items
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    description TEXT NOT NULL,
    image_urls TEXT[]
);

CREATE TABLE favorites (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    item_id INT NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    UNIQUE (user_id, item_id)
);

CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    item_id INT REFERENCES items(id),
    size VARCHAR(10),
    UNIQUE(user_id, item_id)
);