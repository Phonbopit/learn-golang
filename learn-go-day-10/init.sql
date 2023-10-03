CREATE TABLE IF NOT EXISTS posts (
  id INT AUTO_INCREMENT,
  title TEXT NOT NULL,
  body TEXT NOT NULL,
  created_at DATETIME,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
  id INT AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  username varchar(255) NOT NULL,
  active boolean,
  PRIMARY KEY (id)
);