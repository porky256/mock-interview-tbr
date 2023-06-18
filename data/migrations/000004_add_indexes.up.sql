CREATE UNIQUE INDEX users_email_idx ON users (email);

CREATE UNIQUE INDEX users_skills_user_idx ON users_skills (user_id);

CREATE UNIQUE INDEX users_matches_user_asker_idx ON users_matches (user_asker);

CREATE UNIQUE INDEX match_requests_user_idx ON match_requests (user_id);

CREATE UNIQUE INDEX skills_in_requests_request_idx ON skills_in_requests (request_id);