CREATE TABLE notes (
                       id SERIAL PRIMARY KEY,
                       user_id INT REFERENCES users(id),
                       title VARCHAR(200),
                       content TEXT,
                       category VARCHAR(50),
                       created_at TIMESTAMPTZ DEFAULT NOW(),
                       updated_at TIMESTAMPTZ,
                       deleted_at TIMESTAMPTZ
);