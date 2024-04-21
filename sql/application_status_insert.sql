INSERT INTO public.application_status (
id, description) VALUES (
'1'::integer, 'OnGoing'::character varying)
 returning id;
 INSERT INTO public.application_status (
id, description) VALUES (
'2'::integer, 'Interview'::character varying)
 returning id;
 INSERT INTO public.application_status (
id, description) VALUES (
'3'::integer, 'Accepted'::character varying)
 returning id;
 INSERT INTO public.application_status (
id, description) VALUES (
'4'::integer, 'Rejected'::character varying)
 returning id;