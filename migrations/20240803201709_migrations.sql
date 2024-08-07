-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM ('client', 'moderator');
CREATE TYPE flat_status AS ENUM ('created', 'on_moderate', 'approved', 'declined');

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
    address    VARCHAR(100) UNIQUE NOT NULL,
    year       INT                NOT NULL,
    developer  VARCHAR(20)        NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS flats
(
    id       BIGSERIAL PRIMARY KEY,
    number   INTEGER     NOT NULL,
    rooms    INTEGER     NOT NULL,
    price    INTEGER     NOT NULL,
    house_id INT         NOT NULL REFERENCES houses (id) ON DELETE CASCADE,
    status   flat_status NOT NULL DEFAULT 'created',

    UNIQUE (number, house_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS flats;
DROP TABLE IF EXISTS houses;
DROP TYPE IF EXISTS user_role;
DROP TYPE IF EXISTS flat_status;
-- +goose StatementEnd
