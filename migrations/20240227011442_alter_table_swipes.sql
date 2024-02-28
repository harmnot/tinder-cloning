-- migrate:up
CREATE TABLE IF NOT EXISTS swipes
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    swiper_id  UUID NOT NULL,
    FOREIGN KEY(swiper_id) REFERENCES accounts(id),
    swiped_id  UUID NOT NULL,
    FOREIGN KEY(swiped_id) REFERENCES accounts(id),
    swipe_type VARCHAR(4) CHECK (LOWER(swipe_type) IN ('like', 'pass') OR UPPER(swipe_type) IN ('LIKE', 'PASS')),
    swipe_date TIMESTAMPTZ NOT NULL DEFAULT (now())
);

-- migrate:down
DROP TABLE swipes;
