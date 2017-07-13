CREATE TABLE articles (
  id serial PRIMARY KEY,

  name text NOT NULL,
  content text NOT NULL,
  publish_date timestamp NOT NULL
)
