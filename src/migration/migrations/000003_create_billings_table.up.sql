create table billings
(
	id            int auto_increment,
	username      varchar(32)             not null,
	call_duration bigint                  not null,
	call_count    int                     not null,
	created_at    timestamp default NOW() not null,
	updated_at    timestamp default NOW() not null,
	constraint billings_pk
		primary key (id)
);
