package main

//--global scraper object
//--if too few proxies, it won't let you continue. just an example for you to build on
var GoScraper = Scraper{
	Proxies: []string{
		"","","","",
		"","","","",
		"","","","",
		"","","","",
		"","","","",
		"","","","",
		"","","","",		"","","","",
		"","","","",
		"","","","",
		"","","","",
		"","","","",
		"","","","",		"","","","",
		"","","","",
		"","","","",
		"","","","",
		"","","","",
		"","","","",		"","","","",
		"","","","",
		"","","","",
		"","","","",
		"","","","",
		"","","","",
	},
	Targets: []string{
		"https://likelihood.us/",
		"https://shop.pharmabergen.no/",
		"https://kith.com/",
		"https://www.notre-shop.com/",
		"https://www.saintalfred.com/",
		"https://yeezysupply.com/",
		"https://wishatl.com/",
		"https://thebettergeneration.com/",
		"https://thenextdoor.fr/",
		"https://www.trophyroomstore.com/",
		"https://sneakerpolitics.com/",
		"https://shop.havenshop.com/",
		"https://shop.extrabutterny.com/",
		"https://shopnicekicks.com/",
		"https://www.bbbranded.com/",
		"https://undefeated.com/",
		"https://shop.travisscott.com/",
		"https://store.unionlosangeles.com/",
		"https://shoegallerymiami.com/",
		"https://crusoeandsons.com/",
		"https://www.onenessboutique.com/",
		"https://www.xhibition.co/",
		"https://feature.com/",
		"https://fearofgod.com/",
		"https://nrml.ca/",
		"https://www.hanon-shop.com/",
		"https://cncpts.com/",
		"https://creme321.com/",
		"https://www.pampamlondon.com/",
		"https://www.deadstock.ca/",
		"https://www.theclosetinc.com/",
		"https://properlbc.com/",
		"https://www.socialstatuspgh.com/",
		"https://bdgastore.com/",
		"https://us.bape.com/",
		"https://juicestore.com/",
		"https://www.blendsus.com/",
		"https://nowhere.ie/",
		"https://shopsizeusa.com/",
		"https://johngeigerco.com/",
		"https://rsvpgallery.com/",
		"https://www.solestop.com/",
		"https://offthehook.ca/",
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


