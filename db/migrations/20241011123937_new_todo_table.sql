-- +goose Up
-- +goose StatementBegin
create table if not exists todo (
    todo_id bigint generated always as identity primary key,
    task text,
    user_id bigint not null references users (user_id),
    task_status_id smallint not null references task_status (task_status_id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists todo;
-- +goose StatementEnd
