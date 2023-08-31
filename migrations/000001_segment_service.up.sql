CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    email VARCHAR (50) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS segments(
    id serial PRIMARY KEY,
    slug VARCHAR (255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users_segments(
    user_id int references users (id) on delete cascade not null,
    segment_id int references segments (id) on delete cascade not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (user_id, segment_id)
);