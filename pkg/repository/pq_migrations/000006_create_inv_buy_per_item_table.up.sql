create table if not exists inv_buy_per_item
(
    id           serial not null,
    commodity_id int not null,
    buy_price    int not null,
    created_at   timestamp default now(),
    updated_at   timestamp    null,
    primary key (id),
    foreign key (commodity_id) references inv_commodity (id)
);