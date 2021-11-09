package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

var status RefreshStatus

type RefreshStatus struct {
	HistoryLastRefreshTime time.Time `json:"history_last_refresh_time"`
	FaviconLastRefreshTime time.Time `json:"favicon_last_refresh_time"`
}

func (r *RefreshStatus) load() error {
	bytes, err := os.ReadFile("refresh.json")
	if err != nil && os.IsNotExist(err) {
		err = r.save()
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, r)
	if err != nil {
		return err
	}

	return nil
}

func (r *RefreshStatus) save() error {
	bytes, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.OpenFile("refresh.json", os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func statusInit() {
	err := status.load()
	if err != nil {
		log.Printf("refresh status load failed, err: %v\n", err)
	}
}
