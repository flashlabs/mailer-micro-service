create table if not exists message
(
    id          uuid                     not null
        constraint message_pk
            primary key,
    email       text                     not null,
    title       text                     not null,
    content     text                     not null,
    mailing_id  integer                  not null,
    insert_time timestamp with time zone not null
);

alter table message
    owner to postgres;

create index if not exists message_mailing_id_index
    on message (mailing_id);

create unique index if not exists message_email_mailing_id_uindex
    on message (email, mailing_id);
