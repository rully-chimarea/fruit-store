CREATE TABLE public.fruits
(
    id bigserial,
    "fruit-name" character varying,
    "fruit-type" character varying,
    price numeric(102),
    quantity bigint,
    PRIMARY KEY (id)
);

ALTER TABLE public.fruits
    OWNER to postgres;