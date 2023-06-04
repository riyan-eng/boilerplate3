create table users(
    id uuid not null DEFAULT gen_random_uuid(),
    username varchar(255) unique,
    password varchar(255) not null,
    email varchar(255) unique nulls not distinct,
    phone varchar(255) unique nulls not distinct,
    role varchar(255),
    user_data uuid,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,

    primary key(id),
    constraint fk_user_data foreign key(user_data) references user_datas(id) on update cascade,
    constraint fk_role foreign key(role) references roles(code) on update cascade
);