CREATE TABLE companies (
    id SERIAL PRIMARY KEY
    ,company_name TEXT
    ,representative TEXT
    ,phone_number TEXT
);

CREATE TABLE departments (
    id SERIAL PRIMARY KEY
    ,department_name TEXT
    ,email TEXT
    ,company_id BIGINT NOT NULL
);

CREATE TABLE employees (
    id SERIAL PRIMARY KEY
    ,name TEXT
    ,gender TEXT
    ,email TEXT
    ,latest_login_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    ,dependents_num INT NOT NULL DEFAULT 0
    ,is_manager BOOLEAN NOT NULL DEFAULT false
    ,department_id BIGINT
    ,company_id BIGINT NOT NULL
);


