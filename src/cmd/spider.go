package cmd

import (
	"figoxu/towerspider/common/config"
	"figoxu/towerspider/module/spider/crawl"
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
		user, password := config.TowerInfo()
		logrus.WithField("user", user).WithField("password", password).Println("startup")
		fetchUrl := "https://tower.im/teams/9e7383ee99514ab497e5654832f2d522/events/"
		actionLogSpider := crawl.NewActionLogSpider(fetchUrl, config.Ds)
		for i := 0; i < 30; i++ {
			actionLogSpider.More()
			actionLogSpider.ParseAndSave()
		}

		logrus.WithField("tp", taskTp).WithField("url", taskUrl).Println("start up")
	},
}

func init() {
	rootCmd.AddCommand(spiderCmd)
	spiderCmd.PersistentFlags().StringVar(&taskTp, "tp", "", "task type")
	spiderCmd.PersistentFlags().StringVar(&taskUrl, "url", "", "task url")
}
