create table if not exists inv_commodity
(
    id                    serial      not null,
    name                  varchar(100) not null,
    commodity_desc        text         null,
    commodity_category_id int      not null,
    primary key (id),
    foreign key (commodity_category_id) references inv_commodity_category (id)
)