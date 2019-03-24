package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Version string

var taskTp string
var taskUrl string

var spiderCmd = &cobra.Command{
	Use:   "spider",
	Short: "启动爬虫",
	Long:  `启动爬虫`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.WithField("tp", taskTp).WithField("url", taskUrl).Println("start up")
	},
}

func init() {
	rootCmd.AddCommand(spiderCmd)
	spiderCmd.PersistentFlags().StringVar(&taskTp, "tp", "", "task type")
	spiderCmd.PersistentFlags().StringVar(&taskUrl, "url", "", "task url")
}
