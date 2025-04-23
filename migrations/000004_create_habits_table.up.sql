CREATE TABLE habits (
                        id SERIAL PRIMARY KEY,
                        user_id INT REFERENCES users(id),
                        name VARCHAR(100),
                        created_at TIMESTAMPTZ DEFAULT NOW(),
                        updated_at TIMESTAMPTZ,
                        deleted_at TIMESTAMPTZ
);