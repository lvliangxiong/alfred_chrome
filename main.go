package main

import (
	"flag"
	"os"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	doCheck      bool
	showBookmark bool
	showHistory  bool
	enableFilter bool

	queries []string // command-line argument

	// Icon to show if an update is available
	iconAvailable = &aw.Icon{Value: "update-available.png"}
	iconChrome    = &aw.Icon{Value: "icon.png"}

	homeDir string

	// For auto update
	repo    = "lvliangxiong/alfred_chrome"
	helpUrl = "https://www.google.com"

	// Workflow stuff
	wf *aw.Workflow
)

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

func flagInit() {
	flag.BoolVar(&doCheck, "check", false, "check for a new version")
	flag.BoolVar(&showBookmark, "bm", false, "search in the bookmarks")
	flag.BoolVar(&showHistory, "hi", false, "search in the browser history")
	flag.BoolVar(&enableFilter, "filter", false, "enable filtering in history searching")
	flag.Parse()
}

func wfInit() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(helpUrl)) // alfred workflow stuff

	wf.Args()
	queries = flag.Args()

	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
}

func main() {
	flagInit() // init flags
	wfInit()   // init wf and query

	statusInit()   // init last refresh status
	confInit()     // load config
	faviconInit()  // load favicon mapping and refresh if necessary
	bookmarkInit() // load bookmark, always the newest
	historyInit()  // load histories and refresh if necessary

	wf.Run(run)
}
