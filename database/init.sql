-- init.sql

CREATE TYPE role AS ENUM ('coach', 'normal_user' );

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    role role,
    password_hash VARCHAR(255),
    salt VARCHAR(255),
    creation_time timestamp
);

CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    coach_id INT REFERENCES users(id),
    appointmet_date timestamp,
    appointment_user_id INT REFERENCES users(id),
    creation_time timestamp
);
-- Aggiungi altre query per creare ulteriori tabelle se necessario
