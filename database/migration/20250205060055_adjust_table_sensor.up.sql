ALTER TABLE sensors MODIFY COLUMN id VARCHAR(36) NOT NULL;
ALTER TABLE sensors MODIFY COLUMN image LONGBLOB;
ALTER TABLE sensors DROP latitude;
ALTER TABLE sensors DROP langitude;
ALTER TABLE sensors ADD gateway_id VARCHAR(36) NOT NULL;
ALTER TABLE sensors ADD node_id VARCHAR(36) NOT NULL;
ALTER TABLE sensors ADD coordinate VARCHAR(25) NOT NULL;
