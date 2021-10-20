package crontab

var Work *work

func init() {
	Work = NewWork()

	Work.Add("cost", NewCost())
}
