-- name: GetStation :one
SELECT * FROM stations
WHERE id = ? LIMIT 1;

-- name: ListStations :many
SELECT * FROM stations
ORDER BY name;

-- name: CreateStation :execresult
INSERT INTO stations (
    name
) VALUES (
    ?
);

-- name: UpdateStation :execresult
UPDATE stations
SET name = ?
WHERE id = ?;

-- name: DeteteStation :exec
DELETE FROM stations
WHERE id = ?;
