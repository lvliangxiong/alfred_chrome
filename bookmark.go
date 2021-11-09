package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	aw "github.com/deanishe/awgo"

	"github.com/lvliangxiong/alfred_chrome/model"
)

var bookmarks []model.Bookmark

func bookmarkInit() {
	bookmarksFile := readBookmarks()
	bookmarks = parse(bookmarksFile)
}

func showAllBookmarkItems() {
	for _, bm := range bookmarks {
		item := wf.NewItem(bm.Name).
			Subtitle(bm.Url).
			Arg(bm.Url).
			UID(bm.Guid).
			Valid(true).
			Match(bm.Name + " " + bm.Url) // used for awgo's fuzzy sort

		if iconId, ok := mappings[bm.Url]; ok {
			item.Icon(&aw.Icon{Value: fmt.Sprintf("icons/%010d.png", iconId)})
		} else {
			wf.NewItem(bm.Name).Icon(iconChrome)
		}
	}
}

func readBookmarks() *model.BookmarksFile {
	bytes, err := os.ReadFile(config.getProfileBookmarkFilePath())
	if err != nil {
		log.Fatalf("parsing bookmark error, err is %s\n", err.Error())
	}
	bookmarksFile := &model.BookmarksFile{}
	err = json.Unmarshal(bytes, bookmarksFile)
	if err != nil {
		log.Fatalf("parsing bookmark error, err is %s\n", err.Error())
	}
	return bookmarksFile
}

func parse(bookmarksFile *model.BookmarksFile) []model.Bookmark {
	bs := make([]model.Bookmark, 0)
	bookmarks := []model.Bookmark{bookmarksFile.Roots.BookmarkBar, bookmarksFile.Roots.Other, bookmarksFile.Roots.Synced}
	for _, b := range bookmarks {
		bs = append(bs, retrieveBookmarks(b)...)
	}
	return bs
}

func retrieveBookmarks(bookmark model.Bookmark) []model.Bookmark {
	if len(bookmark.Children) == 0 && bookmark.Type == "url" {
		return []model.Bookmark{bookmark}
	}

	bs := make([]model.Bookmark, 0)
	if len(bookmark.Children) != 0 {
		for _, child := range bookmark.Children {
			childBookmarks := retrieveBookmarks(child)
			bs = append(bs, childBookmarks...)
		}
	}
	return bs
}
