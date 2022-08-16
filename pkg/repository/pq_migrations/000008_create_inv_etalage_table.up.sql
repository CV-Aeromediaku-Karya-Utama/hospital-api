create table if not exists inv_etalage
(
    id               serial not null,
    product_id     int not null,
    sell_per_item_id int not null,
    current_stock    int not null default 0,
    min_stock        int not null default 0,
    status           int          default 0,
    created_at       timestamp             default now(),
    updated_at       timestamp    null,
    primary key (id),
    foreign key (product_id) references inv_product (id),
    foreign key (sell_per_item_id) references inv_sell_per_item (id)
);