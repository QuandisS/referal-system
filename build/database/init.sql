create table referers (
    id bigint generated always as identity,
    email varchar(255) unique not null,
    code varchar(255),
    primary key (id)
);

create table referals (
    email varchar(255) not null,
    referer_id bigint not null,
    primary key (email, referer_id),
    foreign key (referer_id) references referers (id) on delete cascade
);