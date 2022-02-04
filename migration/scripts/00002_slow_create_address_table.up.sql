SELECT pg_sleep(10);

CREATE TABLE IF NOT EXISTS address
(
    id   bigserial,
    address varchar(64)
);

