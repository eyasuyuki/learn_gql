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
);

CREATE TABLE employee (
    id SERIAL PRIMARY KEY
    ,gender TEXT
    ,email TEXT
    ,latestLogin_at TIMESTAMP
    ,dependents_num INT
    ,isManager BOOLEAN
    ,department_id BIGINT
    ,company_id BIGINT
);


