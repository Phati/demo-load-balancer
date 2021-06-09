CREATE TABLE users (
    id  int PRIMARY KEY,
    email TEXT,
    pwd VARCHAR(15)
);

CREATE SEQUENCE users_sequence start 1  increment 2;