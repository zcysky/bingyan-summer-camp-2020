package bot

type MessageChain struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Message struct {
	Session string         `json:"sessionKey"`
	Target  int            `json:"target"`
	Message []MessageChain `json:"messageChain"`
}

type AuthResponse struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

type VerifyResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type MessageResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"msg"`
	MessageID int    `json:"messageId"`
}

type SingleData struct {
	Type         string `json:"type"`
	MessageChain []struct {
		Type string `json:"type"`
		ID   int    `json:"id"`
		Time int    `json:"time"`
		Text string `json:"text"`
	} `json:"messageChain"`
	Sender struct {
		ID       int    `json:"id"`
		Nickname string `json:"nickname"`
		Remark   string `json:"remark"`
		Group    struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Permission string `json:"member"`
		}
	} `json:"sender"`
}

type FetchResponse struct {
	Code int          `json:"code"`
	Data []SingleData `json:"data"`
}

type Chat struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Reminder struct {
	ID      int    `bson:"id"`
	Type    string `bson:"type"`
	Due     int64  `bson:"due"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
	Advance int    `bson:"advance"`
	Gap     int    `bson:"gap"`
}
