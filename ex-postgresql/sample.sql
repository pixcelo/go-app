CREATE TABLE department
(
    department_code character varying(10) NOT NULL,
	department_name character varying(100),
    CONSTRAINT pk_department PRIMARY KEY (department_code)
);