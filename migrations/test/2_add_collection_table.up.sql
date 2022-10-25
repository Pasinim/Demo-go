CREATE TABLE IF NOT EXISTS demo.collezione (
    id serial4 PRIMARY KEY NOT NULL,
    nome varchar(20),
    FOREIGN KEY (collezione_id)
    REFERENCES demo.articolo (collezione_id)
);