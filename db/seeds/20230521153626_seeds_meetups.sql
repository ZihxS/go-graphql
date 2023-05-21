-- +goose Up
-- +goose StatementBegin
INSERT INTO meetups (name, description, user_id) VALUES
('A first meetup for Bob', 'A description first meetup for Bob', 1),
('A first meetup for Jon', 'A first meetup for Jon', 2),
('A second meetup for Jon', 'A second meetup for Jon', 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
