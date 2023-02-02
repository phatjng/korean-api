-- name: GetCard :one
select * from cards
where id = ? limit 1;