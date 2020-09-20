package repository

import (
	"log"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/lib/timer"
	"github.com/keimoon/gore"
)

type Report struct{}

var conn *gore.Conn

func InitRedis() {
	if conn == nil {
		var err error
		conn, err = gore.Dial("localhost:6379") //Connect to redis server at localhost:6379
		gore.NewCommand("SELECT", 15).Run(conn)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (r *Report) incrementByPeriod(telemarketerID, periodType, action string, amount int, t timer.Time) (err error) {
	periodString := t.Format("20060102")
	key := "REPORT_" + action + "_" + telemarketerID + "_" + periodType + "_" + periodString
	log.Println("increment ", key)
	_, err = gore.NewCommand("INCRBY", key, amount).Run(conn)
	if err != nil {
		log.Println(err, "error closings")
		return
	}
	return
}
func (r *Report) Increment(telemarketerID, action string, amount int, t timer.Time) (err error) {

	periodType := "daily"
	r.incrementByPeriod(telemarketerID, periodType, action, amount, t.BeginningDay())
	periodType = "weekly"
	r.incrementByPeriod(telemarketerID, periodType, action, amount, t.BeginningWeek())
	periodType = "monthly"
	r.incrementByPeriod(telemarketerID, periodType, action, amount, t.BeginningMonth())

	// periodString := t.BeginningDay().Format("20060102")
	// key := "REPORT_" + action + "_" + telemarketerID + "_" + periodType + "_" + periodString
	// log.Println("increment ", key)
	// _, err = gore.NewCommand("INCRBY", key, amount).Run(conn)
	// if err != nil {
	// 	log.Println(err, "error closings")
	// 	return
	// }

	// periodType = "weekly"
	// periodString = t.BeginningWeek().Format("20060102")
	// key = "REPORT_" + action + "_" + telemarketerID + "_" + periodType + "_" + periodString
	// log.Println("increment ", key)
	// _, err = gore.NewCommand("INCRBY", key, amount).Run(conn)
	// if err != nil {
	// 	log.Println(err, "error closings")
	// 	return
	// }

	// periodType = "monthly"
	// periodString = t.BeginningMonth().Format("20060102")
	// key = "REPORT_" + action + "_" + telemarketerID + "_" + periodType + "_" + periodString
	// log.Println("increment ", key)
	// _, err = gore.NewCommand("INCRBY", key, amount).Run(conn)
	// if err != nil {
	// 	log.Println(err, "error closings")
	// 	return
	// }

	return
}

func (r *Report) getAction(telemarketerID, periodType, action string, t timer.Time) (score int) {

	periodString := t.Format("20060102")
	key := "REPORT_" + action + "_" + telemarketerID + "_" + periodType + "_" + periodString
	log.Println("Get ", key)
	temp, err := gore.NewCommand("GET", key).Run(conn)
	if err != nil {
		log.Println(err, "error closings")
		return
	}
	get, _ := temp.Int()
	score = int(get)
	return
}

func (r *Report) getByPeriod(telemarketerID string, periodType string, t timer.Time) (tu entity.TargetUnit) {
	action := "CLOSING"
	tu.Closing = r.getAction(telemarketerID, periodType, action, t)

	action = "BUYAMOUNT"
	tu.BuyAmount = r.getAction(telemarketerID, periodType, action, t)

	action = "CALL"
	tu.Call = r.getAction(telemarketerID, periodType, action, t)
	return
}

func (r *Report) Get(telemarketerID string, t timer.Time) (tp entity.TargetPeriod) {
	periodType := "daily"
	tp.Daily = r.getByPeriod(telemarketerID, periodType, t.BeginningDay())

	periodType = "weekly"
	tp.Weekly = r.getByPeriod(telemarketerID, periodType, t.BeginningWeek())

	periodType = "monthly"
	tp.Monthly = r.getByPeriod(telemarketerID, periodType, t.BeginningMonth())
	log.Printf("%+v", tp)
	return
}
