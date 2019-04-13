package main

//--global scraper object
//--if too few proxies, it won't let you continue. just an example for you to build on
var GoScraper = Scraper{
	Proxies: []string{
		"",
		"",
		"",
		"",
		"",
	},
	Targets: []string{
		"https://undefeated.com/",
		"https://kith.com/",
		"https://nrml.ca/",
		"https://bdgastore.com",
		"https://cncpts.com/",
	},
}

type Scraper struct {
	Targets []string
	Proxies []string
}

type Task struct {
	ID     int
	Delay  int
	Target string
	Proxy  string
}


