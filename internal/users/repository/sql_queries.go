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
)
