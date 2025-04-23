CREATE TABLE goals (
                       id SERIAL PRIMARY KEY,
                       user_id INT REFERENCES users(id),
                       title VARCHAR(200),
                       description TEXT,
                       deadline DATE,
                       is_completed BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMPTZ DEFAULT NOW(),
                       updated_at TIMESTAMPTZ,
                       deleted_at TIMESTAMPTZ
);