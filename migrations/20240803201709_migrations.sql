-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM ('client', 'moderator');

CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL PRIMARY KEY,
    email         VARCHAR(20) UNIQUE NOT NULL,
    password_hash TEXT               NOT NULL,
    role          user_role          NOT NULL
);

CREATE TABLE IF NOT EXISTS houses
(
    id         BIGSERIAL PRIMARY KEY,
    address    VARCHAR(30) UNIQUE NOT NULL,
    year       INT                NOT NULL,
    developer  VARCHAR(20)        NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS flats
(
    id       BIGSERIAL PRIMARY KEY,
    number   INTEGER NOT NULL,
    price    INTEGER NOT NULL,
    house_id INT     NOT NULL REFERENCES houses (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS user_role;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS houses;
DROP TABLE IF EXISTS flats;
-- +goose StatementEnd
