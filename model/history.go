package model

type History struct {
	Id            int64  `column:"id"`
	Url           string `column:"url"`
	Title         string `column:"title"`
	VisitCount    int64  `column:"visit_count"`
	TypedCount    int64  `column:"typed_count"`
	LastVisitTime int64  `column:"last_visit_time"`
	Hidden        int64  `column:"hidden"`
}

func (History) TableName() string {
	return "urls"
}
