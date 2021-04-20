
---------------------------------------------------------------------------------------

CREATE SEQUENCE if not exists bc_most_recent_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;

drop table if exists bc_block ;

create table bc_block (
	  "id"					uuid DEFAULT uuid_generate_v4() not null primary key
	, "seq"					bigint DEFAULT nextval('bc_most_recent_seq'::regclass) NOT NULL 
	, "data"				text not null	-- insert : JSON data from row of table
	, "hash_of_this"		text 
	, "merkle_hash"			text 	-- Hash of Tranactions
	, "prev_block_hash"		text 
	, "prev_block_id"		uuid 
	, "created" 			timestamp default current_timestamp not null 						
);

create unique index bc_data_u1 on bc_block ( "hash_of_this" );
create unique index bc_data_u2 on bc_block ( "prev_block_id" );
create unique index bc_data_u3 on bc_block ( "prev_block_hash" );
create index bc_data_p1 on bc_block ( "created" );

---------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION bc_data_tg() RETURNS trigger AS $$
DECLARE
	l_seq_id bigint;
	l_id text;
	l_hash text;
BEGIN
    IF tg_op = 'UPDATE' or tg_op = 'DELETE' THEN
		RAISE EXCEPTION 'Can not UPDATE or DELETE from ' USING ERRCODE='20808';
        RETURN NULL;	-- Prevent All Updtes/Deletes
    END IF;
    IF tg_op = 'INSERT' THEN
		NEW.hash_of_this = digest(NEW.data||NEW.merkle_hash||NEW.prev_block_hash||(NEW.prev_block_id::text), 'sha256');
		-- if ithe genesis block!
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

CREATE TRIGGER bc_data_trigger
	BEFORE INSERT OR UPDATE OR DELETE ON bc_block 
	FOR EACH ROW EXECUTE PROCEDURE bc_data_tg();

---------------------------------------------------------------------------------------

-- Create the "genesis" block in the chian.
insert into bc_block (
	  "id"			
	, "data"				
) values (
	  'e9d963e1-d2a2-4e35-8c37-22742008bce4'::uuid
	, 'genesis row'
);















drop table if exists bc_tx ;

---------------------------------------------------------------------------------------
create table bc_tx (
	  "id"					uuid DEFAULT uuid_generate_v4() not null primary key
	, "block_id"			uuid NOT NULL	-- FK TODO
	, "tx_from"				uuid[] not null
	, "tx_to"				uuid[] not null
	, "tx_to_amt"			numeric[] not null
	, "tx_from_amt"			numeric not null
	, "tx_hash"				text 
	, "tx_signature"		text not null
	, "memo"				text not null
	, "smemo"				jsonb 
	, "words"				tsvector 
);

create index bc_tx_g1 on bc_tx using gin ( smemo ) where smemo is not null;
create index bc_tx_g2 on bc_tx using gin ( words ) where words is not null;

---------------------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION bc_tx_trigger() RETURNS trigger AS $$
BEGIN
    IF tg_op = 'INSERT' THEN
		-- attach pending Tx to the genesis block
	  	if new.block_id is null then			
			new.block_id = 'e9d963e1-d2a2-4e35-8c37-22742008bce4'::uuid;
		end if;
        RETURN NEW;
    ELSIF tg_op = 'UPDATE' THEN
		if new.block_id <> old.block_id then		
			if isMostRecent ( new.block_id )  
				or new.block_id = 'e9d963e1-d2a2-4e35-8c37-22742008bce4'::uuid
			then		
				new.id = old.id;
				new.tx_from = old.tx_from;
				new.tx_to = old.tx_to;
				new.tx_from_amt = old.tx_from_amt;
				new.tx_to_amt = old.tx_to_amt;
				new.tx_signature = old.tx_signature;
				new.tx_hash = digest(tx_searilize(NEW), 'sha256');
				RETURN NEW;
			else
				RAISE EXCEPTION 'Can not UPDATE Tx ' USING ERRCODE='20808';
				RETURN NULL;	-- Prevent 
			end if;
		end if;
        RETURN NULL;
    ELSIF tg_op = 'DELETE' THEN
		RAISE EXCEPTION 'Can not DELETE Tx ' USING ERRCODE='20808';
        RETURN NULL;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER example_table_trigger
	BEFORE INSERT or UPDATE or DELETE ON bc_tx 
	FOR EACH ROW EXECUTE PROCEDURE bc_tx_trigger();


---------------------------------------------------------------------------------------
--- Test Data
-- 
-- delete from t_output;
-- 
-- insert into bc_tx ( "name", "username" ) values ( 'Philip', 'pschlump' );
-- 
-- select count(1) from bc_tx;
-- select * from bc_block;
-- 
-- select output from t_output;
-- 
-- insert into bc_tx ( "name", "username" ) values ( 'Philip Schlump', 'pschlump@uwyo.edu' );
-- 
-- select count(1) from bc_tx;
-- select * from bc_block;
-- 
-- select output from t_output;
-- 
