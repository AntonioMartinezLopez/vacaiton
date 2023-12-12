-- Create userservice database and user
CREATE DATABASE userdata;
CREATE USER userservice WITH PASSWORD 'userservice';
GRANT ALL PRIVILEGES ON DATABASE userdata TO userservice;
ALTER DATABASE userdata OWNER TO userservice;

-- Create tripService database and user
CREATE DATABASE trip;
CREATE USER tripservice WITH PASSWORD 'tripservice';
GRANT ALL PRIVILEGES ON DATABASE trip TO tripservice;
ALTER DATABASE trip OWNER TO tripservice;

-- Create queryService database and user
CREATE DATABASE query;
CREATE USER queryservice WITH PASSWORD 'queryservice';
GRANT ALL PRIVILEGES ON DATABASE query TO queryservice;
ALTER DATABASE query OWNER TO queryservice;

-- -- Create root user with superuser privileges
-- CREATE USER root WITH ENCRYPTED PASSWORD 'rootpassword';
-- ALTER USER root WITH SUPERUSER;