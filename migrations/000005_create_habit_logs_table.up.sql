CREATE TABLE habit_logs (
                            id SERIAL PRIMARY KEY,
                            habit_id INT REFERENCES habits(id),
                            log_date DATE,
                            created_at TIMESTAMPTZ DEFAULT NOW(),
                            updated_at TIMESTAMPTZ,
                            deleted_at TIMESTAMPTZ
);