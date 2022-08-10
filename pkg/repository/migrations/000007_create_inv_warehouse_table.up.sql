create table if not exists inv_warehouse
(
    id                  integer      not null auto_increment,
    commodity_id integer not null ,
    buy_per_item_id integer not null ,
    current_stock       integer      not null default 0,
    min_stock           integer      not null default 0,
    status              integer               default 0,
    created_at          timestamp                  default now(),
    updated_at          timestamp         null,
    primary key (id),
    foreign key (commodity_id) references inv_commodity(id),
    foreign key (buy_per_item_id) references inv_buy_per_item(id)
);