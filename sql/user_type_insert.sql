INSERT INTO public.user_type (
id, description) VALUES (
'1'::integer, 'Talent'::character varying)
 returning id;

INSERT INTO public.user_type (
id, description) VALUES (
'2'::integer, 'Employer'::character varying)
 returning id;