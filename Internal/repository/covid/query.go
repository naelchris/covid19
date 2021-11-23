package covid

const (
	getBatchCovidCaseQuery = `
		SELECT 
			*
		FROM 
			report
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
			record
		WHERE 
			id = $1
	`

	addCaseQuery = `
		INSERT INTO report( 
			country,   
			country_code,
			province, 
			lat,
			lon,
			confirmed, 
			deaths, 
			recovered,  
			active
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
		) returning id
	`

	updateCaseQuery = `
		UPDATE
			report
		SET
			%s
		WHERE
			id=%d
	`
)
