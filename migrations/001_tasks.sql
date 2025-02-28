-- +goose Up

create table if not exists tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
    );

create index tasks_id on tasks (id);

-- +goose Down
drop table if exists tasks;
drop index tasks_id;