-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS feedback (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255)   NOT NULL,
    email           VARCHAR(255)   NOT NULL,
    subject         VARCHAR(255)   NOT NULL,
    message         TEXT           NOT NULL DEFAULT '', 
    created_at      TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP
);

CREATE TRIGGER update_feedback_updated_at 
    BEFORE UPDATE ON feedback 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS feedback;
-- +goose StatementEnd