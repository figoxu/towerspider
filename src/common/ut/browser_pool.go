package ut

import (
	"github.com/sclevine/agouti"
	"sync"
)

const selenium_path = "/Users/xujianhui/figospace/workspace_go/towerspider/resources/selenium-server-standalone-3.11.0.jar"

type Pool struct {
	sync.Mutex
	pws []*PageWarp
}

func NewPagePool() *Pool {
	return &Pool{
		pws: make([]*PageWarp, 0),
	}
}

func (p *Pool) Put(pw *PageWarp) {
	p.Lock()
	defer p.Unlock()
	defer Catch()
	pw.WorkFlag = false
	if pw.Page != nil {
		pw.Page.CloseWindow()
	}
}

func (p *Pool) Get() *PageWarp {
	defer Catch()
	p.Lock()
	defer p.Unlock()
	for _, pw := range p.pws {
		if !pw.WorkFlag {
			page, err := pw.WebDriber.NewPage(agouti.Browser("chrome"))
			if err != nil {
				pw.restart()
			} else {
				pw.Page = page
			}
			pw.WorkFlag = true
			return pw
		}
	}

	pw := &PageWarp{}
	pw.start()
	p.pws = append(p.pws, pw)
	return pw
}

type PageWarp struct {
	Page      *agouti.Page
	WebDriber *agouti.WebDriver
	WorkFlag  bool
}

func (p *PageWarp) start() {
	webDriver := agouti.Selendroid(selenium_path)
	for err := webDriver.Start(); err != nil; {
		webDriver = agouti.Selendroid(selenium_path)
		err = webDriver.Start()
	}
	p.WebDriber = webDriver
	page, err := p.WebDriber.NewPage(agouti.Browser("chrome"))
	for err != nil {
		page, err = p.WebDriber.NewPage(agouti.Browser("chrome"))
	}
	p.Page = page
}

func (p *PageWarp) restart() {
	p.release()
	p.start()
}

func (p *PageWarp) release() {
	stop := func() {
		defer Catch()
		p.WebDriber.Stop()
	}
	close := func() {
		defer Catch()
		p.Page.CloseWindow()
	}
	destory := func() {
		defer Catch()
		p.Page.Destroy()
	}
	close()
	destory()
	stop()
}
