ALTER TABLE subtracker.users 
ADD CHECK (char_length(username) BETWEEN 3 AND 30);

ALTER TABLE subtracker.users
ADD CHECK (username !~ '^[0-9]+$');

ALTER TABLE subtracker.subscriptions
ADD CHECK(char_length(title) <= 30);