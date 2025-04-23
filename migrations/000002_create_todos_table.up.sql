CREATE TABLE todos (
                       id SERIAL PRIMARY KEY,
                       user_id INT REFERENCES users(id),
                       title VARCHAR(255),
                       description TEXT,
                       is_done BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMPTZ DEFAULT NOW(),
                       updated_at TIMESTAMPTZ,
                       deleted_at TIMESTAMPTZ
);