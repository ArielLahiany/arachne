CREATE TABLE IF NOT EXISTS attack (
    id          SERIAL          PRIMARY KEY,
    created_at  TIMESTAMP       NOT NULL        DEFAULT NOW(),
    updated_at  TIMESTAMP       NOT NULL        DEFAULT NOW(),
    type        VARCHAR(20)     NOT NULL,
    target      VARCHAR(2048)   NOT NULL
);
