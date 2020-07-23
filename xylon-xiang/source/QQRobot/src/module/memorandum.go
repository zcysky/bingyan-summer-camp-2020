package module

type Memorandum struct {
	UserId     uint   `json:"user_id"`
	GroupId    uint   `json:"group_id"`
	MemoId     string `json:"memo_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	Deadline   int64  `json:"deadline"`
	RemindTime int64  `json:"remind_time"`
}
