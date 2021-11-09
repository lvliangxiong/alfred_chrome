package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/lvliangxiong/alfred_chrome/model"
)

var (
	fvWG     sync.WaitGroup
	mappings map[string]int64
)

func faviconInit() {
	needRefresh := false
	if time.Since(status.FaviconLastRefreshTime) >= config.Favicon.RefreshInterval {
		needRefresh = true
	}

	if !needRefresh {
		_, err := os.Stat("Favicons")
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

		if err := copyFile(config.getProfileFaviconFilePath(), "Favicons"); err != nil {
			log.Fatal(err)
		}
	}

	db, err := gorm.Open(sqlite.Open("Favicons"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if needRefresh {
		fvWG.Add(1)
		go syncFavicons(db)
	}

	fvWG.Add(1)
	go getUrl2IconMapping(db)

	fvWG.Wait()

	if needRefresh {
		status.FaviconLastRefreshTime = time.Now()
	}
}

func syncFavicons(db *gorm.DB) {
	defer fvWG.Done()

	fvs := make([]model.FaviconBitmap, 0)
	if err := db.Where("height == 32 AND width == 32").Find(&fvs).Error; err != nil {
		logger.Default.Warn(context.Background(), db.Error.Error())
	}

	if err := os.RemoveAll("icons"); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("icons", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	for _, favicon := range fvs {
		savePNG(favicon.ImageData, fmt.Sprintf("icons/%010d.png", favicon.IconID))
	}
}

func savePNG(imageData []byte, filename string) {
	out, err := os.Create(filename) // ignore_security_alert
	if err != nil {
		log.Println(err)
	}

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Println(err)
		}
	}(out)

	_, err = out.Write(imageData)
	if err != nil {
		log.Println(err)
	}
}

func getUrl2IconMapping(db *gorm.DB) {
	defer fvWG.Done()

	iconMappings := make([]model.IconMapping, 0)
	if err := db.Find(&iconMappings).Error; err != nil {
		logger.Default.Warn(context.Background(), db.Error.Error())
	}

	mappings = make(map[string]int64, len(iconMappings))
	for _, mapping := range iconMappings {
		if _, ok := mappings[mapping.PageUrl]; ok {
			continue
		}
		mappings[mapping.PageUrl] = mapping.IconID
	}
}
