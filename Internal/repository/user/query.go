package user

const (
	addUserQuery = `
		INSERT INTO user_data(
			name,
			email,
			password,
			dateofbirth,
			lat,
			lng,
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
			$7,
			$8,
			$9
		) returning ID, name, email, password, dateofbirth, lat, lng, vaccineType, healthstatus, createdat
	`

	getUserQuery = `
		SELECT
			id,
			name,
			email,
			dateofbirth,
			lat,
			lng,
			vaccinetype,
			password,
			healthstatus,
			createdat,
			updatedat
		FROM
			user_data
		WHERE
			email = $1
		LIMIT 1
	`

	validateLoginQuery = `
		SELECT
			id,
			name,
			email,
			dateofbirth,
			lat,
			lng,
			vaccinetype,
			password,
			healthstatus
		FROM
			user_data
		WHERE
			email = $1
		AND
			password = $2
		LIMIT 1
	`

	updateUserQuery = `
		UPDATE 
			user_data
		SET
			name = $1,
			dateofbirth = $2,
			lat = $3,
			lng = $4,
			vaccinecertificate1 = $5,
			vaccinecertificate2 = $6,
			healthstatus = $7
		WHERE
			email = $8
		RETURNING email, name, dateofbirth, lat, lng, vaccinecertificate1, vaccinecertificate2, healthstatus, vaccinetype
	`
)
