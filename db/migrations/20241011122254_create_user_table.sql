-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    user_id bigint generated always as identity primary key,
    username text,
    hash_password text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
