CREATE TABLE pages (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    banner_media VARCHAR(255),
    content TEXT,
    publication_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    CONSTRAINT fk_pages_created_by FOREIGN KEY (created_by) REFERENCES users(id),
    CONSTRAINT fk_pages_updated_by FOREIGN KEY (updated_by) REFERENCES users(id)
);


-- Index pada kolom slug
CREATE INDEX idx_pages_slug ON pages(slug);

-- Index pada kolom title
CREATE INDEX idx_pages_title ON pages(title);

-- Index pada kolom publication_date
CREATE INDEX idx_pages_publication_date ON pages(publication_date);

