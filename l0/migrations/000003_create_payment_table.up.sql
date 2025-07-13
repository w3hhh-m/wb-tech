CREATE TABLE IF NOT EXISTS payment (
    transaction TEXT PRIMARY KEY,
    order_uid TEXT UNIQUE REFERENCES orders(order_uid) ON DELETE CASCADE,
    request_id TEXT,
    currency TEXT NOT NULL,
    provider TEXT NOT NULL,
    amount INTEGER NOT NULL,
    payment_dt BIGINT NOT NULL,
    bank TEXT NOT NULL,
    delivery_cost INTEGER NOT NULL,
    goods_total INTEGER NOT NULL,
    custom_fee INTEGER NOT NULL
);