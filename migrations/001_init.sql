-- +goose Up
CREATE TABLE IF NOT EXISTS departments (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    parent_id BIGINT REFERENCES departments(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS uniq_dept_name_parent
    ON departments (name, COALESCE(parent_id, 0));

CREATE TABLE IF NOT EXISTS employees (
    id BIGSERIAL PRIMARY KEY,
    department_id BIGINT NOT NULL REFERENCES departments(id),
    full_name VARCHAR(200) NOT NULL,
    position VARCHAR(200) NOT NULL, 
    hired_at DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS employees;
DROP TABLE IF EXISTS departments;

