package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	aw "github.com/deanishe/awgo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/lvliangxiong/alfred_chrome/model"
)

var histories []model.History

func showAllHistoryItems() {
	if enableFilter {
		histories = filter(histories)
	}

	for _, h := range histories {
		item := wf.NewItem(h.Title).
			Subtitle(h.Url).
			Arg(h.Url).
			UID(h.Url).
			Valid(true).
			Match(h.Url + "" + h.Title)

		if iconId, ok := mappings[h.Url]; ok {
			item.Icon(&aw.Icon{
				Value: fmt.Sprintf("icons/%010d.png", iconId),
			})
		} else {
			item.Icon(iconChrome)
		}
	}
}

func filter(hs []model.History) []model.History {
	urls := make(map[string]struct{}, len(hs))

	i := 0
	for _, h := range hs {
		url, err := url.Parse(h.Url)
		if err != nil {
			continue
		}

		url.Fragment = ""
		url.RawQuery = ""
		_, ok := urls[url.String()]
		if ok {
			continue
		}

		hs[i] = h
		urls[url.String()] = struct{}{}
		i++
	}

	return hs[0:i]
}

func historyInit() {
	needRefresh := false
	if time.Since(status.HistoryLastRefreshTime) >= config.History.RefreshInterval {
		needRefresh = true
	}

	if !needRefresh {
		_, err := os.Stat("History")
		if os.IsNotExist(err) {
			needRefresh = true
		}
	}

	if needRefresh {
		defer func() {
			err := status.save()
			if err != nil {
				log.Printf("refresh status saved failed, err: %#v", err)
			}
		}()

		if err := copyFile(config.getProfileHistoryFilePath(), "History"); err != nil {
			log.Fatal(err)
		}
	}

	db, err := gorm.Open(sqlite.Open("History"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	histories = fetchHistory(db)

	if needRefresh {
		status.HistoryLastRefreshTime = time.Now()
	}
}

func fetchHistory(db *gorm.DB) []model.History {
	hs := make([]model.History, 0)

	if len(queries) > 0 {
		q := "%" + queries[0] + "%"
		db = db.Where("title like ?", q).Or("url like ?", q)
	}

	if err := db.Order("last_visit_time DESC, visit_count DESC, typed_count DESC").
		Limit(int(config.App.Limit)).Find(&hs).Error; err != nil {
		log.Println("", db.Error.Error())
	}

	return hs
}
