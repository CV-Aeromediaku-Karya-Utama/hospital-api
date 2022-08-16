create table if not exists inv_product
(
    id                    serial      not null,
    name                  varchar(100) not null,
    product_desc        text         null,
    product_category_id int      not null,
    primary key (id),
    foreign key (product_category_id) references inv_product_category (id)
)