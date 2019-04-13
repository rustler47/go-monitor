package main

import(
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func collyOnRequest(r *colly.Request) {
	fmt.Println("Visiting", r.URL.String())
}

func collyOnResponse(r *colly.Response) {
	if !(r.StatusCode == 200) {
		fmt.Println("Error", r.StatusCode)//--stuff like error 404, etc
		return
	}
	fmt.Println("Response for", r.Request.URL.String())
	//fmt.Println(string(r.Body))
}

func createCollector(proxy string) (*colly.Collector, *colly.Context, bool) {
	//---Our trusty storage
	//--so we can loop stuff, carry data over safely
	Jar := colly.NewContext()
	
	//---Our trusty steed
	c := colly.NewCollector()
	c.IgnoreRobotsTxt = true
	
	//---Sets random user agent for our trusted stallion
	extensions.RandomUserAgent(c) 
	
	//--On each request (for headers, proxies etc)
	c.OnRequest(func(r *colly.Request)   {
		//jar example: something put in jar here will be avaliable to get in c.OnScraped 
		//using jar.Get
		collyOnRequest(r) 
	})
	
	//--On each response (url visit, post req etc)
	c.OnResponse(func(r *colly.Response) { collyOnResponse(r) })
	
	//--onError
	c.OnError(func(r *colly.Response, err error) {
		ret := fmt.Sprintf("Request URL: %s\nError: %v\n", r.Request.URL, err)
		fmt.Println(ret + fmt.Sprintf(" - Connection %d", r.StatusCode))
	})

	
	if !(proxy == "") { //dont set proxy if its ""
		err := c.SetProxy(proxy)
		if !(err == nil) { 
			return c, Jar, false
		}
	}
	return c, Jar, true
}