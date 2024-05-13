package repository

const (
	getCommentByID = `SELECT id, message, likes,user_id,demand_id,active, created_at, updated_at FROM comment WHERE id = $1`
	deleteComment  = `UPDATE comment SET active = false, updated_at = now() WHERE id = $1`
	createComment  = `INSERT INTO comment(message,
											likes,
											user_id,
											demand_id,
											active,
											created_at,
											updated_at)
						VALUES ($1,
								0,
								$2,
								$3,
								true,
								now(),
								now()) RETURNING *`

	updateComment     = `UPDATE comment SET message = $1, updated_at = now() WHERE id = $2 RETURNING *`
	increaseLikeCount = `UPDATE  comment SET likes = likes + 1, updated_at = now() WHERE id = $1`
)
