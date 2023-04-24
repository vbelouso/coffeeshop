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

INSERT INTO public.customers(customer_id, full_name, created_at)
VALUES
('1','John Smith', CURRENT_TIMESTAMP),
('2','Jane Doe', CURRENT_TIMESTAMP),
('3','Bob Johnson',CURRENT_TIMESTAMP),
('4','Sara Lee',CURRENT_TIMESTAMP),
('5','Mike Myers',CURRENT_TIMESTAMP),
('6','Lisa Jones',CURRENT_TIMESTAMP),
('7','David Brown',CURRENT_TIMESTAMP),
('8','Amy Adams',CURRENT_TIMESTAMP),
('9','Mark Taylor',CURRENT_TIMESTAMP),
('10','Emily Wilson',CURRENT_TIMESTAMP);


INSERT INTO public.orders(order_id, customer_id, status, created_at)
VALUES
('1','1','New',CURRENT_TIMESTAMP),
('2','2','In Progress',CURRENT_TIMESTAMP),
('3','3','Out for Delivery',CURRENT_TIMESTAMP),
('4','4','Ready',CURRENT_TIMESTAMP),
('5','5','Ready for Pickup',CURRENT_TIMESTAMP),
('6','6','Cancelled',CURRENT_TIMESTAMP),
('7','7','On Hold',CURRENT_TIMESTAMP),
('8','8','Completed',CURRENT_TIMESTAMP),
('9','9','Refunded',CURRENT_TIMESTAMP);