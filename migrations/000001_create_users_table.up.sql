
CREATE TABLE IF NOT EXISTS "bazooka_user" (
    id SERIAL PRIMARY KEY,

    uid BIGINT NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    username character varying(255),
    lang character varying(32),

    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,

    UNIQUE(uid)
);
