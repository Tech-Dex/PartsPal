package main

import (
	"context"
	"github.com/Tech-Dex/PartsPal/pkg/providers"
	"github.com/Tech-Dex/PartsPal/pkg/scraper"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/gorilla/websocket"
	"log"
	"math"
	"net/http"
	"sync"
	"time"
)

var up = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://localhost:80" || origin == "https://localhost:443" || origin == "http://localhost"
	},
}

const Timeout = 10 * time.Second

func scrapeSKU(w http.ResponseWriter, r *http.Request) {
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading websocket connection: %v", err)
	}
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading websocket message: %v", err)
			return
		}

		go func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			search := string(msg)
			bd := &structs.BestDeal{
				Product: "",
				Price:   -1,
				Store:   "",
				Link:    "",
			}

			var wg sync.WaitGroup
			pipe := make(chan *structs.Deal, providers.SizeURLs)
			defer close(pipe)

			scraper.FindBestDeal(bd, &search, &pipe, &wg, ctx)
			bdStoreTmp := ""
			for {
				select {
				case deal := <-pipe:
					bdProduct, bdPrice, bdStore, bdLink := bd.Get()
					if bdStoreTmp != bdStore && bdPrice != -1 {
						bdStoreTmp = bdStore
						err = conn.WriteJSON(
							&structs.DealJson{
								Type:        "bestDeal",
								Product:     bdProduct,
								Price:       math.Round(bdPrice*100) / 100,
								Store:       bdStore,
								Link:        bdLink,
								NotFound:    false,
								Unavailable: false,
								Requestable: false,
							})
					}

					dProduct, dPrice, dStore, dLink, dErr, notFound, unavailable, requstable := deal.Get()
					if dErr != "" {
						log.Printf("Error scraping %s: %s", dStore, dErr)
					}

					err = conn.WriteJSON(
						&structs.DealJson{
							Type:        "deal",
							Product:     dProduct,
							Price:       math.Round(dPrice*100) / 100,
							Store:       dStore,
							Link:        dLink,
							NotFound:    notFound,
							Unavailable: unavailable,
							Requestable: requstable,
						})
				case <-time.After(Timeout):
					wg.Wait()
					cancel()
					err := conn.WriteMessage(msgType, []byte("done"))
					if err != nil {
						log.Printf("Error writing websocket message: %v", err)
					}
					return
				}
			}
		}()

		if err != nil {
			log.Printf("Error writing websocket message: %v", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws/scrape", scrapeSKU)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started on port 3000")
}
