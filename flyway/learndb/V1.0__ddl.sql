CREATE TABLE company (
    id SERIAL PRIMARY KEY
    ,company_name TEXT
    ,representative TEXT
    ,phone_number TEXT
);

CREATE TABLE department (
    id SERIAL PRIMARY KEY
    ,department_name TEXT
    ,email TEXT
    ,company_id BIGINT NOT NULL
);

CREATE TABLE employee (
    id SERIAL PRIMARY KEY
    ,name TEXT
    ,gender TEXT
    ,email TEXT
    ,latestLogin_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    ,dependents_num INT NOT NULL DEFAULT 0
    ,isManager BOOLEAN NOT NULL DEFAULT false
    ,department_id BIGINT NOT NULL
    ,company_id BIGINT NOT NULL
);


