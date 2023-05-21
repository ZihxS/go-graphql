-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS meetups (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) UNIQUE NOT NULL,
    `description` TEXT,
    user_id BIGINT NOT NULL REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS meetups;
-- +goose StatementEnd
