CREATE TABLE matches (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    project_id BIGINT NOT NULL,
    vendor_id BIGINT NOT NULL,
    score FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY (project_id, vendor_id),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (vendor_id) REFERENCES vendors(id)
);