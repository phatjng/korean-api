-- name: CreateCard :exec
insert into cards (
  id, front, back, deck_id
) values (
  ?, ?, ?, ?
);

-- name: GetCard :one
select * from cards
where id = ? limit 1;

-- name: GetAllCardsFromDeck :many
select * from cards
where deck_id = ?;

-- name: CreateDeck :exec
insert into decks (
  id, title
) values (
  ?, ?
);

-- name: GetDeck :one
select * from decks
where id = ? limit 1;

-- name: GetDeckByTitle :one
select * from decks
where title = ? limit 1;