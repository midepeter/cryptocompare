-- +migrate Up
CREATE TABLE pricedb (
    req_id     int auto_increment primary key,
    change24hour    varchar(255),
	changepct24hour varchar(255),
	open24hour     varchar(255),
	volume24hour    varchar(255),
	volumde24hourto  varchar(255),
	low24hour    varchar(255),
    high24hour     varchar(255),
	price         varchar(255),
	supply         varchar(255),
	mktcap         varchar(255)
);