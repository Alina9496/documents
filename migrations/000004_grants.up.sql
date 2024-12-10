CREATE TABLE IF NOT EXISTS grants(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id uuid not null,
    document_id uuid not null,
    grant_user_login text not null,
    created_at timestamp not null DEFAULT CURRENT_TIMESTAMP
);