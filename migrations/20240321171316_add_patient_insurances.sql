-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS insurancesOnPatient (
    userId integer NOT NULL,
    insuranceId integer NOT NULL,
    PRIMARY KEY(userId, insuranceId),
    FOREIGN KEY(userId) REFERENCES patient(userId),
    FOREIGN KEY(insuranceId) REFERENCES insurances(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS insurancesOnPatient
-- +goose StatementEnd
