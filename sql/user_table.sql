CREATE TABLE public.user(
    id serial4 NOT NULL,
    "username" varchar NULL,
    "email" varchar NULL,
    "password" varchar NULL,
    CONSTRAINT user_pk PRIMARY KEY (id),
    user_type INT REFERENCES user_type(id)
)