package structs

type TikTokMeta struct {
	Props       Props  `json:"props"`
	Page        string `json:"page"`
	Query       Query  `json:"query"`
	BuildID     string `json:"buildId"`
	AssetPrefix string `json:"assetPrefix"`
}
type Children struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
type LanguageList struct {
	Value    string     `json:"value"`
	Alias    string     `json:"alias"`
	Label    string     `json:"label"`
	Children []Children `json:"children"`
}
type Amazon struct {
	Visible bool   `json:"visible"`
	Normal  string `json:"normal"`
}
type Google struct {
	Visible bool   `json:"visible"`
	Normal  string `json:"normal"`
}
type Apple struct {
	Visible bool   `json:"visible"`
	Normal  string `json:"normal"`
}
type DownloadLink struct {
	Amazon Amazon `json:"amazon"`
	Google Google `json:"google"`
	Apple  Apple  `json:"apple"`
}
type AbTestVersion struct {
	ClientParameters  string `json:"clientParameters"`
	ClientVersionName string `json:"clientVersionName"`
	VersionName       string `json:"versionName"`
	Parameters        string `json:"parameters"`
	StartTime         string `json:"startTime"`
	EndTime           string `json:"endTime"`
}
type LegalList struct {
	Title string `json:"title"`
	Key   string `json:"key"`
	Href  string `json:"href"`
}
type InitialProps struct {
	IsMobile         bool           `json:"$isMobile"`
	IsIOS            interface{}    `json:"$isIOS"`
	IsAndroid        bool           `json:"$isAndroid"`
	Host             string         `json:"$host"`
	PageURL          string         `json:"$pageUrl"`
	Language         string         `json:"$language"`
	OriginalLanguage string         `json:"$originalLanguage"`
	LanguageList     []LanguageList `json:"$languageList"`
	Region           string         `json:"$region"`
	AppID            int            `json:"$appId"`
	Os               string         `json:"$os"`
	BaseURL          string         `json:"$baseURL"`
	DownloadLink     DownloadLink   `json:"$downloadLink"`
	AbTestVersion    AbTestVersion  `json:"$abTestVersion"`
	AppType          string         `json:"$appType"`
	Gray             string         `json:"$gray"`
	ReflowType       string         `json:"$reflowType"`
	LegalList        []LegalList    `json:"$legalList"`
}
type PageState struct {
	RegionAppID int    `json:"regionAppId"`
	Os          string `json:"os"`
	Region      string `json:"region"`
	BaseURL     string `json:"baseURL"`
	AppType     string `json:"appType"`
	FullURL     string `json:"fullUrl"`
}
type UserData struct {
	SecUID       string   `json:"secUid"`
	UserID       string   `json:"userId"`
	IsSecret     bool     `json:"isSecret"`
	UniqueID     string   `json:"uniqueId"`
	NickName     string   `json:"nickName"`
	Signature    string   `json:"signature"`
	Covers       []string `json:"covers"`
	CoversMedium []string `json:"coversMedium"`
	Following    int      `json:"following"`
	Fans         int      `json:"fans"`
	Heart        int      `json:"heart"`
	Video        int      `json:"video"`
	Verified     bool     `json:"verified"`
	Digg         int      `json:"digg"`
}
type ShareUser struct {
	SecUID       string        `json:"secUid"`
	UserID       string        `json:"userId"`
	UniqueID     string        `json:"uniqueId"`
	NickName     string        `json:"nickName"`
	Signature    string        `json:"signature"`
	Covers       []interface{} `json:"covers"`
	CoversMedium []interface{} `json:"coversMedium"`
	CoversLarger []interface{} `json:"coversLarger"`
	IsSecret     bool          `json:"isSecret"`
}
type ShareMeta struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type LangList struct {
	Value    string     `json:"value"`
	Alias    string     `json:"alias"`
	Label    string     `json:"label"`
	Children []Children `json:"children"`
}
type FeedConfig struct {
	PageType   int    `json:"pageType"`
	SecUID     string `json:"secUid"`
	ID         string `json:"id"`
	ShowAvatar bool   `json:"showAvatar"`
	EmptyTip   string `json:"emptyTip"`
}
type TeaApp struct {
	Open bool `json:"open"`
}
type Tea struct {
	Open bool `json:"open"`
}
type Config struct {
	TeaApp TeaApp `json:"teaApp"`
	Tea    Tea    `json:"tea"`
}
type Footer struct {
	ShowDownload bool `json:"showDownload"`
}
type Header struct {
	ShowUpload bool   `json:"showUpload"`
	Type       string `json:"type"`
}
type PageOptions struct {
	Footer      Footer      `json:"footer"`
	Header      Header      `json:"header"`
	HeadOptions interface{} `json:"headOptions"`
}
type PageProps struct {
	UniqueID      string        `json:"uniqueId"`
	PageState     PageState     `json:"pageState"`
	UserData      UserData      `json:"userData"`
	ShareUser     ShareUser     `json:"shareUser"`
	ShareMeta     ShareMeta     `json:"shareMeta"`
	StatusCode    int           `json:"statusCode"`
	ItemList      []interface{} `json:"itemList"`
	LangList      []LangList    `json:"langList"`
	TestID        string        `json:"testId"`
	IsInSeoulTest bool          `json:"isInSeoulTest"`
	FeedConfig    FeedConfig    `json:"feedConfig"`
	Config        Config        `json:"config"`
	PageOptions   PageOptions   `json:"pageOptions"`
}
type Props struct {
	InitialProps InitialProps `json:"initialProps"`
	PageProps    PageProps    `json:"pageProps"`
	Pathname     string       `json:"pathname"`
}
type ZeroFea6A13C52B4D4725368F24B045Ca84 struct {
}
type Aa59D67C2123F094D0D6798Ffe651C4D struct {
}
type Query struct {
	UniqueID                            string                              `json:"uniqueId"`
	Webtoken                            string                              `json:"webtoken"`
	ZeroFea6A13C52B4D4725368F24B045Ca84 ZeroFea6A13C52B4D4725368F24B045Ca84 `json:"0fea6a13c52b4d4725368f24b045ca84"`
	Aa59D67C2123F094D0D6798Ffe651C4D    Aa59D67C2123F094D0D6798Ffe651C4D    `json:"aa59d67c2123f094d0d6798ffe651c4d"`
}
