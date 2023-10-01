CREATE DATABASE active_charity
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

create table if not exists public.activities
(
    id   serial
        primary key,
    name varchar(127),
    unit varchar(63)
);

create table if not exists public.users
(
    id               serial
        primary key,
    email            varchar(255),
    password         varchar(255),
    name             varchar(127),
    date_of_birthday varchar(127),
    organization_id  integer,
    wallet_id        integer,
    created_at       timestamp,
    updated_at       timestamp,
    fund_id          integer
);

create table if not exists public.organizations
(
    id         serial
        primary key,
    name       varchar(127),
    owner_id   integer,
    wallet_id  integer,
    created_at timestamp,
    updated_at timestamp
);

create table if not exists public.invite_codes
(
    id              serial
        primary key,
    organization_id integer,
    owner_id        integer,
    created_at      timestamp,
    updated_at      timestamp,
    code            varchar(63)
);

create table if not exists public.tokens
(
    id            serial
        primary key,
    user_id       integer,
    access_token  varchar(127),
    refresh_token varchar(127),
    exp           timestamp
);

create table if not exists public.transactions
(
    id             serial
        primary key,
    type           varchar(127),
    from_wallet_id integer,
    to_wallet_id   integer,
    coins          integer,
    rubles         integer,
    status         varchar(127),
    created_at     timestamp,
    updated_at     timestamp
);

create table if not exists public.wallets
(
    id         serial
        primary key,
    type       varchar(127),
    coins      integer,
    rubles     integer,
    created_at timestamp,
    updated_at timestamp
);

create table if not exists public.steps_history
(
    id          serial
        primary key,
    user_id     integer,
    steps       integer,
    activity_id integer,
    created_at  timestamp,
    updated_at  timestamp
);

create table if not exists public.funds
(
    id          serial
        primary key,
    description text,
    owner_id    integer,
    created_at  timestamp,
    updated_at  timestamp,
    name        varchar(127)
);

create table if not exists public.push_ups_history
(
    id          serial
        primary key,
    user_id     integer,
    repeats     integer,
    activity_id integer,
    created_at  timestamp,
    updated_at  timestamp
);

create table if not exists public.bench_press_history
(
    id          serial
        primary key,
    user_id     integer,
    repeats     integer,
    activity_id integer,
    created_at  timestamp,
    updated_at  timestamp
);

create table if not exists public.cycling_history
(
    id          serial
        primary key,
    user_id     integer,
    metres      integer,
    activity_id integer,
    created_at  timestamp,
    updated_at  timestamp
);

create table if not exists public.crunches_history
(
    id          serial
        primary key,
    user_id     integer,
    repeats     integer,
    activity_id integer,
    created_at  timestamp,
    updated_at  timestamp
);

create table if not exists public.pull_ups_history
(
    id          serial
        primary key,
    user_id     integer,
    repeats     integer,
    activity_id integer,
    created_at  timestamp,
    updated_at  timestamp
);

INSERT INTO public.activities (id, name, unit) VALUES (1, 'Шаги', '');
INSERT INTO public.activities (id, name, unit) VALUES (2, 'Подтягивания', 'Повторений');
INSERT INTO public.activities (id, name, unit) VALUES (3, 'Отжимания', 'Повторений');
INSERT INTO public.activities (id, name, unit) VALUES (4, 'Жим лёжа', 'Повторений');
INSERT INTO public.activities (id, name, unit) VALUES (5, 'Скручивания', 'Повторений');
INSERT INTO public.activities (id, name, unit) VALUES (6, 'Велопрогулка', 'Расстояние(в метрах)');