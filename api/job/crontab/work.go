package crontab

import (
	"github.com/robfig/cron"
)

type work struct {
	c *cron.Cron
}

func NewWork() *work {
	w := work{}
	w.c = cron.New()
	return &w
}

func (w *work) Add(spec string, job cron.Job) {
	w.c.AddJob(spec, job)
}

func (w *work) Remove(job cron.Job) {
	//TODO
	//how to update job's spec
}

func (w *work) Start() {
	//启动计划任务
	w.c.Start()
	//TODO: 这里是不是不需要？ 阻塞主线程不退出，执行计划任务需要
	select {}
}

func (w *work) Stop() {
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	w.c.Stop()
}
