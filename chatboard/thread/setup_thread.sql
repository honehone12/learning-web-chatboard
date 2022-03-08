DROP TABLE posts;
DROP TABLE threads;

CREATE TABLE threads (
  id         SERIAL PRIMARY KEY,
  uu_id       VARCHAR(64) NOT NULL UNIQUE,
  topic      TEXT,
  user_id    INTEGER REFERENCES users(id),
  created_at TIMESTAMP NOT NULL       
);

CREATE TABLE posts (
  id         SERIAL PRIMARY KEY,
  uu_id       VARCHAR(64) NOT NULL UNIQUE,
  body       TEXT,
  user_id    INTEGER REFERENCES users(id),
  thread_id  INTEGER REFERENCES threads(id),
  created_at TIMESTAMP NOT NULL  
);
