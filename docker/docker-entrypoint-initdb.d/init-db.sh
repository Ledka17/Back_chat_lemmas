#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    create table message
    (
        id           serial    not null
            constraint message_pk
                primary key,
        user_from_id integer   not null,
        user_to_id   integer   not null,
        text         text      not null,
        sent_date    timestamp not null
    );

    create table "user"
    (
        id         serial                not null
            constraint user_pk
                primary key,
        name       varchar               not null,
        email      varchar               not null,
        is_support boolean default false not null
    );

    create unique index user_email_uindex
        on "user" (email);
EOSQL