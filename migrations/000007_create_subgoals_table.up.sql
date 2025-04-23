CREATE TABLE subgoals (
                          id SERIAL PRIMARY KEY,
                          goal_id INT REFERENCES goals(id),
                          title VARCHAR(200),
                          is_done BOOLEAN DEFAULT FALSE,
                          created_at TIMESTAMPTZ DEFAULT NOW(),
                          updated_at TIMESTAMPTZ,
                          deleted_at TIMESTAMPTZ
);