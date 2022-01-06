CREATE TABLE public.promocions (
    id bigserial NOT NULL,
    descripcion varchar(100) NOT NULL,
    porcentaje numeric NOT NULL,
    fecha_inicio timestamptz NOT NULL,
    fecha_fin timestamptz NOT NULL,
    CONSTRAINT promocions_pkey PRIMARY KEY (id)
);

CREATE TABLE public.medicines (
    id bigserial NOT NULL,
    "name" varchar(50) NOT NULL,
    price numeric NOT NULL,
    "location" varchar(50) NOT NULL,
    CONSTRAINT medicines_pkey PRIMARY KEY (id)
);

CREATE TABLE public.invoices (
    id bigserial NOT NULL,
    fecha_creacion timestamptz NOT NULL,
    pago_total numeric NOT NULL,
    id_promotion int8 NOT NULL,
    CONSTRAINT invoices_pkey PRIMARY KEY (id)
);

ALTER TABLE public.invoices ADD CONSTRAINT fk_invoices_promotion FOREIGN KEY (id_promotion) REFERENCES public.promocions(id);

CREATE TABLE public.invoice_medicines (
    invoice_id int8 NOT NULL,
    medicine_id int8 NOT NULL,
    CONSTRAINT invoice_medicines_pkey PRIMARY KEY (invoice_id, medicine_id)
);

ALTER TABLE public.invoice_medicines ADD CONSTRAINT fk_invoice_medicines_invoice FOREIGN KEY (invoice_id) REFERENCES public.invoices(id);
ALTER TABLE public.invoice_medicines ADD CONSTRAINT fk_invoice_medicines_medicine FOREIGN KEY (medicine_id) REFERENCES public.medicines(id);