CREATE TABLE IF NOT EXISTS catalogo (
    id serial PRIMARY KEY NOT NULL,
    name VARCHAR(20),
    articolo_id serial REFERENCES articolo(id)
);