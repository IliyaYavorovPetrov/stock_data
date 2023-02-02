CREATE TABLE stocks.record_stocks_by_uuid
(
    record_stock_uuid timeuuid,
    stock_name      text,
    price           float,
    PRIMARY KEY (record_stock_uuid)
);
