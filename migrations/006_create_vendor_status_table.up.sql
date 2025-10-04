CREATE TABLE vendor_statues (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    vendor_id BIGINT NOT NULL,
    last_response_at DATETIME,
    sla_expired BOOLEAN DEFAULT FALSE,
    check_run_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_vendor FOREIGN KEY (vendor_id) REFERENCES vendors(id)
);