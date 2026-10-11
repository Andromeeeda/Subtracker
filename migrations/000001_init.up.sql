CREATE SCHEMA subtracker;

CREATE TABLE subtracker.users (
    user_id SERIAL PRIMARY KEY,
    username TEXT,
    hashPassword TEXT 
);

CREATE TABLE subtracker.subscriptions (
    sub_id  SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES subtracker.users(user_id) ON DELETE CASCADE,
    title TEXT NOT NULL, 
    cost NUMERIC(10,2) NOT NULL, 
    date SMALLINT NOT NULL CHECK (date BETWEEN 1 AND 31)
);