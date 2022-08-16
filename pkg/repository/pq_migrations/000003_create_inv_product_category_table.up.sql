create table if not exists inv_product_category
(
    id   SERIAL      not null,
    name varchar(100) not null,
    primary key (id)
)