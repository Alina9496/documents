CREATE TABLE IF NOT EXISTS token(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id text not null,
    token text not null,
    created_at timestamp not null DEFAULT CURRENT_TIMESTAMP
);