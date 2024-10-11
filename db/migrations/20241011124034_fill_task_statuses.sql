-- +goose Up
-- +goose StatementBegin
insert into task_status(task_status) values
('Backlog'),
('In work'),
('Done');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from task_status
where task_status in ('Backlog', 'In work', 'Done');
-- +goose StatementEnd
