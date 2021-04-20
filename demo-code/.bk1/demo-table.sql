
CREATE OR REPLACE FUNCTION send_notify_after_trig()
RETURNS TRIGGER AS $body$
declare 
	msg text;
begin

	-- msg = '{"id":'||to_json(new.id)||'}';		-- xyzzy - use function to get full row	
	msg = row_to_json(NEW);
	PERFORM pg_notify('events',msg);
		
	return null; 
end;
$body$ LANGUAGE plpgsql;

-- TODO - modify table to have a text field with the row of data in JSON format
-- TODO - add a timeout field (probably in text) to create a TTL, can be null
-- TODO - add a cmd/action field that tells if this is an inset/update/delete
-- TODO - add a new_id field that is the ID if it changes
CREATE TABLE bk_block (
	id 			uuid not null,
	seq			searial not null,
	data		text not null,
	hash		text,
	prev_hash	uuid,
	prev_id		uuid,
	whenmod		timestamp default current_timestamp not null
);

create index send_notify_p1 on bk_block ( whenmod );

CREATE TRIGGER send_notify_after_tg
AFTER INSERT ON bk_block
    FOR EACH ROW EXECUTE PROCEDURE send_notify_after_trig();


-- xyzzy- chain blocks, hash them
















create table "c_inventory_item" (
	"id"					uuid DEFAULT uuid_generate_v4() not null primary key,
	"item_id" 				text not null ,
	"item_count" 			int default 1 not null ,
	"user_id" 				uuid not null , 
	"created" 				timestamp default current_timestamp not null,
);

create index "c_inventory_item_u1" on "c_inventory_item" ( "key" );




CREATE OR REPLACE FUNCTION notify_for_inventory_item()
RETURNS TRIGGER AS $$
declare 
	msg text;
begin
	insert into bk_block ( id ) values ( row_to_json(NEW) );
	return NEW; 
end;
$$ LANGUAGE plpgsql;

CREATE TRIGGER imutable_inventory_item_imutable
AFTER INSERT ON c_inventory_item
    FOR EACH ROW EXECUTE PROCEDURE notify_for_inventory_item();




CREATE OR REPLACE FUNCTION imutable_data()
RETURNS TRIGGER AS $$
begin
	return NULL;
end;
$$ LANGUAGE plpgsql;


CREATE TRIGGER notify_inventory_item_imutable
AFTER DELETE or UPDATE ON c_inventory_item
    FOR EACH ROW EXECUTE PROCEDURE imutable_data();

CREATE TRIGGER send_notify_imutable
AFTER DELETE or UPDATE ON send_notity
    FOR EACH ROW EXECUTE PROCEDURE imutable_data();

