-- +goose Up
-- +goose StatementBegin
INSERT INTO genericUser (name, age, gender) VALUES
('John Doe', 30, 'Male'),
('Joanna Kim', 35, 'Female'),
('Keith Dodds', 20, 'Nonbinary'),
('Mark Adams', 15, 'Male'),
('Olivia Craine', 70, 'Female'),
('Liz Lizzle', 65, 'Female'),
('Taylor Swift', 30, 'Female'),
('Kim Petras', 29, 'Female');

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
