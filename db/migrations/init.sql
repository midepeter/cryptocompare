-- +migrate Up
create table pricedb
(
    req_id     int auto_increment primary key,
    par_title  varchar(255)  not null,
    par_height int default 0 not null,
)