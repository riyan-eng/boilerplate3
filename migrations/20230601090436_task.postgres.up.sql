create table tasks(
    id uuid not null DEFAULT gen_random_uuid(),
    name varchar(255) not null,
    detail text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id)
)
