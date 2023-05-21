-- +goose Up
-- +goose StatementBegin
INSERT INTO users (username, email) VALUES
('Bob', 'bob@gmail.com'),
('Jon', 'jon@gmail.com'),
('Jane', 'jane@gmail.com'),
('John', 'john@gmail.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
