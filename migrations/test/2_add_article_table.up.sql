CREATE TABLE IF NOT EXISTS articolo (
    id serial PRIMARY KEY NOT NULL,
    nome varchar(20),
    sku integer,
    collection_id serial  REFERENCES collezione(id)
);