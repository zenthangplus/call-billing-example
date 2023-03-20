create table calls
(
	id         int auto_increment,
	username   varchar(32)             not null,
	duration   bigint                  not null,
	created_at timestamp default NOW() not null,
	updated_at timestamp default NOW() not null,
	constraint calls_pk
		primary key (id)
);
