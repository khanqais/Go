 -- name: GetAuthor :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM product
ORDER BY name;


-- name: DeleteAuthor :exec
DELETE FROM product
WHERE id = $1;
