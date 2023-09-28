CREATE TABLE calendar
(
    id serial not null unique,
    description text not null,
    date text not null,
    title text not null
)
