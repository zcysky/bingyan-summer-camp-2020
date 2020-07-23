package memo

import (
	"QQRobot/src/module"
	"QQRobot/src/util"
	"context"
	"github.com/Logiase/gomirai"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"regexp"
	"strconv"
	"time"
)

const (
	NoteCtxPattern = "备忘录内容[:：]+.*"
	NoteDDLPattern = "\\[.*\\]DDL[:：]+.*"
	ShowALLPattern = "展示所有备忘录"
	DeletePattern  = "删除\\[.*\\]"
)

//regex the ctx and send the message
func Memorandum(b *gomirai.Bot, userId uint, ctx string, groupId ...uint) {
	match, pattern, err := regexCtx(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	if !match {
		return
	}

	switch pattern {
	case NoteCtxPattern:
		re := regexp.MustCompile(`.*[:：]+(.*)`)
		contents := re.FindSubmatch([]byte(ctx))
		content := string(contents[1])

		memo := module.Memorandum{
			UserId:  userId,
			Content: content,
		}

		if groupId != nil {
			memo.GroupId = groupId[0]
		}

		id, err := setMemorandum(memo)
		if err != nil {
			util.SendStateMessage(b, userId, util.FAILURE, groupId)
			return
		}

		util.SendStateMessage(b, userId, util.SUCCESS+"Your new memorandum id is "+id, groupId)

	case NoteDDLPattern:

		re := regexp.MustCompile(`\[(.*)\]`)
		results := re.FindSubmatch([]byte(ctx))
		memoId := string(results[1])

		// []DDL: year-month-day-hour-minute  advanced-minutes
		re = regexp.MustCompile(`\[\w*\]DDL[:：]+\s*(\d*)-(\d*)-(\d*)-(\d*)-(\d*)\s*(\d*)`)
		results = re.FindSubmatch([]byte(ctx))
		year, _ := strconv.Atoi(string(results[1]))
		month, _ := strconv.Atoi(string(results[2]))
		day, _ := strconv.Atoi(string(results[3]))
		hour, _ := strconv.Atoi(string(results[4]))
		minute, _ := strconv.Atoi(string(results[5]))
		advancedMinute, _ := strconv.Atoi(string(results[6]))

		timeObj := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local)
		ddl := timeObj.Unix()

		remindTime := ddl - int64(advancedMinute*60)

		memo := module.Memorandum{
			UserId:     userId,
			MemoId:     memoId,
			Deadline:   ddl,
			RemindTime: remindTime,
		}

		if groupId != nil {
			memo.GroupId = groupId[0]
		}

		err := setDDL(memo)
		if err != nil {
			util.SendStateMessage(b, userId, util.FAILURE, groupId)
			return
		}

		util.SendStateMessage(b, userId, util.SUCCESS, groupId)

	case ShowALLPattern:
		result, err := showAll(userId)
		if err != nil {
			util.SendStateMessage(b, userId, util.FAILURE, groupId)
			return
		}

		util.SendStateMessage(b, userId, result, groupId)

	case DeletePattern:

		re := regexp.MustCompile(`.*\[(\w{3,})\]`)
		results := re.FindSubmatch([]byte(ctx))
		memoId := string(results[1])

		err := deleteMemo(memoId)
		if err != nil {
			util.SendStateMessage(b, userId, "fail tto delete the memorandum", groupId)
		}

		util.SendStateMessage(b, userId, "Delete the memorandum successful!", groupId)
	}

}

func regexCtx(ctx string) (bool, string, error) {

	match, err := regexp.Match(NoteCtxPattern, []byte(ctx))
	if err != nil {
		return false, "", err
	}
	if match {
		return true, NoteCtxPattern, nil
	}

	match, err = regexp.Match(NoteDDLPattern, []byte(ctx))
	if err != nil {
		return false, "", err
	}
	if match {
		return true, NoteDDLPattern, nil
	}

	match, err = regexp.Match(ShowALLPattern, []byte(ctx))
	if err != nil {
		return false, "", err
	}
	if match {
		return true, ShowALLPattern, nil
	}

	match, err = regexp.Match(DeletePattern, []byte(ctx))
	if err != nil {
		return false, "", err
	}
	if match {
		return true, DeletePattern, nil
	}

	return false, "", nil
}

func setMemorandum(memo module.Memorandum) (string, error) {
	memo.CreateTime = time.Now().Unix()
	id := util.RandString()
	memo.MemoId = id

	// insert the memorandum into the mongo db
	_, err := module.MemoCol.InsertOne(context.Background(), bson.D{
		{"user_id", memo.UserId},
		{"group_id", memo.GroupId},
		{"memo_id", memo.MemoId},
		{"create_time", memo.CreateTime},
		{"content", memo.Content},
		{"deadline", memo.Deadline},
		{"remind_time", memo.RemindTime},
	})
	if err != nil {
		return "", err
	}

	return id, nil
}

func setDDL(memo module.Memorandum) error {
	var result bson.M

	filter := bson.D{
		{"memo_id", memo.MemoId},
	}

	err := module.MemoCol.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return err
	}

	memo.CreateTime = result["create_time"].(int64)
	memo.Content = result["content"].(string)

	update := bson.D{
		{"$set", bson.D{
			{"user_id", memo.UserId},
			{"memo_id", memo.MemoId},
			{"content", memo.Content},
			{"create_time", memo.CreateTime},
			{"deadline", memo.Deadline},
			{"group_id", memo.GroupId},
			{"remind_time", memo.RemindTime},
		}},
	}

	module.MemoCol.FindOneAndUpdate(context.Background(), filter, update)

	return nil
}

func showAll(id uint) (string, error) {
	cursor, err := module.MemoCol.Find(context.Background(), bson.D{{"user_id", id}})
	if err != nil {
		return "", err
	}

	var results []bson.M

	if err = cursor.All(context.Background(), &results); err != nil {
		return "", err
	}

	resultStr := ""
	for _, result := range results {
		resultStr += "[" + result["memo_id"].(string) + "]" + result["content"].(string) + "\n"
	}

	return resultStr, nil
}

func deleteMemo(memoId string) error {
	_, err := module.MemoCol.DeleteOne(context.Background(), bson.D{
		{"memo_id", memoId},
	})
	if err != nil {
		return err
	}
	return nil
}
