package covid

const (
	getBatchCovidCaseQuery = `
		SELECT 
			*
		FROM 
			covid19_data
		LIMIT $1
		OFFSET $2
	`

	getCovidCaseQuery = `
		SELECT 
			Country,
			Province,
			confirmed,
			deaths,
			recovered,
			active
		FROM
			covid19_data
		WHERE 
			id = $1
	`

	addCaseQuery = `
		INSERT INTO covid19_data( 
			country,   
			countrycode,
			province, 
			city,
			citycode,
			lat,
			lon,
			confirmed, 
			deaths, 
			recovered,  
			active,
			date
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12
		) returning *
	`

	updateCaseQuery = `
		UPDATE
			covid19_data
		SET
			%s
		WHERE
			id=%d
	`
)
