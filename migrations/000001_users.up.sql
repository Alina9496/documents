BEGIN;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    login text not null,
    password text not null,
    created_at timestamp not null DEFAULT CURRENT_TIMESTAMP
);
COMMIT;