CREATE TABLE IF NOT EXISTS segments (
    id               BIGINT PRIMARY KEY,
    start_sec        INT,
    duration_sec     INT,
    summary          TEXT,

    -- embedding vector for summary; dimension must match Gemini's default
    embedding        VECTOR(768)
);