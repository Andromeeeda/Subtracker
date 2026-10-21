ALTER TABLE subtracker.users ALTER COLUMN username SET NOT NULL;
ALTER TABLE subtracker.users ADD CONSTRAINT users_username_unique UNIQUE(username);
