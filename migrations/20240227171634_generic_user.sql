-- +goose Up
CREATE TABLE IF NOT EXISTS genericUser (
    id SERIAL NOT NULL,
    name varchar(50) NOT NULL,
    age integer,
    gender varchar(20),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS genericUser;
