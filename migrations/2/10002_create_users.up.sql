CREATE TABLE users (
    id  int PRIMARY KEY,
    email TEXT,
    pwd VARCHAR(15)
);

CREATE SEQUENCE users_sequence start 2  increment 2;