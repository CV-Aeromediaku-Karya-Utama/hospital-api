BEGIN ;
CREATE TABLE IF NOT EXISTS role
(
    id   SERIAL NOT NULL,
    name VARCHAR(100),
    PRIMARY KEY (id)
);
INSERT INTO role (name) VALUES ('ADMIN'),('USER');
COMMIT ;