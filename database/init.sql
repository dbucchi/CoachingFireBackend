-- init.sql

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
        CREATE TYPE role AS ENUM ('coach', 'normal_user');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    role role,
    password_hash VARCHAR(255),
    salt VARCHAR(255),
    creation_time timestamp
);

CREATE TABLE IF NOT EXISTS games (
    id SERIAL PRIMARY KEY,
    name    VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    coach_id INT REFERENCES users(id),
    appointmet_date timestamp,
    appointment_user_id INT REFERENCES users(id),
    game_id INT REFERENCES games(id),
    creation_time timestamp
);
-- Aggiungi altre query per creare ulteriori tabelle se necessario
