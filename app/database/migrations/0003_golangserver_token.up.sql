CREATE TABLE IF NOT EXISTS golangserver.Token (
  userID varchar(15) NOT NULL PRIMARY KEY,
  token varchar(20),
  tokenTime varchar(250)
);