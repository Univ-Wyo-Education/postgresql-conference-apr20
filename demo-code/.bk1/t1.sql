insert into c_auth_key ( key, user_id, valid_till ) values ( '1234', uuid_generate_v4(), current_timestamp + interval '1 hour' );
