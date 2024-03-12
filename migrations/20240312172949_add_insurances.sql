-- +goose Up
-- +goose StatementBegin
INSERT INTO insurances (name) VALUES 
('cigna'),
('healthcare'),
('aetna');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
