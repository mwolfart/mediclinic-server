-- +goose Up
CREATE TABLE IF NOT EXISTS specialties (
    id SERIAL NOT NULL,
    name varchar(30) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS specialtiesOnDoctor (
    userId integer NOT NULL,
    specialtyId integer NOT NULL,
    PRIMARY KEY(userId, specialtyId),
    FOREIGN KEY(userId) REFERENCES doctor(userId),
    FOREIGN KEY(specialtyId) REFERENCES specialties(id)
);

-- +goose Down
DROP TABLE IF EXISTS specialtiesOnDoctor;
DROP TABLE IF EXISTS specialties;
