-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (description) VALUES 
('manager'),
('doctor'),
('patient');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
