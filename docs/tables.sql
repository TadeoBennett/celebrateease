
-- this must be done as a postgresql user 

DROP DATABASE IF EXISTS  celebrateease; 
CREATE DATABASE  celebrateease;

REVOKE ALL PRIVILEGES ON DATABASE  celebrateease FROM celebrateease_user;
DROP ROLE IF EXISTS celebrateease_user;
CREATE ROLE celebrateease_user LOGIN PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE  celebrateease TO celebrateease_user;
ALTER ROLE celebrateease_user SUPERUSER;


