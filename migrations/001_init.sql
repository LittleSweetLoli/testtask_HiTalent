-- +goose Up
CREATE TABLE departments (
    id INTEGER PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    parent_id INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW() 

    CONSTRAINT fk_department_parent
        FOREIGN KEY (parent_id)
        REFERENCES department(id)
        ON DELETE SET NULL 
);

CREATE TABLE employees (
    id INTEGER PRIMARY KEY,
    department_id INTEGER NOT NULL,
    full_name VARCHAR(200) NOT NULL,
    position VARCHAR(200) NOT NULL, 
    hired_at DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()

    CONSTRAINT fk_employee_department
        FOREIGN KEY (department_id)
        REFERENCES department(id) 
        ON DELETE CASCADE, 
);

-- +goose Down
DROP TABLE IF EXISTS employees;
DROP TABLE IF EXISTS departments;

