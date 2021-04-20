
CREATE OR REPLACE FUNCTION t5()
RETURNS text AS $body$
declare 
	msg text;
begin
	perform pg_notify ( 'events', '{"cmd":"test-echo"}' );
	return 'Yep';
end;
$body$ LANGUAGE plpgsql;


select t5();

