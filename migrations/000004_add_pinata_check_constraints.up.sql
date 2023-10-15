ALTER TABLE Pinatas ADD CONSTRAINT pinata_weight_check CHECK (weight >= 0);
ALTER TABLE Pinatas ADD CONSTRAINT pinata_height_check CHECK (weight >= 0);
ALTER TABLE Pinatas ADD CONSTRAINT pinata_width_check CHECK (weight >= 0);
ALTER TABLE Pinatas ADD CONSTRAINT pinata_depth_check CHECK (weight >= 0);