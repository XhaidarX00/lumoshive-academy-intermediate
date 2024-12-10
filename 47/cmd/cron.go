package cmd

import (
	"project_auth_jwt/controllers"

	"github.com/robfig/cron/v3"
)

func CronJob(ctx *controllers.Controller) error {
	c := cron.New()
	_, err := c.AddFunc("* * */1 * *", func() {
		ctx.Report.ReportService.Report.GenerateReport(0)
	})
	if err != nil {
		return err
	}

	c.Start()
	return nil
}
