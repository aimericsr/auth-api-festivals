-- création trigger pour updated_at
CREATE OR REPLACE FUNCTION function_set_timestamp_update()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- à faire par le superuser
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    username VARCHAR(50) NOT NULL PRIMARY KEY,
    hashed_password VARCHAR(200) NOT NULL,
    full_name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    password_changed_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);