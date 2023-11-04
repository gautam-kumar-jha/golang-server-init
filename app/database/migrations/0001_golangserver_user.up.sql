CREATE TABLE IF NOT EXISTS golangserver.User (
  userID varchar(15) NOT NULL PRIMARY KEY,
  mobileNo varchar(13),  
  pswd varchar(25),
  isActive Char(1)
);