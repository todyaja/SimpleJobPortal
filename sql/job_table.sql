CREATE TABLE public.job(
    id serial4 NOT NULL,
    "title" varchar NULL,
    "detail" varchar NULL,
    "requirement" varchar NULL,
    CONSTRAINT job_pk PRIMARY KEY (id),
    posted_by INT REFERENCES public.user(id)
)