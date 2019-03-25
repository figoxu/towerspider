package crawl

import (
	"figoxu/towerspider/common/config"
	"figoxu/towerspider/common/ut"
	"figoxu/towerspider/module/spider/service"
	"fmt"
	"github.com/quexer/utee"
	"github.com/sirupsen/logrus"
	"time"
)

var pagePool = ut.NewPagePool()

type ActionLogSpider struct {
	PageWrap  *ut.PageWarp
	MoreCount int
	ds        *config.DataSource
}

func NewActionLogSpider(httpUrl string, ds *config.DataSource) *ActionLogSpider {
	pw := pagePool.Get()
	pw.Page.Navigate(httpUrl)
	script := `
document.getElementById("email").value = "%s";
document.getElementsByName("password")[0].value="%s";
document.getElementById("btn-signin").click()
`
	user, password := config.TowerInfo()
	script = fmt.Sprintf(script, user, password)

	pw.Page.RunScript(script, map[string]interface{}{}, map[string]interface{}{})
	return &ActionLogSpider{
		PageWrap: pw,
		ds:       ds,
	}

}

func (p *ActionLogSpider) More() {
	p.MoreCount = p.MoreCount + 1
	y := p.MoreCount * 3000
	script := fmt.Sprint("window.scrollTo(", 0, ",", y, ")")
	logrus.Println(script)
	err := p.PageWrap.Page.RunScript(script, map[string]interface{}{}, map[string]interface{}{})
	utee.Chk(err)
	time.Sleep(time.Second * time.Duration(3))
	logrus.Println("获取更多")
}

func (p *ActionLogSpider) ParseAndSave() {
	elements, err := p.PageWrap.Page.All(".event-common").Elements()
	utee.Chk(err)
	for _, element := range elements {
		service.ActionLog(element, p.ds).Save()
	}
	logrus.Println("解析并保存页面")
}
