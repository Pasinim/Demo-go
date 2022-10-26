CREATE TABLE IF NOT EXISTS articolo (
    id serial PRIMARY KEY NOT NULL,
    nome varchar(20),
    sku varchar(20),
    collection_id serial  REFERENCES collezione(id)
);