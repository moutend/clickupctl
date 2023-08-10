CREATE TABLE responses (
  path TEXT NOT NULL,
  body TEXT NOT NULL,
  cached_at INTEGER NOT NULL,
  expired_at INTEGER NOT NULL,

  PRIMARY KEY (path)
);
