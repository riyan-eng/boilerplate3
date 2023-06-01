create table tasks(
    id uuid not null DEFAULT gen_random_uuid(),
    name varchar(255) not null,
    detail text,
    primary key(id)
)
