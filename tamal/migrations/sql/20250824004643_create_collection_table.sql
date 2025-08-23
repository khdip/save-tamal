-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS collection
(
    collection_id               SERIAL PRIMARY KEY,
    account_type                VARCHAR(100) NOT NULL DEFAULT '',
    account_number              VARCHAR(100) NOT NULL DEFAULT '',
    sender                      VARCHAR(255) NOT NULL DEFAULT 'Anonymous',
    date                        TIMESTAMP    DEFAULT current_timestamp,
    created_at                  TIMESTAMP    DEFAULT current_timestamp,
    created_by                  VARCHAR(100) NOT NULL DEFAULT '',
    updated_at                  TIMESTAMP    DEFAULT current_timestamp,
    updated_by                  VARCHAR(100) NOT NULL DEFAULT '',
    deleted_at                  TIMESTAMP    DEFAULT NULL,
    deleted_by                  VARCHAR(100) NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS collection;
-- +goose StatementEnd
