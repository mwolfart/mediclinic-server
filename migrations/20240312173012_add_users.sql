-- +goose Up
-- +goose StatementBegin
INSERT INTO genericUser (name, email, age, gender, passwordHash) VALUES
('John Doe', 'john.doe@mailinator.com', 30, 'Male', '$2a$10$tTUZo5C31y5/mnvrexv4jubIGZCy8f4Xe8Hz2v3s1IRZwGJxB1DSK'),
('Joanna Kim', 'joanna.kim@mailinator.com', 35, 'Female', '$2a$10$tTUZo5C31y5/mnvrexv4jubIGZCy8f4Xe8Hz2v3s1IRZwGJxB1DSK'),
('Keith Williams', 'keith.williams@mailinator.com', 20, 'Nonbinary', 'abc123'),
('Mark Adams', 'mark.adams@mailinator.com', 15, 'Male', 'abc123'),
('Olivia Evans', 'olivia.evans@mailinator.com', 70, 'Female', 'abc123'),
('Liz Lizzle', 'liz.lizzle@mailinator.com', 65, 'Female', 'abc123'),
('Paul Evans', 'paul.evans@mailinator.com', 30, 'Female', 'abc123'),
('Pauline Harris', 'pauline.harris@mailinator.com', 29, 'Female', 'abc123');

INSERT INTO rolesOnUser VALUES
(1, 1),
(2, 1),
(3, 3),
(4, 3),
(5, 3),
(6, 2),
(7, 2),
(8, 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
