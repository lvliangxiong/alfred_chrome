package model

type FaviconBitmap struct {
	ID            int64  `column:"id"`
	IconID        int64  `column:"icon_id"`
	LastUpdated   int64  `column:"last_updated"`
	ImageData     []byte `column:"image_data"`
	Width         int32  `column:"width"`
	Height        int32  `column:"height"`
	LastRequested int64  `column:"last_requested"`
}

func (FaviconBitmap) TableName() string {
	return "favicon_bitmaps"
}

type IconMapping struct {
	ID      int64  `column:"id"`
	PageUrl string `column:"page_url"`
	IconID  int64  `column:"icon_id"`
}

func (IconMapping) TableName() string {
	return "icon_mapping"
}
