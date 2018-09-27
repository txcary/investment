package crawler

type Strategy interface {
	CrawlNeeded(id string)(bool)
	GetUrl() (string)
	Process(interface{}) error
}

