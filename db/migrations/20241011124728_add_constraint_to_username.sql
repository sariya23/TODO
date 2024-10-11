-- +goose Up
-- +goose StatementBegin
alter table users add constraint unique_username unique (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP CONSTRAINT unique_username;
-- +goose StatementEnd
