-- +goose Up
-- +goose StatementBegin
INSERT INTO users (user_id, name, batch, email, password, created_at, created_by, updated_at, updated_by)
VALUES ('ffb541de-a833-4d8d-99c9-4ec229b55716', 'dip', 15, 'khdip.ku@gmail.com', 'Secure@123','2025-08-24 06:30:19.7526','ffb541de-a833-4d8d-99c9-4ec229b55716','2021-08-24 06:30:19.7526','ffb541de-a833-4d8d-99c9-4ec229b55716');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE user_id = 'ffb541de-a833-4d8d-99c9-4ec229b55716';
-- +goose StatementEnd
