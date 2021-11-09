package main

import (
	aw "github.com/deanishe/awgo"
)

func run() {
	if doCheck {
		CheckAndCache()
		return
	}

	autoCheck()
	showUpdateIfAvailable()

	if showBookmark {
		showAllBookmarkItems()
	}

	if showHistory {
		showAllHistoryItems()
	}

	// Add an extra item to reset update status for demo purposes.
	// As with the update notification, this item triggers a Magic
	// Action that deletes the cached list of releases.
	wf.NewItem("Reset update status").
		Autocomplete("workflow:delcache").
		Icon(aw.IconTrash).
		Valid(false)

	wf.NewItem("Help").
		Autocomplete("workflow:help").
		Icon(aw.IconHelp).
		Valid(false)

	// Filter results on user query if present
	for _, query := range queries {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
}
