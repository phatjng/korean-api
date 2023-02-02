CREATE TABLE IF NOT EXISTS "migrations" (version varchar(255) primary key);
CREATE TABLE decks (
  id varchar(12) primary key,
  title varchar(50)
);
CREATE TABLE cards (
  id varchar(12) primary key,
  front varchar(255) not null,
  back varchar(255) not null,
  deck_id varchar(12) not null,
  foreign key (deck_id) references decks(id) on delete cascade
);
-- Dbmate schema migrations
INSERT INTO "migrations" (version) VALUES
  ('20230202184251');
