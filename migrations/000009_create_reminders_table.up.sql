CREATE TABLE reminders (
                           id SERIAL PRIMARY KEY,
                           user_id INT REFERENCES users(id),
                           title VARCHAR(255),
                           message TEXT,
                           remind_at TIMESTAMPTZ,
                           is_sent BOOLEAN DEFAULT FALSE,
                           created_at TIMESTAMPTZ DEFAULT NOW(),
                           updated_at TIMESTAMPTZ,
                           deleted_at TIMESTAMPTZ
);