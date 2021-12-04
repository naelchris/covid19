package user

const (
	addUserQuery = `
		INSERT INTO user_data(
			name,
			email,
			password,
			dateofbirth,
			vaccinetype,
			createdat
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		) returning *
	`

	getUserQuery = `
		SELECT
			id,
			name,
			email,
			dateofbirth,
			vaccinetype,
			healthstatus
		FROM
			user_data
		WHERE
			email = $1
		AND
			password = $2
		LIMIT 1
	`
)
