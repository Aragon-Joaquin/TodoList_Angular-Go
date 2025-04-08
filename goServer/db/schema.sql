CREATE TABLE IF NOT EXISTS tasks (
    id integer primary key autoincrement unique not null,
    name varchar(64) not null check(length(name) > 3),
    description varchar(128) null,
    status varchar(16) not null default('pending') CHECK (status IN ('done', 'pending', 'cancelled')),
    photo text null,
    hex_color char(7) not null default('#000000') CHECK(hex_color LIKE '#%')
);