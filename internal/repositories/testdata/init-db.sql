CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    hashed_password TEXT NOT NULL,
    birth_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    extract TEXT NOT NULL,
    content TEXT NOT NULL,
    author_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Insert users
INSERT INTO users (id, first_name, last_name, email, username, hashed_password, birth_date, created_at, updated_at)
VALUES
    ('0853f607-2422-4631-8526-832edaa479c4', 'Alice', 'Smith', 'alice@example.com', 'alice_s', 'hashed_pwd_1', '1990-04-12', '2006-01-02 00:00 UTC', '2006-01-02 00:00 UTC'),
    ('b2ccc80d-606e-422f-a9e1-5fd7371163db', 'Bob', 'Johnson', 'bob@example.com', 'bobby_j', 'hashed_pwd_2', '1988-09-25', '2006-01-02 00:00 UTC', '2006-01-02 00:00 UTC'),
    ('2cdc1c8f-9985-4b6c-b007-038a5bef22b5', 'Charlie', 'Brown', 'charlie@example.com', 'charlie_b', 'hashed_pwd_3', '1995-02-07', '2006-01-02 00:00 UTC', '2006-01-02 00:00 UTC')
ON CONFLICT (id) DO NOTHING;

-- Insert posts for User 1 (2 posts)
INSERT INTO posts (id, title, extract, content, author_id, created_at, updated_at)
VALUES
    ('91c1538a-518c-4b05-9a1e-180c561a70b3', 'My First Post', 'This is my first post extract.', 'This is the full content of my first post.', '0853f607-2422-4631-8526-832edaa479c4', '2006-01-02 00:00 UTC', '2006-01-02 00:00 UTC'),
    ('4c09ea12-30ec-4fea-a667-15be9f13e476', 'Another Day in the Life', 'A short story extract.', 'A longer text describing my second post.', '0853f607-2422-4631-8526-832edaa479c4', '2006-01-02 00:00 UTC', '2006-01-02 00:00 UTC')
ON CONFLICT (id) DO NOTHING;

-- Insert posts for User 2 (1 post)
INSERT INTO posts (id, title, extract, content, author_id, created_at, updated_at)
VALUES
    ('bbf19b79-cf9c-4a07-8e43-299baf69b418', 'Hello World', 'My first blog entry.', 'This is the main content of my post.', 'b2ccc80d-606e-422f-a9e1-5fd7371163db', '2006-01-02 00:00 UTC', '2006-01-02 00:00 UTC')
ON CONFLICT (id) DO NOTHING;
