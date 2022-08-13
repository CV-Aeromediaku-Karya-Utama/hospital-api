create table if not exists inv_pre_order_detail
(
    id      serial not null,
    po_item text    null,
    primary key (id)
);