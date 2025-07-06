CREATE TABLE product {
    id BIGSERIAL PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL
    description text,
    price numeric NOT NULL,
    stock integer NOT NULL,
    category_id interger NOT NULL,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES product_category(id) ON DELETE CASCADE
}