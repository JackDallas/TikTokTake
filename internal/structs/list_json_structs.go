package structs

type ListJSON struct {
	StatusCode int         `json:"statusCode"`
	Body       Body        `json:"body"`
	ErrMsg     interface{} `json:"errMsg"`
}

type VideoMeta struct {
	Width    int `json:"width"`
	Height   int `json:"height"`
	Ratio    int `json:"ratio"`
	Duration int `json:"duration"`
}
type Video struct {
	Urls      []string  `json:"urls"`
	VideoMeta VideoMeta `json:"videoMeta"`
}
type ItemInfos struct {
	ID             string        `json:"id"`
	Text           string        `json:"text"`
	CreateTime     string        `json:"createTime"`
	AuthorID       string        `json:"authorId"`
	MusicID        string        `json:"musicId"`
	Covers         []string      `json:"covers"`
	CoversOrigin   []string      `json:"coversOrigin"`
	CoversDynamic  []string      `json:"coversDynamic"`
	Video          Video         `json:"video"`
	DiggCount      int           `json:"diggCount"`
	ShareCount     int           `json:"shareCount"`
	PlayCount      int           `json:"playCount"`
	CommentCount   int           `json:"commentCount"`
	IsOriginal     bool          `json:"isOriginal"`
	IsOfficial     bool          `json:"isOfficial"`
	IsActivityItem bool          `json:"isActivityItem"`
	WarnInfo       []interface{} `json:"warnInfo"`
}
type AuthorInfos struct {
	SecUID       string   `json:"secUid"`
	UserID       string   `json:"userId"`
	UniqueID     string   `json:"uniqueId"`
	NickName     string   `json:"nickName"`
	Signature    string   `json:"signature"`
	Verified     bool     `json:"verified"`
	Covers       []string `json:"covers"`
	CoversMedium []string `json:"coversMedium"`
	CoversLarger []string `json:"coversLarger"`
	IsSecret     bool     `json:"isSecret"`
}
type MusicInfos struct {
	MusicID      string   `json:"musicId"`
	MusicName    string   `json:"musicName"`
	AuthorName   string   `json:"authorName"`
	Original     bool     `json:"original"`
	PlayURL      []string `json:"playUrl"`
	Covers       []string `json:"covers"`
	CoversMedium []string `json:"coversMedium"`
	CoversLarger []string `json:"coversLarger"`
}
type ItemListData struct {
	ItemInfos         ItemInfos     `json:"itemInfos"`
	AuthorInfos       AuthorInfos   `json:"authorInfos"`
	MusicInfos        MusicInfos    `json:"musicInfos"`
	ChallengeInfoList []interface{} `json:"challengeInfoList"`
	DuetInfo          string        `json:"duetInfo"`
}
type Body struct {
	PageState    PageState      `json:"pageState"`
	ItemListData []ItemListData `json:"itemListData"`
	ItemABParams []interface{}  `json:"itemABParams"`
	HasMore      bool           `json:"hasMore"`
	MaxCursor    string         `json:"maxCursor"`
	MinCursor    string         `json:"minCursor"`
}
