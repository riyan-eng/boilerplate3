create table genders(
    id uuid not null DEFAULT gen_random_uuid(),
    code varchar(255) unique,
    name varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id)
);

insert into genders (code, name) values ('L', 'Laki-laki'), ('P', 'Perempuan');