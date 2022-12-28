CREATE TABLE IF NOT EXISTS products
(
    product_uuid UUID DEFAULT uuid_generate_v4() UNIQUE,
    product_name VARCHAR(120) NOT NULL,
    description VARCHAR(200) NOT NULL,
    image VARCHAR(120) NOT NULL,
    seller VARCHAR(200) NOT NULL,
    price DECIMAL(8,2) NOT NULL,
    quantity INT,
    likes INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
);
CREATE TABLE IF NOT EXISTS users
(
    user_uuid UUID DEFAULT uuid_generate_v4() UNIQUE,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    nickname VARCHAR(60),
    age INT NOT NULL,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    password VARCHAR(40) NOT NULL,
    salt VARCHAR(64) NOT NULL,
    purchases INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
