-- +goose Up
-- +goose StatementBegin
CREATE TABLE checklist_items (
    id SERIAL PRIMARY KEY,
    item_name VARCHAR(255) NOT NULL,
    checklist_id INT NOT NULL,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE checklist_items;
-- +goose StatementEnd
