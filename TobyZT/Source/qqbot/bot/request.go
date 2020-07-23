package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(url string, body *bytes.Buffer) (res []byte, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Get(url string) (res []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Auth(url string, authKey string) (session string, err error) {
	jsonStr := fmt.Sprintf(`{"authKey": "%s"}`, authKey)
	buf, err := Post(url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return "", err
	}
	var res AuthResponse
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return "", err
	}
	if res.Code != 0 {
		return "", fmt.Errorf(session)
	}
	return res.Session, nil
}

func Verify(url string, session string, qq int) (err error) {
	jsonStr := fmt.Sprintf(`{"sessionKey": "%s", "qq": %d}`, session, qq)
	buf, err := Post(url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return err
	}
	var res VerifyResponse
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return err
	}
	if res.Code != 0 {
		return fmt.Errorf(res.Message)
	}
	return nil
}

func SendMessage(url string, msg Message) (msgID int, err error) {
	msgBuf := new(bytes.Buffer)
	err = json.NewEncoder(msgBuf).Encode(msg)
	if err != nil {
		return 0, err
	}
	resBuf, err := Post(url, msgBuf)
	if err != nil {
		return 0, err
	}
	var res MessageResponse
	err = json.Unmarshal(resBuf, &res)
	if err != nil {
		return 0, err
	}
	if res.Code != 0 {
		return 0, fmt.Errorf(res.Message)
	}
	return res.MessageID, nil
}
