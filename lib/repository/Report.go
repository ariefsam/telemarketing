package repository

import (
	"log"
	"time"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/lib/timer"
	"github.com/keimoon/gore"
)

type Report struct{}

var conn *gore.Conn
var RedisDB int

func InitRedis() {
	if conn == nil {
		var err error
		conn, err = gore.Dial("localhost:6379") //Connect to redis server at localhost:6379
		gore.NewCommand("SELECT", RedisDB).Run(conn)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (r *Report) incrementByPeriod(telemarketerID, periodType, action string, amount int, t timer.Time) (currentScore int64, err error) {
	periodString := t.Format("20060102")
	key := "REPORT_" + action + "_" + telemarketerID + "_" + periodType + "_" + periodString
	log.Println("increment ", key)
	data, err := gore.NewCommand("INCRBY", key, amount).Run(conn)
	if err != nil {
		log.Println(err, "error closings")
		return 0, err
	}
	currentScore, err = data.Int()
	if err != nil {
		log.Println(err, "error amount snapshot")
		return
	}
	return
}

func (r *Report) dailySnapshotReport(telemarketerID, periodType, action string, amountSnapshot int64, t timer.Time, snapshotTime timer.Time) (err error) {
	periodString := t.Format("20060102")
	key := "REPORT_SNAPSHOT_" + action + "_" + telemarketerID + "_" + periodType + "_" + periodString + "_" + snapshotTime.BeginningDay().Format("20060102")

	_, err = gore.NewCommand("SET", key, amountSnapshot).Run(conn)
	if err != nil {
		log.Println(err, "error closings")
		return
	}
	return
}

func (r *Report) Increment(telemarketerID, action string, amount int, t timer.Time) (err error) {

	periodType := "daily"
	score, err := r.incrementByPeriod(telemarketerID, periodType, action, amount, t.BeginningDay())
	if err == nil {
		r.dailySnapshotReport(telemarketerID, periodType, action, score, t.BeginningDay(), t)
	}
	periodType = "weekly"
	score, err = r.incrementByPeriod(telemarketerID, periodType, action, amount, t.BeginningWeek())
	if err == nil {
		r.dailySnapshotReport(telemarketerID, periodType, action, score, t.BeginningWeek(), t)
		var i, j int64
		for i = int64(t.Weekday()); i <= 6; i++ {
			r.dailySnapshotReport(telemarketerID, periodType, action, score, t.BeginningWeek(), timer.NewFromTime(t.Add(time.Duration(j)*24*time.Hour)))

			j++
		}
	}
	periodType = "monthly"
	score, err = r.incrementByPeriod(telemarketerID, periodType, action, amount, t.BeginningMonth())
	if err == nil {
		var i, j int64
		for i = int64(t.Day()); i <= 31; i++ {
			r.dailySnapshotReport(telemarketerID, periodType, action, score, t.BeginningMonth(), timer.NewFromTime(t.Add(time.Duration(j)*24*time.Hour)))

			j++
		}
	}

	return
}

func (r *Report) getAction(telemarketerID, periodType, action string, t, snapshotTime timer.Time) (score int) {

	periodString := t.Format("20060102")
	key := "REPORT_SNAPSHOT_" + action + "_" + telemarketerID + "_" + periodType + "_" + periodString + "_" + snapshotTime.Format("20060102")
	// log.Println(key)
	temp, err := gore.NewCommand("GET", key).Run(conn)
	if err != nil {
		log.Println(err, "error closings")
		return
	}
	get, _ := temp.Int()
	score = int(get)

	// if score == 0 && (time.Now().Unix()-snapshotTime.Unix() < 86000) {
	// 	key := "REPORT_" + action + "_" + telemarketerID + "_" + periodType + "_" + periodString

	// 	temp, err := gore.NewCommand("GET", key).Run(conn)
	// 	if err != nil {
	// 		log.Println(err, "error closings")
	// 		return
	// 	}
	// 	get, _ := temp.Int()
	// 	score = int(get)
	// }
	return
}

func (r *Report) getByPeriod(telemarketerID string, periodType string, t, snapshotTime timer.Time) (tu entity.TargetUnit) {
	action := "CLOSING"
	tu.Closing = r.getAction(telemarketerID, periodType, action, t, snapshotTime)

	action = "BUYAMOUNT"
	tu.BuyAmount = r.getAction(telemarketerID, periodType, action, t, snapshotTime)

	action = "CALL"
	tu.Call = r.getAction(telemarketerID, periodType, action, t, snapshotTime)
	return
}

func (r *Report) Get(telemarketerID string, t timer.Time) (tp entity.TargetPeriod) {
	snapshotTime := t
	periodType := "daily"
	tp.Daily = r.getByPeriod(telemarketerID, periodType, t.BeginningDay(), snapshotTime)

	periodType = "weekly"
	tp.Weekly = r.getByPeriod(telemarketerID, periodType, t.BeginningWeek(), snapshotTime)

	periodType = "monthly"
	tp.Monthly = r.getByPeriod(telemarketerID, periodType, t.BeginningMonth(), snapshotTime)
	return
}
