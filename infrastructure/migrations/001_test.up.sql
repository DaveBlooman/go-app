CREATE SCHEMA IF NOT EXISTS articles;

CREATE TABLE articles.data (
    name text NOT NULL,

    content text NOT NULL,

    publish_date timestamp NOT NULL,
);
