CREATE TABLE IF NOT EXISTS todos (
    id INTEGER PRIMARY KEY,
    text VARCHAR(255) NULL,
    checked INTEGER DEFAULT 0 NOT NULL
);

INSERT INTO todos (id, text, checked) VALUES (1, 'This is an example', 0);