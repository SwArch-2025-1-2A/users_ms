-- Used to automatically generate UUIDs for new Categories
DROP EXTENSION IF EXISTS "uuid-ossp";

-- These two tables are the most directly related to a user's profile
-- I assume that other tables that have some relationship to users (joinRequests, member,
-- participant) belong to the groups microservice
DROP TABLE IF EXISTS "User";
DROP TABLE IF EXISTS "UserInterests";
