CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE meets (
    meet_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    creator_id UUID NOT NULL REFERENCES users(id),
    title VARCHAR(100) NOT NULL,
    description TEXT,
    location VARCHAR(100),
    meet_date TIMESTAMP,
    theme VARCHAR(25),
    meet_banner BYTEA,
    meet_thumbnail BYTEA,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);