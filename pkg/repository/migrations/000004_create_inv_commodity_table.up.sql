create table if not exists inv_commodity
(
    id                    integer      not null auto_increment,
    name                  varchar(100) not null,
    commodity_desc        text         null,
    commodity_category_id integer      not null,
    primary key (id),
    foreign key (commodity_category_id) references inv_commodity_category (id)
)