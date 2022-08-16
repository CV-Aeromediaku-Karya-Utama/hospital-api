create table if not exists inv_additional
(
    id   serial      not null,
    name varchar(100) not null,
    type int      not null,
    primary key (id)
);