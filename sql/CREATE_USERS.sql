CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    mobile_phone VARCHAR(15),
    password_hash VARCHAR(255) NOT NULL,
    host_status VARCHAR(50) NOT NULL,
    interests TEXT[] DEFAULT '{}',
    states TEXT[] DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);