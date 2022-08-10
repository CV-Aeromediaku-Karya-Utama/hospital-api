create table if not exists inv_pre_order
(
    id      integer      not null auto_increment,
    po_name varchar(100) not null,
    status  integer      not null default 0,
    po_desc text         null,
    primary key (id)
);

