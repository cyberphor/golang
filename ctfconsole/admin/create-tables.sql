CREATE TABLE players (
  id    INT AUTO_INCREMENT NOT NULL,
  user  VARCHAR(64) NOT NULL,
  pass  VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO players (user, pass)

VALUES
  ('cyberphor', MD5('ctfconsole'));