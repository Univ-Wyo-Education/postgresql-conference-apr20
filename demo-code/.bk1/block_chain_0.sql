
drop TRIGGER bk_block_trigger on bk_block;
drop TRIGGER send_notify_after_tg on bk_block;
drop TABLE if exists bk_block cascade ;
DROP FUNCTION if exists bk_block_tg() cascade;
DROP FUNCTION if exists send_notify_after_trig() cascade;
drop table if exists c_inventory_item cascade;

---------------------------------------------------------------------------------------

CREATE TABLE bk_block (
	id					uuid DEFAULT uuid_generate_v4() not null primary key,
	seq					serial not null,
	data				text not null,
	hash_of_this		text not null,
	prev_block_hash		text default '*' not null,
	prev_block_id		uuid not null,
	created 			timestamp default current_timestamp not null
);

create unique index bk_block_u1 on bk_block ( hash_of_this );
create index bk_block_p1 on bk_block ( created );
create index bk_block_p2 on bk_block ( prev_block_id );
create index bk_block_p3 on bk_block ( prev_block_hash );

---------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION bk_block_tg() RETURNS trigger AS $$
DECLARE
	l_seq_id bigint;
	l_id text;
	l_hash text;
BEGIN
    IF tg_op = 'UPDATE' or tg_op = 'DELETE' THEN
		RAISE EXCEPTION 'Can not UPDATE or DELETE from ' USING ERRCODE='20808';
        RETURN NULL;	-- Prevent All Updtes/Deletes
    ELSIF tg_op = 'INSERT' THEN
		NEW.hash_of_this = digest(CONCAT(NEW.data,NEW.prev_block_hash,NEW.hash_of_this,(NEW.prev_block_id::text)), 'sha256');
		IF new.id = 'e9d963e1-d2a2-4e35-8c37-22742008bce4'::uuid
			and NEW.data = 'genesis row' THEN
			NEW.prev_block_hash = NEW.hash_of_this;
			NEW.prev_block_id = NEW.id;
		ELSE
			select  "id", "hash_of_this" into l_id, l_hash
				from "bk_block"
				where "seq" in (
					select max("seq") from "bk_block"
				) ;
			NEW.prev_block_id = l_id;
			NEW.prev_block_hash = l_hash;
		END IF;
        RETURN NEW;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER bk_block_trigger
	BEFORE INSERT OR UPDATE OR DELETE ON bk_block
	FOR EACH ROW EXECUTE PROCEDURE bk_block_tg();

---------------------------------------------------------------------------------------
 CREATE OR REPLACE FUNCTION send_notify_after_trig()
 RETURNS TRIGGER AS $body$
 declare
 	msg text;
 begin
 	msg = '{"Cmd":"LogEvent","Hash":'||to_json(NEW.hash_of_this)||'}';
 	PERFORM pg_notify('events',msg);
 	return NEW;
 end;
 $body$ LANGUAGE plpgsql;
 
 CREATE TRIGGER send_notify_after_tg
 AFTER INSERT ON bk_block
     FOR EACH ROW EXECUTE PROCEDURE send_notify_after_trig();



---------------------------------------------------------------------------------------
-- Create the "genesis" block in the chian.
---------------------------------------------------------------------------------------
insert into bk_block (
	  "id"
	, "data"
) values (
	  'e9d963e1-d2a2-4e35-8c37-22742008bce4'::uuid
	, 'genesis row'
);















---------------------------------------------------------------------------------------

create table c_inventory_item (
	id					uuid DEFAULT uuid_generate_v4() not null primary key,
	item_id 			text not null ,
	item_count 			int default 1 not null ,
	user_id 			uuid not null ,
	created 			timestamp default current_timestamp not null
);

create index c_inventory_item_u1 on c_inventory_item ( item_id );




CREATE OR REPLACE FUNCTION notify_for_inventory_item()
RETURNS TRIGGER AS $$
declare
	msg text;
begin
	insert into bk_block ( data ) values ( row_to_json(NEW) );
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



---------------------------------------------------------------------------------------
insert into c_inventory_item ( item_id, user_id ) values ( '11223344', '7a955820-050a-405c-7e30-310da8152b6d');
