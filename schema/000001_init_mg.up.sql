CREATE TABLE users
(
    id            serial       primary key     not null unique,
    username      varchar(255)                 not null unique,
    email         varchar(255)                 not null unique,
    password_hash varchar(255)                 not null,
    role          varchar(255)                 not null
);

CREATE TABLE categories
(
    id            serial       primary key not null unique,
    name          varchar(255)             not null unique
);

CREATE TABLE products 
(
    id            serial       primary key not null unique,
    name          varchar(255)             not null unique,
    price         numeric,
    category_id   integer references categories (id)   
);

INSERT INTO users (username, email, password_hash, role)
VALUES 
('Nikittansk', 'niknosokov@mail.ru', '$2y$10$bqfQ9NUGcaQBsCz6G30UIe13s8OsDBw2/.16E96Wmr1PkowC5K5YG', 'admin'),
('Anton', 'test@mail.ru', '$2y$10$bqfQ9NUGcaQBsCz6G30UIe13s8OsDBw2/.16E96Wmr1PkowC5K5YG', 'role');

INSERT INTO categories (name)
VALUES 
('maps'),
('cards'),
('objects');

INSERT INTO products (name, price, category_id)
VALUES 
('Lightning McQueen', 1500, 2),
('Mater', 1000, 2);