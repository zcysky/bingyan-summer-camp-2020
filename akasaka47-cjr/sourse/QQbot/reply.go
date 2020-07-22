package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MessageModel struct {
	Tag   string
	Count int
}

func Switch(m *PostInfo, r *Respond) {
	if m.Message == "[CQ:at,qq=3135858576] 退出" || strings.Contains(m.Message, "退出"){
		M.Count = 0
		r.Reply = "退出" + M.Tag + "模式"
		M.Tag = "default"
	} else if m.Message == "[CQ:at,qq=3135858576] yygq" || m.Message == "yygq" {
		M.Tag = "阴阳怪气"
		r.Reply = "开启阴阳人模式"
	} else if m.Message == "备忘录" {
		M.Tag = "备忘录"
		RemindID = ""
		r.Reply = "开启备忘录模式\n1.新建备忘录\n2.查看所有备忘录\n3.删除指定备忘录"
	}
}

func YYGQ(m *PostInfo, r *Respond) {
	if M.Tag != "阴阳怪气" {
		return
	}
	M.Count++
	if strings.Contains(m.Message, "?") || strings.Contains(m.Message, "？") || strings.Contains(m.Message, "为什么") {
		r.Reply = RandomReply(QuestionAnswers)
		fmt.Println(r.Reply)
	} else if strings.Contains(m.Message, "爬") || strings.Contains(m.Message, "不要"){
		r.Reply = RandomReply(EmotionAnswers)
	} else {
		r.Reply = RandomReply(UniversalAnswers)
	}
	return
}

func Reminder(m *PostInfo, r *Respond, client *mongo.Client){
	if M.Tag == "备忘录" {
		if m.Message == "说明" {
			r.Reply = "开启备忘录模式\n1.新建备忘录\n2.查看所有备忘录\n3.删除指定备忘录"
			return
		}else if m.Message == "1" {
			n := new(Remind)
			for {
				n.ID = CreateRandomString(4)
				judge, _ := CheckID(client, n)
				if judge{
				}else{
					break
				}
			}
			RemindID = n.ID
			err := Insert(client, n)
			if err != nil {
				r.Reply = "备忘录添加失败"
			}else{
				M.Tag = "备忘录1.1"
				r.Reply = "备忘录标题是？"
			}
			return
		}else if m.Message == "2" {
			ReplyAll(client, m)
			return
		}else if m.Message == "3" {
			r.Reply = "要删除的备忘录的ID是？"
			M.Tag = "备忘录3"
			return
		}
	}
	if M.Tag == "备忘录1.1" {
		if ChangeTitle(client, m) != 0 {
			M.Tag = "备忘录1.2"
			r.Reply = "内容是？"
			return
		}
	}else if M.Tag == "备忘录1.2" {
		if ChangeContent(client, m) != 0 {
			M.Tag = "备忘录"
			r.Reply = "备忘录添加成功！"
			RemindID = ""
			return
		}
	}
	if M.Tag == "备忘录3" {
		if Delete(client, m) {
			M.Tag = "备忘录"
			r.Reply = "备忘录删除成功！"
			return
		}else{
			M.Tag = "备忘录"
			r.Reply = "未找到该备忘录"
			return
		}
	}
	return
}


func SendMessage (QQ int64, message string){
	url := "http://192.168.31.6:5700/send_private_msg?user_id="
	url = url + strconv.FormatInt(QQ, 10) + "&&message=" + message
	client := &http.Client{Timeout: 2 * time.Second}
	_, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateRandomString(len int) string {
	var container string
	//var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var str = "1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}