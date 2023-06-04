create table user_datas(
    id uuid not null DEFAULT gen_random_uuid(),
    gender varchar(255),
    name varchar(255),
    address text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,

    primary key(id),
    constraint fk_gender foreign key(gender) references genders(code) on update cascade
);