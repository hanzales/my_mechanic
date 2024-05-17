package repository

const (
	getUserByEmail = `
					SELECT id,
						   first_name,
						   last_name,
						   email,
						   password,
						   role,
						   about,
						   avatar,
						   phone_number,
						   address,
						   city,
						   country,
						   gender,
						   postcode,
						   birthday,
						   active,
						   created_at,
						   updated_at,
						   login_date
					FROM users WHERE email = $1`

	userById = `
					SELECT id,
						   first_name,
						   last_name,
						   email,
						   password,
						   role,
						   about,
						   avatar,
						   phone_number,
						   address,
						   city,
						   country,
						   gender,
						   postcode,
						   birthday,
						   active,
						   created_at,
						   updated_at,
						   login_date
					FROM users WHERE id = $1`

	createUser = `INSERT INTO users (first_name, last_name, email, password, role, about, avatar, phone_number, address,
	               		city, gender, postcode, birthday,active, created_at, updated_at, login_date)
						VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'user'), $6, $7, $8, $9, $10, $11, $12, $13,true, now(), now(), now()) 
						RETURNING *`
)
