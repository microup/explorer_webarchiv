package app

import (
	"context"
	"flag"
	"log"

	"explorer_webarchiv/internal/web_domain"
)

func Run() {
	var targetDomain string
	var timeStamp string
	var maxWorkers int

	flag.StringVar(&targetDomain, "domain", "", "Specify the target domain (only lowercase)")
	flag.StringVar(&timeStamp, "timestamp", "", "Specify timestamp in the format:'yyyymmdd' (also: 'yyyy' > download only a specific year; 'yyyymm' > year and month; '2' or '1' > everything for the years past 20** or 19**")
	flag.IntVar(&maxWorkers, "workers", 10, "Specify the max workers (default=10)")

	flag.Parse()

	if targetDomain == "" || timeStamp == "" {
		log.Fatalf("Please provide both domain and timestamp.\n\n`explorer_webarchiv --domain=YOUR_SITE --timestamp=YOUR_DATE --workers=COUNT_WORKERS`\n\n--domain=YOUR_SITE Specify the target domain (only lowercase)\n--timestamp=YYYYMM Specify timestamp in the format:'yyyymmdd' (also: 'yyyy' > download only a specific year; 'yyyymm' > year and month; '2' or '1' > everything for the years past 20** or 19**\n--workers=XXX Specify the max workers (default=10)\n\nemail: contact@microup.ru\ntelegram: @microupp\nhttps://microup.ru")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	webDomain := web_domain.New(targetDomain, timeStamp)
	err := webDomain.Init()
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = webDomain.Download(ctx, maxWorkers)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Print("finished downloading")
}
