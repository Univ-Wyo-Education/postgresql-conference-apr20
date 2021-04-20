
create table "c_auth_key" (
	"id"					uuid DEFAULT uuid_generate_v4() not null primary key,
	"key" 					text not null ,
	"user_id" 				uuid not null , -- references "c_user" ( "id" ),
	"valid_till" 			timestamp not null,
	"created" 				timestamp default current_timestamp not null,
	"updated" 				timestamp
);

create index "c_auth_key_u1" on "c_auth_key" ( "key" );

