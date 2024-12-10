CREATE TABLE IF NOT EXISTS document(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name text not null,
    file text not null,
    mime text not null,
    is_public boolean not null,
    user_id uuid  not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);
