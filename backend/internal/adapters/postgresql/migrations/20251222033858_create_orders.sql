-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
    );

CREATE TABLE IF NOT EXISTS order_items (
    ID BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    price_in_cents INTEGER NOT NULL CHECK (price_in_cents >= 0),
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id)
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS order_items;

-- +goose StatementEnd
