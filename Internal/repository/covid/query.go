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
		) on CONFLICT (date, country) DO UPDATE
		SET confirmed = $8, deaths = $9, recovered = $10,  active = $11
		RETURNING *
	`

	updateCaseQuery = `
		UPDATE
			covid19_data
		SET
			%s
		WHERE
			id=%d
	`

	getCasesByDay = `
		SELECT 
			confirmed,
			deaths,
			recovered,
			active,
			date
		FROM
			covid19_data
		WHERE 
			country = $1
		AND	
			date BETWEEN $2 AND $3
		ORDER BY
			date ASC
		`
	filterMonthCasesQuery = `
		SELECT 
			COALESCE(SUM(confirmed), 0) as confirmed, 
			COALESCE(SUM(deaths), 0) as deaths, 
			COALESCE(SUM(recovered), 0) as recovered,
			COALESCE(SUM(active), 0) as active
		FROM 
			covid19_data 
		WHERE 
			country = $1 
				AND 
			"date" between $2 and $3
	`
)
