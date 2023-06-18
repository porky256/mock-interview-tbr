CREATE TABLE IF NOT EXISTS users (
    id          SERIAL NOT NULL PRIMARY KEY,
    username    VARCHAR(256) NOT NULL DEFAULT '',
    first_name  VARCHAR(256) NOT NULL DEFAULT '',
    last_name   VARCHAR(256) NOT NULL DEFAULT '',
    email       VARCHAR(256) NOT NULL,
    password    VARCHAR(60) NOT NULL,
    phone       VARCHAR(256) NOT NULL,
    user_status INTEGER NOT NULL DEFAULT 1,
    description TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS skills (
    id          SERIAL NOT NULL PRIMARY KEY,
    name        VARCHAR(256) NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users_skills (
    id         SERIAL NOT NULL PRIMARY KEY,
    skill_id   INTEGER,
    user_id    INTEGER,
    score      INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users_matches (
    id          SERIAL NOT NULL PRIMARY KEY,
    user_asker  INTEGER,
    user_match  INTEGER,
    match_score INTEGER,
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS match_requests (
    id         SERIAL NOT NULL PRIMARY KEY,
    user_id    INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS skills_in_requests (
    id         SERIAL NOT NULL PRIMARY KEY,
    request_id INTEGER,
    skill_id   INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS interviews (
    id             SERIAL NOT NULL PRIMARY KEY,
    match_id       INTEGER,
    status         INTEGER,
    interview_date DATE NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT now(),
    updated_at     TIMESTAMP NOT NULL DEFAULT now()
);
