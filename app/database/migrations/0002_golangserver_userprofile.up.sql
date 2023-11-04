CREATE TABLE IF NOT EXISTS golangserver.UserProfile (
  userID varchar(15) NOT NULL PRIMARY KEY,
  uName varchar(50),
  uEmail varchar(80),
  uAddress json
);