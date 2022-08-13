create table if not exists inv_warehouse
(
    id                  serial      not null,
    commodity_id int not null ,
    buy_per_item_id int not null ,
    current_stock       int      not null default 0,
    min_stock           int      not null default 0,
    status              int               default 0,
    created_at          timestamp                  default now(),
    updated_at          timestamp         null,
    primary key (id),
    foreign key (commodity_id) references inv_commodity(id),
    foreign key (buy_per_item_id) references inv_buy_per_item(id)
);