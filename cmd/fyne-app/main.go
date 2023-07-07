package main

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/Tech-Dex/PartsPal/pkg/providers"
	"github.com/Tech-Dex/PartsPal/pkg/scraper"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const Timeout = 10 * time.Second

func main() {
	a := app.New()

	w := a.NewWindow("PartsPal")

	providerDealListB := binding.BindStringList(
		&[]string{},
	)
	bestDealB := binding.NewString()
	bestDealB.Set("No best deal found")
	var urlHl *url.URL
	bestDealOpenLinkBtn := widget.NewHyperlinkWithStyle("Go to shop", urlHl, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	titleW := widget.NewLabel("PartsPal")
	searchW := widget.NewEntry()
	searchW.SetPlaceHolder("Product code")

	var ctx context.Context
	var cancel context.CancelFunc
	searchBtn := widget.NewButton("Search", func() {
		if ctx != nil {
			providerDealListB.Set([]string{})
			cancel()
		}
		ctx, cancel = context.WithCancel(context.Background())
		if searchW.Text == "" {
			return
		}
		providerDealListB.Set([]string{})
		go func() {
			bd := &structs.BestDeal{
				Product: "",
				Price:   -1,
				Store:   "",
				Link:    "",
			}

			var wg sync.WaitGroup
			pipe := make(chan *structs.Deal, providers.SizeURLs)
			defer close(pipe)

			scraper.FindBestDeal(bd, &searchW.Text, &pipe, &wg, ctx)

			for {
				select {
				case deal := <-pipe:
					bdProduct, bdPrice, bdStore, bdLink := bd.Get()
					if bdPrice == -1 {
						bestDealB.Set("No best deal found")
					} else {
						bestDealB.Set(bdProduct + " - " + strconv.FormatFloat(bdPrice, 'f', 2, 64) + " RON @ " + bdStore)
						bestDealOpenLinkBtn.SetURLFromString(bdLink)
					}
					_, dPrice, dStore, _, err, notFound, unavailable, requstable := deal.Get()

					providerDeal := strconv.FormatFloat(dPrice, 'f', 2, 64) + " RON @ " + dStore
					if notFound {
						providerDeal = utils.ProductNotFoundMsg + " @ " + dStore
					}
					if unavailable {
						providerDeal = utils.IndisponibilMsg + " @ " + dStore
					}
					if requstable {
						providerDeal = utils.LaCerereMsg + " @ " + dStore
					}
					if err != "" {
						providerDeal = utils.GenericProviderErrorMsg + " @ " + dStore
					}
					providerDealListB.Append(providerDeal)
				case <-time.After(Timeout):
					wg.Wait()
					cancel()
					return
				}
			}
		}()
	})

	bestDealLbl := widget.NewLabel("Best Deal")
	bestDealW := widget.NewLabelWithData(bestDealB)
	bestDealW.Wrapping = fyne.TextWrapWord

	content := container.NewBorder(
		container.NewGridWithColumns(1,
			container.NewGridWithRows(1,
				container.NewCenter(titleW),
				searchW,
				searchBtn,
			),
		),
		nil, nil, nil,
		container.NewGridWithRows(1,
			container.NewGridWithColumns(1,
				container.NewMax(
					widget.NewListWithData(providerDealListB,
						func() fyne.CanvasObject {
							return widget.NewLabel("template")
						},
						func(i binding.DataItem, o fyne.CanvasObject) {
							o.(*widget.Label).Bind(i.(binding.String))
						}),
				),
			),
			container.NewGridWithColumns(1,
				container.NewBorder(
					bestDealLbl,
					nil, nil, nil,
					container.NewBorder(
						bestDealW, nil, nil, nil,
						container.NewBorder(
							bestDealOpenLinkBtn, nil, nil, nil, nil,
						),
					),
				),
			),
		),
	)

	w.SetContent(content)

	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()

}
