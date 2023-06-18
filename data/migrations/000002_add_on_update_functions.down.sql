DROP TRIGGER IF EXISTS row_mod_on_users_trigger_ ON users;
DROP TRIGGER IF EXISTS row_mod_on_users_skills_trigger_ ON users_skills;
DROP TRIGGER IF EXISTS row_mod_on_users_matches_trigger_ ON users_matches;
DROP TRIGGER IF EXISTS row_mod_on_match_requests_trigger_ ON match_requests;
DROP TRIGGER IF EXISTS row_mod_on_skills_in_requests_trigger_ ON skills_in_requests;
DROP TRIGGER IF EXISTS row_mod_on_interviews_trigger_ ON interviews;

DROP FUNCTION IF EXISTS update_row_modified_function_();