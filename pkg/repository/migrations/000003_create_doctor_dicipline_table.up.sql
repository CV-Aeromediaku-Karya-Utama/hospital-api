BEGIN ;
CREATE TABLE IF NOT EXISTS doctor_dicipline
(
    id   SERIAL  NOT NULL,
    name VARCHAR NOT NULL,
    primary key (id)
);
INSERT INTO doctor_dicipline (id,name) VALUES (1,'mata');
COMMIT ;