package module

type User struct {
	UserId string `json:"user_id"`
}

type CourseApprise struct {
	Good   int `json:"good"`
	Normal int `json:"normal"`
	Bad    int `json:"bad"`
	Total  int `json:"total"`
}

type Course struct {
	// the course id is to identify the course uniquely
	CourseId string `json:"course_id"`

	// the course class which are limited within six kinds
	CourseClass string `json:"course_class"`

	// the credit of the course
	CourseCredit int `json:"course_credit"`

	// which kind of student are the course oriented
	StudentOriented string `json:"student_oriented"`

	TeacherName string `json:"teacher_name"`

	// the starting time of the course, using the unix time stamp format
	StartTime int64 `json:"start_time"`

	// where does the course take place
	Position string `json:"position"`

	// the apprise of the course, divided into three parts, which are good, normal and bad
	// and the total apprise of the course is the means of them
	CourseApprise CourseApprise `json:"course_apprise"`

	// the method of the sign
	SigningMethod string `json:"signing_method"`

	// how the course ending, always paper
	CourseEndingMethod string `json:"course_ending_method"`
}

type CourseComment struct {
	// identify the comment uniquely
	CommentId string `json:"comment_id"`

	// the content of the comment
	Context string `json:"content"`

	// the course id of the comment
	CourseId string `json:"course_id"`

	// the user id of the comment
	UserId string `json:"user_id"`
}

type UserComment struct {
	TargetId    string `json:"target_id"`
	ThumbedNums int    `json:"thumbed_nums"`
	CourseComment
}
