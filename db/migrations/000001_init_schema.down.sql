-- These two tables are the most directly related to a user's profile
-- I assume that other tables that have some relationship to users (joinRequests, member,
-- participant) belong to the groups microservice
DROP TABLE IF EXISTS "User";
DROP TABLE IF EXISTS "UserInterests";
-- I had to add this table to this microservice because UserInterests references it
-- Sadly, this means it is duplicated and is present both in users_ms and groups_ms
DROP TABLE IF EXISTS "Category"
