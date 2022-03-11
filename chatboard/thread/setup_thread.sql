DROP TABLE posts;
DROP TABLE threads;

CREATE TABLE threads (
  id         SERIAL PRIMARY KEY,
  uu_id       VARCHAR(255) NOT NULL UNIQUE,
  topic      TEXT,
  num_replies SERIAL,
  owner VARCHAR(255),
  user_id    SERIAL REFERENCES users(id),
  created_at TIMESTAMP NOT NULL       
);

CREATE TABLE posts (
  id         SERIAL PRIMARY KEY,
  uu_id       VARCHAR(255) NOT NULL UNIQUE,
  body       TEXT,
  contributor VARCHAR(255),
  user_id    SERIAL REFERENCES users(id),
  thread_id  SERIAL REFERENCES threads(id),
  created_at TIMESTAMP NOT NULL  
);
