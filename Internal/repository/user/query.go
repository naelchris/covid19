package user

const (
	addUserQuery = `
		INSERT INTO user_data(
			name,
			email,
			password,
			dateofbirth,
			vaccinetype,
			healthstatus,
			createdat
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		) returning ID, name, email, password, dateofbirth, vaccineType, healthstatus, createdat
	`

	getUserQuery = `
		SELECT
			id,
			name,
			email,
			dateofbirth,
			vaccinetype,
			password,
			healthstatus
		FROM
			user_data
		WHERE
			email = $1
		LIMIT 1
	`

	updateUserQuery = `
		UPDATE 
			user_data
		SET
			name = $1,
			dateofbirth = $2,
			vaccinecertificate1 = $3,
			vaccinecertificate2 = $4,
			healthstatus = $5
		WHERE
			email = $6
		RETURNING email, name, dateofbirth, vaccinecertificate1, vaccinecertificate2, healthstatus, vaccinetype
	`
)
