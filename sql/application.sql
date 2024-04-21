CREATE TABLE public.application(
    id serial4 NOT NULL,
    job_id INT REFERENCES job(id),
    applicant_id INT REFERENCES public.user(id),
    status INT REFERENCES application_status(id),
    CONSTRAINT application_pk PRIMARY KEY (id)
)