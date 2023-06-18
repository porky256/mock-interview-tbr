ALTER TABLE users_skills
    ADD CONSTRAINT fk_users_skills_skill_id
        FOREIGN KEY (skill_id)
            REFERENCES skills(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE users_skills
    ADD CONSTRAINT fk_users_skills_user_id
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE users_matches
    ADD CONSTRAINT fk_users_matches_user_asker
        FOREIGN KEY (user_asker)
            REFERENCES users(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE users_matches
    ADD CONSTRAINT fk_users_matches_user_match
        FOREIGN KEY (user_match)
            REFERENCES users(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE match_requests
    ADD CONSTRAINT fk_match_requests_user_id
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE skills_in_requests
    ADD CONSTRAINT fk_skills_in_requests_request_id
        FOREIGN KEY (request_id)
            REFERENCES match_requests(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE skills_in_requests
    ADD CONSTRAINT fk_skills_in_requests_skill_id
        FOREIGN KEY (skill_id)
            REFERENCES skills(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE interviews
    ADD CONSTRAINT fk_interviews_match_id
        FOREIGN KEY (match_id)
            REFERENCES users_matches(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

