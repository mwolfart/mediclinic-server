-- +goose Up
CREATE TABLE IF NOT EXISTS insurances (
    id SERIAL NOT NULL,
    name varchar(30) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS insurancesOnUser (
    userId integer NOT NULL,
    insuranceId integer NOT NULL,
    PRIMARY KEY(userId, insuranceId),
    FOREIGN KEY(userId) REFERENCES genericUser(id),
    FOREIGN KEY(insuranceId) REFERENCES insurances(id)
);

-- +goose Down
DROP TABLE IF EXISTS insurancesOnUser;
DROP TABLE IF EXISTS insurances;
