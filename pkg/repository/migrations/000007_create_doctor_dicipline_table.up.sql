BEGIN ;
CREATE TABLE IF NOT EXISTS hsp_doctor_dicipline
(
    id   SERIAL  NOT NULL,
    name VARCHAR NOT NULL,
    primary key (id)
);
INSERT INTO hsp_doctor_dicipline (id,name) VALUES (1,'mata');
COMMIT ;