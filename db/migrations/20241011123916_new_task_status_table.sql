-- +goose Up
-- +goose StatementBegin
create table if not exists task_status (
    task_status_id smallint generated always as identity primary key,
    task_status varchar(20) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists task_status;
-- +goose StatementEnd
