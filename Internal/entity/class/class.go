package domain

type Class struct {
	ClassID            int64    `json:"class_id"`
	CurriculumID       int64    `json:"curriculum_id"`
	Title              string   `json:"class_title"`
	ImageURL           string   `json:"image_url"`
	AdditionalImageURL []string `json:"additional_image_url"`
	Rating             float64  `json:"class_rating"`
	MentorID           int64    `json:"mentor_id"`
	MentorName         int64    `json:"mentor_name"`
	MentorBIO          string   `json:"mentor_bio"`
	Price              int64    `json:"price"`
	StartDate          int64    `json:"start_date"` //unix time
}

type ClassUsecase interface {
	Fetch()
	GetByID()
	Update()
	GetByTitle()
	Store()
	Delete()
}

type ClassRepository interface {
	Fetch()
	GetByID()
	Update()
	GetByTitle()
	Store()
	Delete()
}
