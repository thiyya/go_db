CREATE TABLE public.Person
(
    id   serial,
    name text,
    PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE public.Person
    OWNER to postgres;

INSERT INTO Person (name) VALUES ('Aysegul_Postgre');
INSERT INTO Person (name) VALUES ('Zeynep_Postgre');
