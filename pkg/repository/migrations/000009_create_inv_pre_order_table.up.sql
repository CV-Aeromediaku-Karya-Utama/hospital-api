create table if not exists inv_pre_order
(
    id      serial      not null,
    po_name varchar(100) not null,
    status  int      not null default 0,
    po_desc text         null,
    primary key (id)
);

