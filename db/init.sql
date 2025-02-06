CREATE DATABASE pingdb;
\c pingdb;
CREATE TABLE IF NOT EXISTS pings (
    ip TEXT PRIMARY KEY,
    time INT,
    last_active TIMESTAMP
);