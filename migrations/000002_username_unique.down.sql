ALTER TABLE subtracker.users DROP CONSTRAINT users_username_unique;
ALTER TABLE subtracker.users ALTER COLUMN username DROP NOT NULL;