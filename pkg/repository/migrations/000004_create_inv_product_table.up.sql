create table if not exists inv_product
(
    id                  serial       not null,
    name                varchar(100) not null,
    product_desc        text         null,
    product_category_id jsonb         not null,
    primary key (id)
)