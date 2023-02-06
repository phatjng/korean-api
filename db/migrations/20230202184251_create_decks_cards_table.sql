-- migrate:up
create table decks (
  id varchar(12) primary key,
  title varchar(50) not null
);

create table cards (
  id varchar(12) primary key,
  front varchar(255) not null,
  back varchar(255) not null,
  deck_id varchar(12) not null,
  foreign key (deck_id) references decks(id) on delete cascade
);

-- migrate:down
drop table decks;
drop table cards;
