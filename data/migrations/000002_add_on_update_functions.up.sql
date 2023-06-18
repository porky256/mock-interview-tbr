
CREATE OR REPLACE FUNCTION update_row_modified_function_()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END;
$$
language PLPGSQL;

CREATE TRIGGER row_mod_on_users_trigger_ BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE PROCEDURE update_row_modified_function_();

CREATE TRIGGER row_mod_on_skill_trigger_ BEFORE UPDATE ON skills
    FOR EACH ROW EXECUTE PROCEDURE update_row_modified_function_();

CREATE TRIGGER row_mod_on_users_skills_trigger_ BEFORE UPDATE ON users_skills
    FOR EACH ROW EXECUTE PROCEDURE update_row_modified_function_();

CREATE TRIGGER row_mod_on_users_matches_trigger_ BEFORE UPDATE ON users_matches
    FOR EACH ROW EXECUTE PROCEDURE update_row_modified_function_();

CREATE TRIGGER row_mod_on_match_requests_trigger_ BEFORE UPDATE ON  match_requests
    FOR EACH ROW EXECUTE PROCEDURE update_row_modified_function_();

CREATE TRIGGER row_mod_on_skills_in_requests_trigger_ BEFORE UPDATE ON skills_in_requests
    FOR EACH ROW EXECUTE PROCEDURE update_row_modified_function_();

CREATE TRIGGER row_mod_on_interviews_trigger_ BEFORE UPDATE ON interviews
    FOR EACH ROW EXECUTE PROCEDURE update_row_modified_function_();
