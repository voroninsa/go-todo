-- CREATE TABLE IF NOT EXISTS Users (
--     user_id SERIAL PRIMARY KEY,
--     username VARCHAR(255) NOT NULL,
--     email VARCHAR(255) NOT NULL UNIQUE,
--     password_hash VARCHAR(255) NOT NULL
-- );

-- Создание таблицы Tasks
CREATE TABLE IF NOT EXISTS Tasks (
    task_id SERIAL PRIMARY KEY,
    -- user_id INT NOT NULL,
    description TEXT NOT NULL,
    tags  VARCHAR(255)[] NOT NULL,
    deadline TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed BOOLEAN DEFAULT FALSE,
    -- FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);