CREATE TABLE media (
    id SERIAL PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL, -- e.g., image or video
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    CONSTRAINT fk_media_created_by FOREIGN KEY (created_by) REFERENCES users(id),
    CONSTRAINT fk_media_updated_by FOREIGN KEY (updated_by) REFERENCES users(id)
);
