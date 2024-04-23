CREATE TABLE public.application(
    id serial4 NOT NULL,
    job_id INT REFERENCES job(id)  ON DELETE CASCADE,
    applicant_id INT REFERENCES public.user(id)  ON DELETE CASCADE,
    status INT REFERENCES application_status(id)  ON DELETE CASCADE,
    CONSTRAINT application_pk PRIMARY KEY (id)
)