CREATE TABLE product
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50),
    price INT NOT NULL

);