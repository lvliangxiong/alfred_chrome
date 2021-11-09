package main

import (
	"log"
	"os"
	"os/exec"

	aw "github.com/deanishe/awgo"
)

// CheckAndCache 每间隔一段时间就会进行一次更新检查并缓存最新的结果，默认是 24h
func CheckAndCache() {
	wf.Configure(aw.TextErrors(true))
	log.Println("Checking for updates...")
	log.Println(wf.CacheDir())
	log.Println(wf.DataDir())
	if err := wf.CheckForUpdate(); err != nil {
		wf.FatalError(err)
	}
}

func autoCheck() {
	// 一定间隔后，会自动后台起一个进程执行自身，带上 -check 参数，触发 doCheck 部分的逻辑
	if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
		log.Println("Running update check in background...")

		cmd := exec.Command(os.Args[0], "-check")
		if err := wf.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}
}

func showUpdateIfAvailable() {
	// Only show update status if query is empty.
	if len(queries) == 0 && wf.UpdateAvailable() {
		// Turn off UIDs to force this item to the top.
		// If UIDs are enabled, Alfred will apply its "knowledge"
		// to order the results based on your past usage.
		wf.Configure(aw.SuppressUIDs(true))

		// Notify user of update. As this item is invalid (Valid(false)),
		// actioning it expands the query to the Autocomplete value.
		// "workflow:update" triggers the updater Magic Action that
		// is automatically registered when you configure Workflow with
		// an Updater.
		//
		// If executed, the Magic Action downloads the latest version
		// of the workflow and asks Alfred to install it.
		wf.NewItem("Update available!").
			Subtitle("↩ to install").
			Autocomplete("workflow:update").
			Valid(false).
			Icon(iconAvailable)
	}
}
