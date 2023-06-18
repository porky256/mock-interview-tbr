ALTER TABLE users
    DROP CONSTRAINT fk_users_skills_skill_id;

ALTER TABLE users
    DROP CONSTRAINT fk_users_skills_user_id;

ALTER TABLE users_matches
    DROP CONSTRAINT fk_users_matches_user_asker;

ALTER TABLE users_matches
    DROP CONSTRAINT fk_users_matches_user_match;

ALTER TABLE match_requests
    DROP CONSTRAINT fk_match_requests_user_id;

ALTER TABLE skills_in_requests
    DROP CONSTRAINT fk_skills_in_requests_request_id;

ALTER TABLE skills_in_requests
    DROP CONSTRAINT fk_skills_in_requests_skill_id;

ALTER TABLE interviews
    DROP CONSTRAINT fk_interviews_match_id;
