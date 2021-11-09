package main

import (
	"fmt"
	"log"
	"os"
	"time"

	aw "github.com/deanishe/awgo"
	"gopkg.in/yaml.v3"
)

var (
	config Config
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Bookmark BookmarkConfig `yaml:"bookmark"`
	History  HistoryConfig  `yaml:"history"`
	Favicon  FaviconConfig  `yaml:"favicon"`
}

type AppConfig struct {
	Profile string `yaml:"profile"`
	Limit   int32  `yaml:"limit"`
}

type BookmarkConfig struct {
	RefreshInterval time.Duration `yaml:"refresh_interval"`
}

type HistoryConfig struct {
	RefreshInterval time.Duration `yaml:"refresh_interval"`
}

type FaviconConfig struct {
	RefreshInterval time.Duration `yaml:"refresh_interval"`
}

func (c Config) getProfileBookmarkFilePath() string {
	return fmt.Sprintf("%s/Library/Application Support/Google/Chrome/%s/Bookmarks", homeDir, config.App.Profile)
}

func (c Config) getProfileFaviconFilePath() string {
	return fmt.Sprintf("%s/Library/Application Support/Google/Chrome/%s/Favicons", homeDir, config.App.Profile)
}

func (c Config) getProfileHistoryFilePath() string {
	return fmt.Sprintf("%s/Library/Application Support/Google/Chrome/%s/History", homeDir, config.App.Profile)
}

type WorkflowOptions struct {
	Profile                 string `env:"CHROME_PROFILE"`
	Limit                   int32  `env:"LIMIT"`
	BookmarkRefreshInterval string `env:"BOOKMARK_REFRESH_INTERVAL"`
	HistoryRefreshInterval  string `env:"HISTORY_REFRESH_INTERVAL"`
	FaviconRefreshInterval  string `env:"FAVICON_REFRESH_INTERVAL"`
}

func confInit() {
	f, err := os.Open("conf.yaml")
	if err != nil {
		log.Fatalln("read config file failed")
	}

	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		log.Fatalf("decode config file failed: %v\n", err)
	}

	opts := WorkflowOptions{}
	cfg := aw.NewConfig()
	if err := cfg.To(&opts); err != nil {
		panic(err)
	}

	if opts.Limit > 0 {
		config.App.Limit = opts.Limit
	}
	if opts.Profile != "" {
		config.App.Profile = opts.Profile
	}

	if opts.BookmarkRefreshInterval != "" {
		interval, err := time.ParseDuration(opts.BookmarkRefreshInterval)
		if err == nil {
			config.Bookmark.RefreshInterval = interval
		}
	}
	if opts.HistoryRefreshInterval != "" {
		interval, err := time.ParseDuration(opts.HistoryRefreshInterval)
		if err == nil {
			config.History.RefreshInterval = interval
		}
	}
	if opts.FaviconRefreshInterval != "" {
		interval, err := time.ParseDuration(opts.FaviconRefreshInterval)
		if err == nil {
			config.Favicon.RefreshInterval = interval
		}
	}

	log.Printf("%#v", config)
}
