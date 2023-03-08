CREATE TABLE IF NOT EXISTS users
(
    id           VARCHAR(37) PRIMARY KEY NOT NULL,
    username     VARCHAR(30) UNIQUE      NOT NULL,
    password     VARCHAR(64)             NOT NULL,
    email        varchar(45) UNIQUE      NOT NULL,
    phone_number varchar(15),
    role         int
);

