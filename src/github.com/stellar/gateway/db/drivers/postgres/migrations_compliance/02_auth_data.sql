-- +migrate Up
CREATE TABLE AuthData (
  id varchar(255) NOT NULL,
  domain varchar(255) NOT NULL,
  auth_data text NOT NULL,
  
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE AuthData;
