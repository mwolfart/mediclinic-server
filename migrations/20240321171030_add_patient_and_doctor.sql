-- +goose Up
CREATE TABLE IF NOT EXISTS patient (
    userId integer NOT NULL,
    PRIMARY KEY(userId),
    FOREIGN KEY(userId) REFERENCES genericUser(id)    
);

CREATE TABLE IF NOT EXISTS doctor (
    userId integer NOT NULL,
    PRIMARY KEY(userId),
    FOREIGN KEY(userId) REFERENCES genericUser(id)    
);

-- +goose Down
DROP TABLE IF EXISTS patient;
DROP TABLE IF EXISTS doctor;
