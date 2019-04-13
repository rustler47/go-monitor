package main

import(
	"fmt"
	"sync"
	"github.com/gocolly/colly"
)


func main() {
	fmt.Println("Press anything to start scrape")

	var pause string
	fmt.Scanln(&pause)//pause lol
	fmt.Println(pause)
	
	fmt.Println("Starting Concurrent Scraper")
	var wg sync.WaitGroup //---sync all go routines, wg.Done() increments wait counter
	var counter int
	
	for i, target := range GoScraper.Targets {
		newTask := &Task{Delay: 3000, ID: i, Target: target}
		
		if i > (len(GoScraper.Proxies)-1) {
			fmt.Println("Error: Ran out of proxies!!")
			fmt.Scanln(&pause)//pause lol
			fmt.Println(pause)
			return
		}
		newTask.Proxy = GoScraper.Proxies[i]
		counter++
		wg.Add(1) //--tell our waitgroup we have more waiting to do
		go startScrape(&wg, newTask)
	}
	
	wg.Wait() //--unblocks when all wg.Done()'s trigger
	
	fmt.Println(counter, "concurrent scrapes finished!")

	fmt.Scanln(&pause)//pause lol
	fmt.Println(pause)
}

func startScrape(wg *sync.WaitGroup, task *Task) {
	target := task.Target
	proxy  := task.Proxy
	collector, jar, ok := createCollector(proxy)
	if !ok {
		fmt.Println("Error: Failure creating collector!")
		return
	}
	jar.Put("safeGlobalVariable_ToBeSharedBetweenThreads", "Scrape Complete!")
	
	scan := "title"
	collector.OnHTML(scan, func(e *colly.HTMLElement) {
		fmt.Println(scan,":", e.Text)
	})

	collector.OnXML("*", func(e *colly.XMLElement) {
		//fmt.Println(e.Text)
	})
	
	//---final handler fired, after onHtml, onXML, etc
	collector.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		vari := jar.Get("safeGlobalVariable_ToBeSharedBetweenThreads")
		fmt.Println(vari)
	})
	
	collector.Visit(target)
	wg.Done()
}
