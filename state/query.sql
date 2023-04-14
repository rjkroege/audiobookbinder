-- name: ListTracks :many
SELECT * FROM tracks
ORDER BY filename;

-- TODO(rjk): Can I name the parameters?

-- name: CreateTrack :exec
INSERT INTO tracks (
  author, booktitle, trackindex, year, filename, trackname
) VALUES (
  ?, ?, ?, ?, ?, ?
);
