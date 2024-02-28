-- migrate:up
CREATE TABLE memberships
(
    id              SERIAL PRIMARY KEY,
    account_id      UUID        NOT NULL REFERENCES accounts (id),
    membership_type VARCHAR(50) NOT NULL CHECK (LOWER(membership_type) IN ('free', 'premium', 'gold') OR
                                                UPPER(membership_type) IN ('FREE', 'PREMIUM', 'GOLD')),
    start_date      DATE,
    end_date        DATE,
    payment_method  VARCHAR(50),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- migrate:down
DROP TABLE memberships;
