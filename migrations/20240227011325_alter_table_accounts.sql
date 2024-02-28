-- migrate:up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS accounts
(
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    email         VARCHAR(255) UNIQUE                         NOT NULL,
    username      VARCHAR(255) UNIQUE                         NOT NULL,
    gender        VARCHAR(6) CHECK (LOWER(gender) IN ('male', 'female') OR UPPER(gender) IN ('MALE', 'FEMALE')),
    password_hash VARCHAR(255)                                NOT NULL,
    date_of_birth DATE                                        NOT NULL,
    is_verified   BOOLEAN          DEFAULT FALSE              NOT NULL,
    bio           TEXT,
    location      VARCHAR(255) NOT NULL,
    avatar        VARCHAR(255),
    created_at    TIMESTAMPTZ                                 NOT NULL DEFAULT (now()),
    updated_at    TIMESTAMPTZ
);

-- migrate:down
DROP TABLE accounts;