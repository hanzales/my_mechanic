DROP TABLE IF EXISTS comment CASCADE;
DROP TABLE IF EXISTS demand CASCADE;
DROP TABLE IF EXISTS users CASCADE;

--yorumların tutulduğu tablo
CREATE TABLE comment
(
    id serial PRIMARY KEY,
    message    VARCHAR(1024)            NOT NULL CHECK ( message <> '' ),
    likes      BIGINT                   DEFAULT 0,
    user_id int,
    demand_id int,
    active boolean,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

--talep tablosu
create table demand
(
    id         serial PRIMARY KEY,
    title      VARCHAR(1024)                                      NOT NULL CHECK ( title <> '' ),
    message    VARCHAR(5000)                                      NOT NULL CHECK ( message <> '' ),
    status     int                                                not null,
    user_id    int                                                not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

--kullanıcı tablosu
CREATE TABLE users
(
    id      serial PRIMARY KEY,
    first_name   VARCHAR(32)                 NOT NULL CHECK ( first_name <> '' ),
    last_name    VARCHAR(32)                 NOT NULL CHECK ( last_name <> '' ),
    email        VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
    role         VARCHAR(10)                 NOT NULL DEFAULT 'user',
    about        VARCHAR(1024)                        DEFAULT '',
    avatar       VARCHAR(512),
    phone_number VARCHAR(20),
    address      VARCHAR(250),
    city         VARCHAR(30),
    country      VARCHAR(30),
    gender       VARCHAR(20)                 NOT NULL DEFAULT 'male',
    postcode     INTEGER,
    birthday     DATE                                 DEFAULT NULL,
    active boolean,
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
    login_date   TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);