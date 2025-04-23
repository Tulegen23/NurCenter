CREATE TABLE finances (
                          id SERIAL PRIMARY KEY,
                          user_id INT REFERENCES users(id),
                          amount DECIMAL(10,2) NOT NULL,
                          type VARCHAR(10) CHECK (type IN ('income', 'expense')),
                          category VARCHAR(100),
                          description TEXT,
                          date DATE DEFAULT CURRENT_DATE,
                          created_at TIMESTAMPTZ DEFAULT NOW(),
                          updated_at TIMESTAMPTZ,
                          deleted_at TIMESTAMPTZ
);