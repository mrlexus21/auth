-- +goose Up
-- +goose StatementBegin
create table user_ent (
    id serial primary key,
    name text not null,
    email text not null,
    password text not null,
    password_confirm text not null,
    role int not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user_ent;
-- +goose StatementEnd
