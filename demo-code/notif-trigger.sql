CREATE OR REPLACE FUNCTION send_notify_after_trig()
RETURNS TRIGGER AS $body$
declare 
	msg text;
begin

	-- TODO - add so as to send new data to event.
	msg = '{"id":'||to_json(new.id)||'}';
	PERFORM pg_notify('events',msg);
		
	return null; 
end;
$body$ LANGUAGE plpgsql;

-- TODO - modify table to have a text field with the row of data in JSON format
-- TODO - add a timeout field (probably in text) to create a TTL, can be null
-- TODO - add a cmd/action field that tells if this is an inset/update/delete
-- TODO - add a new_id field that is the ID if it changes
CREATE TABLE send_notify (
	id 		text,
	whenmod	timestamp default current_timestamp not null
);

create index send_notify_p1 on send_notify ( whenmod );

CREATE TRIGGER send_notify_after_tg
AFTER INSERT ON send_notify
    FOR EACH ROW EXECUTE PROCEDURE send_notify_after_trig();


