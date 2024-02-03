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
    price         numeric      
);