CREATE TABLE IF NOT EXISTS PinatasContents (
                                               pinata_id INT,
                                               content_id INT,
                                               PRIMARY KEY (pinata_id, content_id),
    FOREIGN KEY (pinata_id) REFERENCES Pinatas(id),
    FOREIGN KEY (content_id) REFERENCES Contents(id)
    );