create table if not exists inv_additional
(
    id   integer      not null auto_increment,
    name varchar(100) not null,
    type integer      not null,
    primary key (id)
);