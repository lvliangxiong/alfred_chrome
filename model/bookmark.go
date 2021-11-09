package model

type BookmarksFile struct {
	CheckSum string `json:"checksum"`
	Roots    struct {
		BookmarkBar Bookmark `json:"bookmark_bar"`
		Other       Bookmark `json:"other"`
		Synced      Bookmark `json:"synced"`
	} `json:"roots"`
	SyncMetadata string `json:"sync_metadata"`
	Version      int    `json:"version"`
}

type Bookmark struct {
	Children     []Bookmark `json:"children"`
	DateAdded    string     `json:"date_added"`
	DateModified string     `json:"date_modified"`
	Guid         string     `json:"guid"`
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	Url          string     `json:"url"`
}
