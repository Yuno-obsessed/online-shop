CREATE TABLE IF NOT EXISTS users
(
    user_uuid VARCHAR(36),
    first_name VARCHAR(40) NOT NULL,
    last_name VARCHAR(40) NOT NULL,
    nickname VARCHAR(90) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    password VARCHAR(64) NOT NULL,
--     salt VARCHAR(64) NOT NULL,
    image VARCHAR(120) NOT NULL,
    purchases INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (user_uuid)
    );

CREATE TABLE IF NOT EXISTS categories
(
    category_id int,
    category_name VARCHAR(100),
    PRIMARY KEY (category_id)
    );

CREATE TABLE IF NOT EXISTS products
(
    product_uuid VARCHAR(36),
    product_name VARCHAR(120) NOT NULL,
    description VARCHAR(200) NOT NULL,
    category VARCHAR(36) NOT NULL,
    image VARCHAR(120) NOT NULL,
    seller VARCHAR(36) NOT NULL,
    price INT NOT NULL,
    quantity INT,
    likes INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (product_uuid)
    -- FOREIGN KEY (seller)
      -- REFERENCES users (user_uuid)
        -- ON DELETE CASCADE,
    -- FOREIGN KEY (category)
    --    REFERENCES categories (category_id)
    );

CREATE TABLE IF NOT EXISTS carts
(
    cart_uuid VARCHAR(36),
    user_uuid VARCHAR(36) NOT NULL,
    product_uuid VARCHAR(36) NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (cart_uuid),
    FOREIGN KEY (user_uuid)
        REFERENCES users (user_uuid)
        ON DELETE CASCADE,
    FOREIGN KEY (product_uuid)
        REFERENCES  products (product_uuid)
        ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS likes
(
    liked_uuid VARCHAR(36),
    user_uuid VARCHAR(36) NOT NULL,
    product_uuid VARCHAR(36) NOT NULL,
    PRIMARY KEY (liked_uuid),
    FOREIGN KEY (user_uuid)
        REFERENCES users (user_uuid)
        ON DELETE CASCADE,
    FOREIGN KEY (product_uuid)
        REFERENCES products (product_uuid)
        ON DELETE CASCADE
    );

INSERT INTO categories
    (category_id, category_name)
    VALUES
        (1, "women's fashion"),
        (2, "accessories"),
        (3, "men's fashion"),
        (4, "phones & telecomunications"),
        (5, "laptops"),
        (6, "storage devices"),
        (7, "security & protection"),
        (8, "camera & photo"),
        (9, "videogames"),
        (10, "jewelry"),
        (11, "arts & crafts & sewing"),
        (12, "home decor"),
        (13, "furniture"),
        (14, "household items"),
        (15, "outdoor fun & sports"),
        (16, "beauty & hair & health"),
        (17, "automobiles & motorcycles"),
        (18, "tools & home improvement");