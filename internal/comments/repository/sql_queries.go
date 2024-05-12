package repository

const (
	getCommentByID = `select id, message, likes,user_id,demand_id, created_at, updated_at from comment where id = $1`
)
