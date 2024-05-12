package repository

const (
	getCommentByID = `select id, message, likes,author_id,demand_id, created_at, updated_at from comments where id = $1`
)
