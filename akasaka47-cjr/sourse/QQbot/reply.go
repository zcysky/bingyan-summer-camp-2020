package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MessageModel struct {
	Tag   string
	Count int
}

func Switch(m *PostInfo, r *Respond) {
	if m.Message == "[CQ:at,qq=3135858576] 退出" || strings.Contains(m.Message, "退出") {
		M.Count = 0
		r.Reply = "退出" + M.Tag + "模式"
		M.Tag = "default"
	} else if m.Message == "[CQ:at,qq=3135858576] yygq" || m.Message == "yygq" {
		M.Tag = "阴阳怪气"
		r.Ban = true
		r.BanDuration = 1
		r.Reply = "开启阴阳人模式"
	} else if m.Message == "备忘录" {
		M.Tag = "备忘录"
		RemindID = ""
		r.Reply = "开启备忘录模式\n1.新建备忘录\n2.查看所有备忘录\n3.修改指定备忘录\n4.设定备忘录提醒时间\n5.删除指定备忘录"
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
	} else if strings.Contains(m.Message, "爬") || strings.Contains(m.Message, "不要") {
		r.Reply = RandomReply(EmotionAnswers)
	} else {
		r.Reply = RandomReply(UniversalAnswers)
	}
	return
}

func Reminder(m *PostInfo, r *Respond, client *mongo.Client, re redis.Conn) {
	if m.Message == "0" && strings.Contains(M.Tag, "备忘录4.2") {
		M.Tag = "备忘录"
		r.Reply = "1.新建备忘录\n2.查看所有备忘录\n3.修改指定备忘录\n4.设定备忘录提醒时间\n5.删除指定备忘录"
		return
	}
	switch M.Tag {
	case "备忘录":
		switch m.Message {
		case "说明":
			r.Reply = "1.新建备忘录\n2.查看所有备忘录\n3.修改指定备忘录\n4.设定备忘录提醒时间\n5.删除指定备忘录"
		case "1":
			n := new(Remind)
			for {
				n.ID = CreateRandomString(4)
				judge, _ := CheckID(client, n.ID)
				if judge {
				} else {
					break
				}
			}
			RemindID = n.ID
			n.QQ = m.UserId
			err := Insert(client, n)
			if err != nil {
				r.Reply = "备忘录添加失败"
			} else {
				M.Tag = "备忘录1.1"
				r.Reply = "备忘录标题是？"
			}
		case "2":
			r.Reply = ""
			if !ReplyAll(client, m) {
				r.Reply = "您还没有建立备忘录！\n"
			}
			r.Reply = r.Reply + "发送“说明”可以继续操作"
		case "3":
			M.Tag = "备忘录3.1"
			r.Reply = "要修改的备忘录ID是？"
		case "4":
			M.Tag = "备忘录4.1"
			r.Reply = "要提醒的备忘录ID是？"
		case "5":
			r.Reply = "要删除的备忘录的ID是？"
			M.Tag = "备忘录4"
		default:
			r.Reply = "请输入要使用功能的序号！"
		}
	case "备忘录1.1":
		if ChangeTitle(client, m) != 0 {
			M.Tag = "备忘录1.2"
			r.Reply = "内容是？"
		}
	case "备忘录1.2":
		if ChangeContent(client, m) != 0 {
			M.Tag = "备忘录"
			r.Reply = "备忘录添加成功！\n发送“说明”可以继续操作"
			RemindID = ""
		}
	case "备忘录3.1":
		RemindID = m.Message
		M.Tag = "备忘录3.2"
		r.Reply = "要修改？\n1.标题\n2.目录"
	case "备忘录3.2":
		switch m.Message {
		case "1":
			M.Tag = "备忘录3.2.1"
			r.Reply = "修改为？"
		case "2":
			M.Tag = "备忘录3.2.2"
			r.Reply = "修改为？"
		default:
			r.Reply = "请输入要使用功能的序号！"
		}
	case "备忘录3.2.1":
		if ChangeTitle(client, m) != 0 {
			M.Tag = "备忘录"
			r.Reply = "标题修改成功！\n发送“说明”可以继续操作"
			RemindID = ""
		} else {
			r.Reply = "标题修改失败\n发送“说明”可以继续操作"
		}
	case "备忘录3.2.2":
		if ChangeContent(client, m) != 0 {
			M.Tag = "备忘录"
			r.Reply = "内容修改成功！\n发送“说明”可以继续操作"
			RemindID = ""
		} else {
			r.Reply = "内容修改失败\n发送“说明”可以继续操作"
		}
	case "备忘录4.1":
		RemindID = m.Message
		judge, _ := CheckID(client, RemindID)
		if !judge {
			r.Reply = "未找到该备忘录！\n请重新发送ID\n发送“0”退出提醒功能"
			RemindID = ""
			return
		}
		M.Tag = "备忘录4.2.1"
		r.Reply = "在哪一天？\n（例如：2020-01-01）"
	case "备忘录4.2.1":
		reg := regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2}`)
		str := reg.FindAllString(m.Message, -1)
		if str == nil {
			r.Reply = "发送格式错误，请重新发送\n发送“0”退出提醒功能"
			return
		}
		err := SetRedis(re, "day", str[0])
		if err != nil {
			fmt.Println(err)
		}
		M.Tag = "备忘录4.2.2"
		r.Reply = "在几点钟？\n（发送格式：00:00:00，使用24小时制）"
	case "备忘录4.2.2":
		reg := regexp.MustCompile(`：`)
		m.Message = reg.ReplaceAllString(m.Message, ":")
		reg = regexp.MustCompile(`[0-9]{2}:[0-9]{2}:[0-9]{2}`)
		str := reg.FindAllString(m.Message, -1)
		if str == nil {
			r.Reply = "发送格式错误，请重新发送\n发送“0”退出提醒功能"
			return
		}

		Time := FindRedis(re, "day")
		if Time != "" {
			Time = Time + " " + str[0]
		}
		fmt.Println(Time)
		t, _ := time.Parse("2006-01-02 15:04:05", Time)
		if ChangeTime(client, t.Unix()-28800) == 0 { //转换时区
			r.Reply = "修改时间失败"
		}
		DelRedis(re, "day")

		judge, result := CheckID(client, RemindID)
		if judge {
			go DDLReply(result)
		}

		r.Reply = "需要提前提醒吗？\n（回复“是”则提前提醒\n回复其它的内容则不会提前提醒）"
		M.Tag = "备忘录4.3.1"

	case "备忘录4.3.1":
		if m.Message != "是" {
			r.Reply = "提醒时间设定成功！\n发送“说明”可继续使用备忘录功能。"
			M.Tag = "备忘录"
			return
		}
		r.Reply = "提前多长时间？\n可发送格式：\n1天 或 1小时 或 1分钟"
		M.Tag = "备忘录4.3.2"
	case "备忘录4.3.2":
		reg := regexp.MustCompile(`[0-9]{1,2}天|[0-9]{1,2}小时|[0-9]{1,2}分钟`)
		str := reg.FindAllString(m.Message, -1)
		if str == nil {
			r.Reply = "发送格式错误，请重新发送\n发送“0”退出提醒功能"
			return
		}
		var t int64
		if strings.Contains(str[0], "天") {
			reg := regexp.MustCompile(`[0-9]{1,2}`)
			st := reg.FindAllString(str[0], -1)
			i, err := strconv.ParseInt(st[0], 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			t = i * 3600 * 24
		} else if strings.Contains(str[0], "小时") {
			reg := regexp.MustCompile(`[0-9]{1,2}`)
			st := reg.FindAllString(str[0], -1)
			i, err := strconv.ParseInt(st[0], 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			t = i * 3600
		} else if strings.Contains(str[0], "分钟") {
			reg := regexp.MustCompile(`[0-9]{1,2}`)
			st := reg.FindAllString(str[0], -1)
			i, err := strconv.ParseInt(st[0], 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			t = i * 60
		}
		err := SetRedis(re, "ahead", strconv.FormatInt(t, 10))
		if err != nil {
			fmt.Println(err)
		}
		r.Reply = "间隔多长时间提醒？\n可发送格式：\n1小时 或 1分钟\n回复“否”则不会间隔提醒"
		M.Tag = "备忘录4.3.3"
	case "备忘录4.3.3":
		var intval int64
		if m.Message == "否" {
			intval = 0
		} else {
			reg := regexp.MustCompile(`[0-9]{1,2}小时|[0-9]{1,2}分钟`)
			str := reg.FindAllString(m.Message, -1)
			if str == nil {
				r.Reply = "发送格式错误，请重新发送\n发送“0”退出提醒功能"
				return
			}
			if strings.Contains(str[0], "小时") {
				reg := regexp.MustCompile(`[0-9]{1,2}`)
				st := reg.FindAllString(str[0], -1)
				i, err := strconv.ParseInt(st[0], 10, 64)
				if err != nil {
					fmt.Println(err)
				}
				intval = i * 3600
			} else if strings.Contains(str[0], "分钟") {
				reg := regexp.MustCompile(`[0-9]{1,2}`)
				st := reg.FindAllString(str[0], -1)
				i, err := strconv.ParseInt(st[0], 10, 64)
				if err != nil {
					fmt.Println(err)
				}
				intval = i * 60
			}
		}

		ah, err := strconv.ParseInt(FindRedis(re, "ahead"), 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		judge, result := CheckID(client, RemindID)
		if judge {
			go AheadReply(result, ah, intval)
		}

		r.Reply = "设定成功！\n发送“说明”可继续使用备忘录功能。"
		M.Tag = "备忘录"

	case "备忘录5":
		if Delete(client, m) {
			M.Tag = "备忘录"
			r.Reply = "备忘录删除成功！\n发送“说明”可以继续操作"
			RemindID = ""
		} else {
			M.Tag = "备忘录"
			r.Reply = "未找到该备忘录"
		}
	}
	return
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

func SendPrivateMessage(UserId int64, message string) {
	url := Url + T.Private.Use + T.Private.UserId
	url = url + strconv.FormatInt(UserId, 10) + T.Private.Message + message
	client := &http.Client{Timeout: 2 * time.Second}
	_, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
}

func SendGroupMessage(GroupId int64, message string) {
	url := Url + T.Group.Use + T.Group.GroupId
	url = url + strconv.FormatInt(GroupId, 10) + T.Group.Message + message
	client := &http.Client{Timeout: 2 * time.Second}
	_, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
}

func DDLReply(result *Remind) {
	for {
		if result.DDL <= time.Now().Unix() {
			break
		}
		fmt.Println(result.DDL)
		fmt.Println(time.Now().Unix())
		fmt.Println()
		time.Sleep(time.Second)
	}
	SendPrivateMessage(result.QQ, "备忘录提醒的时间到了！")
	SendPrivateMessage(result.QQ, "标题："+result.Title)
	SendPrivateMessage(result.QQ, "内容："+result.Content)
	return
}

func AheadReply(result *Remind, ahead int64, interval int64) {
	ah := result.DDL - ahead
	for {
		if ah > time.Now().Unix() {
			time.Sleep(time.Second)
		} else {
			SendPrivateMessage(result.QQ, "备忘录提醒的时间到了！")
			SendPrivateMessage(result.QQ, "标题："+result.Title)
			SendPrivateMessage(result.QQ, "内容："+result.Content)
			time.Sleep(time.Duration(interval * 1000000000))
			if result.DDL <= time.Now().Unix() {
				break
			}
		}
	}
	return
}
