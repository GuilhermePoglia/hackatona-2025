-- +goose Up
-- +goose StatementBegin
CREATE TABLE feedback (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sender_id UUID NOT NULL,
    receiver_id UUID NOT NULL,
    stars INTEGER NOT NULL CHECK (stars >= 1 AND stars <= 5),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_id) REFERENCES employee(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES employee(id) ON DELETE CASCADE
);

CREATE INDEX idx_feedback_receiver_id ON feedback(receiver_id);
CREATE INDEX idx_feedback_sender_id ON feedback(sender_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS feedback;
-- +goose StatementEnd
