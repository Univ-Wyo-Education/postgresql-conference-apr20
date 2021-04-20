
drop TRIGGER bc_block_trigger on bk_block;
drop TRIGGER send_notify_after_tg on bk_block;
drop TABLE if exists bk_block cascade ;
drop index if exists bc_block_u1 ;
drop index if exists bc_block_u2 ;
drop index if exists bc_block_u3 ;
drop index if exists bc_block_p1 ;
DROP FUNCTION if exists bc_block_tg() cascade;
DROP FUNCTION if exists send_notify_after_trig() cascade;

---------------------------------------------------------------------------------------

CREATE TABLE bk_block (
	id 					uuid not null,
	seq					serial not null,
	data				text not null,
	hash_of_this		text not null,
	prev_block_hash		text default '*' not null,
	prev_block_id		uuid not null,
	created 			timestamp default current_timestamp not null
);

create unique index bc_block_u1 on bc_block ( hash_of_this );
create unique index bc_block_u2 on bc_block ( prev_block_id );
create unique index bc_block_u3 on bc_block ( prev_block_hash );
create index bc_block_p1 on bc_block ( created );

---------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION bc_block_tg() RETURNS trigger AS $$
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
				from "bc_block"
				where "seq" in (
					select max("seq") from "bc_block"
				) ;
			NEW.prev_block_id = l_id;
			NEW.prev_block_hash = l_hash;
		END IF;
        RETURN NEW;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER bc_block_trigger
	BEFORE INSERT OR UPDATE OR DELETE ON bc_block
	FOR EACH ROW EXECUTE PROCEDURE bc_block_tg();

---------------------------------------------------------------------------------------
--- CREATE OR REPLACE FUNCTION send_notify_after_trig()
--- RETURNS TRIGGER AS $body$
--- declare
--- 	msg text;
--- begin
--- 	msg = row_to_json(NEW);
--- 	PERFORM pg_notify('events',msg);
--- 	return NEW;
--- end;
--- $body$ LANGUAGE plpgsql;
--- 
--- CREATE TRIGGER send_notify_after_tg
--- AFTER INSERT ON bk_block
---     FOR EACH ROW EXECUTE PROCEDURE send_notify_after_trig();



---------------------------------------------------------------------------------------
-- Create the "genesis" block in the chian.
---------------------------------------------------------------------------------------
insert into bc_block (
	  "id"
	, "data"
) values (
	  'e9d963e1-d2a2-4e35-8c37-22742008bce4'::uuid
	, 'genesis row'
);




