CREATE TABLE IF NOT EXISTS product(
   product_id serial PRIMARY KEY,
   product_code TEXT UNIQUE NOT NULL,
   product_name VARCHAR (200) NOT NULL,
   stock_quantity INTEGER NOT NULL
);