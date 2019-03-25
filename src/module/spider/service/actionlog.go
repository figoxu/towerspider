package service

import (
	"figoxu/towerspider/common/config"
	"figoxu/towerspider/common/db/model"
	"figoxu/towerspider/common/db/mydao"
	"figoxu/towerspider/common/ut"
	"github.com/quexer/utee"
	"github.com/sclevine/agouti/api"
	"github.com/sirupsen/logrus"
)

var existMap = utee.SyncMap{}

type ActionLogService struct {
	element *api.Element
	ds      *config.DataSource
}

func ActionLog(ele *api.Element, ds *config.DataSource) *ActionLogService {
	return &ActionLogService{
		element: ele,
		ds:      ds,
	}
}

func (p *ActionLogService) Save() {

	id := p.attrText("id")
	if _, existFlag := existMap.Get(id); existFlag {
		return
	}
	existMap.Put(id, true)

	actor := p.parseText(".event-actor")
	action := p.parseText(".event-action")
	text := p.parseText(".event-text")
	sysName := p.attrText("data-ancestor-name")
	resourceUrl := p.assetHref()
	create_at, err := p.selAttr(".event-created-at", "data-created-at")
	utee.Chk(err)

	logrus.WithField("actor", actor).
		WithField("action", action).
		WithField("text", text).
		WithField("sysName", sysName).
		WithField("href", resourceUrl).
		WithField("create_at", create_at).
		WithField("id", id).Println("获取信息")

	actionLog := &model.ActionLog{
		EventId:     id,
		Actor:       actor,
		Action:      action,
		Text:        text,
		SysName:     sysName,
		ResourceUrl: resourceUrl,
	}
	timeText := ut.ClearTowerZoneInfo(create_at)
	actionLog.CreatedAt = ut.FormatTime(ut.TOWER_TIME_FMT, timeText)
	actionLogDao := mydao.ActionLog(p.ds.Mysql)
	actionLogDao.Insert(actionLog)
}

func (p *ActionLogService) parseText(sel string) string {
	ele, err := p.element.GetElement(NewCssSelector(sel))
	utee.Chk(err)
	text, err := ele.GetText()
	utee.Chk(err)
	return text
}
func (p *ActionLogService) attrText(name string) string {
	v, err := p.element.GetAttribute(name)
	utee.Chk(err)
	return v
}
func (p *ActionLogService) selAttr(sel, attr string) (string, error) {
	ele, err := p.element.GetElement(NewCssSelector(sel))
	if err != nil {
		return "", err
	}
	v, err := ele.GetAttribute(attr)
	return v, err
}
func (p *ActionLogService) assetHref() string {
	href, err := p.selAttr(".check_item-rest", "href")
	if err != nil {
		href, err = p.selAttr(".todo-rest", "href")
		if err != nil {
			href, err = p.selAttr(".emphasize", "href")
			if err != nil {
				logrus.WithField("err", true).Println(p.element.GetText())
				utee.Chk(err)
			}
		}
	}
	return href
}

func NewCssSelector(v string) api.Selector {
	return api.Selector{"css selector", v}
}
