CREATE TABLE categories (
    id serial PRIMARY KEY,
    description varchar(20) NOT NULL
);

CREATE TABLE themes (
    id serial PRIMARY KEY,
    description varchar(20) NOT NULL
);

CREATE TABLE notes (
    id bigserial PRIMARY KEY,
    category int NOT NULL REFERENCES categories (id),
    theme int NOT NULL REFERENCES themes (id),
    title varchar(30) NOT NULL,
    summary text NOT NULL
);

CREATE TABLE keywords (
    id serial PRIMARY KEY,
    position integer NOT NULL,
    note bigserial NOT NULL REFERENCES notes (id),
    description varchar(50) NOT NULL 
);

CREATE TABLE annotations (
    id serial PRIMARY KEY,
    position integer NOT NULL,
    note bigserial NOT NULL REFERENCES notes (id),
    "value" text NOT NULL
);