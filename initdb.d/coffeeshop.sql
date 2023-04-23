CREATE TABLE IF NOT EXISTS public.customers
(
  customer_id serial8 NOT NULL,
  full_name text,
  created_at timestamp not null DEFAULT now(),
  CONSTRAINT customer_id_pk PRIMARY KEY (customer_id)
);

CREATE UNIQUE INDEX customers_id_uindex ON public.customers USING btree (customer_id);

CREATE TABLE IF NOT EXISTS public.orders
(
  order_id serial8 NOT NULL,
  customer_id serial8 NOT NULL,
  status text,
  created_at timestamp not null DEFAULT now(),
  CONSTRAINT order_id_pk PRIMARY KEY (order_id),
  CONSTRAINT customer_id_fk FOREIGN KEY(customer_id) REFERENCES public.customers(customer_id)
);

CREATE UNIQUE INDEX order_id_uindex ON public.orders USING btree (order_id);

CREATE TABLE IF NOT EXISTS public.products
(
  product_id serial8 NOT NULL,
  name text,
  description text,
  price real,
  CONSTRAINT product_id_pk PRIMARY KEY (product_id)
);

CREATE UNIQUE INDEX product_id_uindex ON public.products USING btree (product_id);

CREATE TABLE IF NOT EXISTS public.orders_details
(
  order_id serial8 NOT NULL,
  product_id serial8 NOT NULL,
  status text,
  created_at timestamp not null DEFAULT now(),
  CONSTRAINT orders_details_id_pk PRIMARY KEY (
    order_id,
    product_id
  ),
  CONSTRAINT order_id_fk FOREIGN KEY(order_id) REFERENCES public.orders(order_id),
  CONSTRAINT product_id_fk FOREIGN KEY(product_id) REFERENCES public.products(product_id)
);