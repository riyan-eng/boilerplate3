create table genders(
    id uuid not null DEFAULT gen_random_uuid(),
    code varchar(255) unique,
    name varchar(255),
    primary key(id)
);

insert into genders (code, name) values ('L', 'Laki-laki'), ('P', 'Perempuan');

create table user_datas(
    id uuid not null DEFAULT gen_random_uuid(),
    gender uuid,
    name varchar(255),
    address text,
    primary key(id),
    constraint fk_gender foreign key(gender) references genders(id)
);

create table users(
    id uuid not null DEFAULT gen_random_uuid(),
    username varchar(255) unique,
    password varchar(255) not null,
    email varchar(255) unique nulls not distinct,
    phone varchar(255) unique nulls not distinct,
    user_data uuid,
    primary key(id),
    constraint fk_user_data foreign key(user_data) references user_datas(id)
);