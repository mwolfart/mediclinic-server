-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL NOT NULL,
    description varchar(50) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS rolesOnUser (
    userId integer NOT NULL,
    roleId integer NOT NULL,
    PRIMARY KEY(userId, roleId),
    FOREIGN KEY(userId) REFERENCES genericUser(id),
    FOREIGN KEY(roleId) REFERENCES roles(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rolesOnUser;
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
