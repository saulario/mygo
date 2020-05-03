
create user myuser1 with login password 'mypassword1';
create user myuser2 with login password 'mypassword2';

create database mydb1 owner myuser1 template template0 encoding utf8 lc_collate 'POSIX' lc_ctype 'POSIX';
create database mydb2 owner myuser2 template template0 encoding utf8 lc_collate 'POSIX' lc_ctype 'POSIX';

// esto es por base de datos

drop table if exists table2;
drop table if exists table1;
drop table if exists table0;

create table table0 (
    id bigint not null primary key,
    texto varchar(80) not null,
    fechaHora timestamp,
    fecha date,
    hora time
);

create table table1 (
    id bigint not null primary key,
    table0id bigint references table0(id),
    texto varchar(40),
    tsmallint smallint,
    tinteger integer,
    tbiging bigint,
    tdecimal decimal(13,4),
    treal real,
    tdouble double precision
);

create table table2 (
    id bigint not null primary key,
    table0id bigint references table0(id),
    table1id bigint references table1(id),
    texto text
);

