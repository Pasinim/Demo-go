CREATE TABLE IF NOT EXISTS demo.articolo (
    id serial PRIMARY KEY NOT NULL,
    name varchar(20),
    sku varchar(10),
    collezione_id serial
);