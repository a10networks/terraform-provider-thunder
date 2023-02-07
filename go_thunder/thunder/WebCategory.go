package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type WebCategorySelf struct {
	Server WebCategorySelfInstance `json:"web-category,omitempty"`
}

type WebCategorySelfInstance struct {
	//UUID                BypassedUrls          `json:"bypassed-urls,omitempty"`
	Name                []CategoryListList `json:"category-list-list,omitempty"`
	CloudQueryCacheSize int                `json:"cloud-query-cache-size,omitempty"`
	CloudQueryDisable   int                `json:"cloud-query-disable,omitempty"`
	DatabaseServer      string             `json:"database-server,omitempty"`
	DbUpdateTime        string             `json:"db-update-time,omitempty"`
	Enable              int                `json:"enable,omitempty"`
	//UUID                BypassedUrls          `json:"intercepted-urls,omitempty"`
	//UUID                BypassedUrls          `json:"license,omitempty"`
	Port               int                    `json:"port,omitempty"`
	ProxyHost          ProxyServerWebCategory `json:"proxy-server,omitempty"`
	RemoteSyslogEnable int                    `json:"remote-syslog-enable,omitempty"`
	GreaterTrustworthy []ReputationScopeList  `json:"reputation-scope-list,omitempty"`
	RtuCacheSize       int                    `json:"rtu-cache-size,omitempty"`
	RtuUpdateDisable   int                    `json:"rtu-update-disable,omitempty"`
	RtuUpdateInterval  int                    `json:"rtu-update-interval,omitempty"`
	Server             string                 `json:"server,omitempty"`
	ServerTimeout      int                    `json:"server-timeout,omitempty"`
	SslPort            int                    `json:"ssl-port,omitempty"`
	Counters1          Statistics             `json:"statistics,omitempty"`
	//UUID                BypassedUrls          `json:"url,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	UseMgmtPort int    `json:"use-mgmt-port,omitempty"`
	//UUID        WebReputation `json:"web-reputation,omitempty"`
}

type BypassedUrls struct {
	UUID string `json:"uuid,omitempty"`
}

type CategoryListList struct {
	Abortion                    int                         `json:"abortion,omitempty"`
	AdultAndPornography         int                         `json:"adult-and-pornography,omitempty"`
	AlcoholAndTobacco           int                         `json:"alcohol-and-tobacco,omitempty"`
	Auctions                    int                         `json:"auctions,omitempty"`
	BotNets                     int                         `json:"bot-nets,omitempty"`
	BusinessAndEconomy          int                         `json:"business-and-economy,omitempty"`
	Cdns                        int                         `json:"cdns,omitempty"`
	Cheating                    int                         `json:"cheating,omitempty"`
	ComputerAndInternetInfo     int                         `json:"computer-and-internet-info,omitempty"`
	ComputerAndInternetSecurity int                         `json:"computer-and-internet-security,omitempty"`
	ConfirmedSpamSources        int                         `json:"confirmed-spam-sources,omitempty"`
	CultAndOccult               int                         `json:"cult-and-occult,omitempty"`
	Dating                      int                         `json:"dating,omitempty"`
	DeadSites                   int                         `json:"dead-sites,omitempty"`
	Drugs                       int                         `json:"drugs,omitempty"`
	DynamicComment              int                         `json:"dynamic-comment,omitempty"`
	EducationalInstitutions     int                         `json:"educational-institutions,omitempty"`
	EntertainmentAndArts        int                         `json:"entertainment-and-arts,omitempty"`
	FashionAndBeauty            int                         `json:"fashion-and-beauty,omitempty"`
	FinancialServices           int                         `json:"financial-services,omitempty"`
	FoodAndDining               int                         `json:"food-and-dining,omitempty"`
	Gambling                    int                         `json:"gambling,omitempty"`
	Games                       int                         `json:"games,omitempty"`
	Government                  int                         `json:"government,omitempty"`
	Gross                       int                         `json:"gross,omitempty"`
	Hacking                     int                         `json:"hacking,omitempty"`
	HateAndRacism               int                         `json:"hate-and-racism,omitempty"`
	HealthAndMedicine           int                         `json:"health-and-medicine,omitempty"`
	HomeAndGarden               int                         `json:"home-and-garden,omitempty"`
	HuntingAndFishing           int                         `json:"hunting-and-fishing,omitempty"`
	Illegal                     int                         `json:"illegal,omitempty"`
	ImageAndVideoSearch         int                         `json:"image-and-video-search,omitempty"`
	InternetCommunications      int                         `json:"internet-communications,omitempty"`
	InternetPortals             int                         `json:"internet-portals,omitempty"`
	JobSearch                   int                         `json:"job-search,omitempty"`
	KeyloggersAndMonitoring     int                         `json:"keyloggers-and-monitoring,omitempty"`
	Kids                        int                         `json:"kids,omitempty"`
	Legal                       int                         `json:"legal,omitempty"`
	LocalInformation            int                         `json:"local-information,omitempty"`
	MalwareSites                int                         `json:"malware-sites,omitempty"`
	Marijuana                   int                         `json:"marijuana,omitempty"`
	Military                    int                         `json:"military,omitempty"`
	MotorVehicles               int                         `json:"motor-vehicles,omitempty"`
	Music                       int                         `json:"music,omitempty"`
	Name                        string                      `json:"name,omitempty"`
	NewsAndMedia                int                         `json:"news-and-media,omitempty"`
	Nudity                      int                         `json:"nudity,omitempty"`
	OnlineGreetingCards         int                         `json:"online-greeting-cards,omitempty"`
	OpenHTTPProxies             int                         `json:"open-http-proxies,omitempty"`
	ParkedDomains               int                         `json:"parked-domains,omitempty"`
	PayToSurf                   int                         `json:"pay-to-surf,omitempty"`
	PeerToPeer                  int                         `json:"peer-to-peer,omitempty"`
	PersonalSitesAndBlogs       int                         `json:"personal-sites-and-blogs,omitempty"`
	PersonalStorage             int                         `json:"personal-storage,omitempty"`
	PhilosophyAndPolitics       int                         `json:"philosophy-and-politics,omitempty"`
	PhishingAndOtherFraud       int                         `json:"phishing-and-other-fraud,omitempty"`
	PrivateIPAddresses          int                         `json:"private-ip-addresses,omitempty"`
	ProxyAvoidAndAnonymizers    int                         `json:"proxy-avoid-and-anonymizers,omitempty"`
	Questionable                int                         `json:"questionable,omitempty"`
	RealEstate                  int                         `json:"real-estate,omitempty"`
	RecreationAndHobbies        int                         `json:"recreation-and-hobbies,omitempty"`
	ReferenceAndResearch        int                         `json:"reference-and-research,omitempty"`
	Religion                    int                         `json:"religion,omitempty"`
	Counters1                   []SamplingEnableWebCategory `json:"sampling-enable,omitempty"`
	SearchEngines               int                         `json:"search-engines,omitempty"`
	SexEducation                int                         `json:"sex-education,omitempty"`
	SharewareAndFreeware        int                         `json:"shareware-and-freeware,omitempty"`
	Shopping                    int                         `json:"shopping,omitempty"`
	SocialNetwork               int                         `json:"social-network,omitempty"`
	Society                     int                         `json:"society,omitempty"`
	SpamUrls                    int                         `json:"spam-urls,omitempty"`
	Sports                      int                         `json:"sports,omitempty"`
	SpywareAndAdware            int                         `json:"spyware-and-adware,omitempty"`
	StockAdviceAndTools         int                         `json:"stock-advice-and-tools,omitempty"`
	StreamingMedia              int                         `json:"streaming-media,omitempty"`
	SwimsuitsAndIntimateApparel int                         `json:"swimsuits-and-intimate-apparel,omitempty"`
	TrainingAndTools            int                         `json:"training-and-tools,omitempty"`
	Translation                 int                         `json:"translation,omitempty"`
	Travel                      int                         `json:"travel,omitempty"`
	UUID                        string                      `json:"uuid,omitempty"`
	Uncategorized               int                         `json:"uncategorized,omitempty"`
	UnconfirmedSpamSources      int                         `json:"unconfirmed-spam-sources,omitempty"`
	UserTag                     string                      `json:"user-tag,omitempty"`
	Violence                    int                         `json:"violence,omitempty"`
	Weapons                     int                         `json:"weapons,omitempty"`
	WebAdvertisements           int                         `json:"web-advertisements,omitempty"`
	WebBasedEmail               int                         `json:"web-based-email,omitempty"`
	WebHostingSites             int                         `json:"web-hosting-sites,omitempty"`
}

type ProxyServerWebCategory struct {
	AuthType     string `json:"auth-type,omitempty"`
	Domain       string `json:"domain,omitempty"`
	HTTPPort     int    `json:"http-port,omitempty"`
	HTTPSPort    int    `json:"https-port,omitempty"`
	Password     int    `json:"password,omitempty"`
	ProxyHost    string `json:"proxy-host,omitempty"`
	SecretString string `json:"secret-string,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	Username     string `json:"username,omitempty"`
}

type ReputationScopeList struct {
	GreaterTrustworthy GreaterThan                 `json:"greater-than,omitempty"`
	LessTrustworthy    LessThan                    `json:"less-than,omitempty"`
	Name               string                      `json:"name,omitempty"`
	Counters1          []SamplingEnableWebCategory `json:"sampling-enable,omitempty"`
	UUID               string                      `json:"uuid,omitempty"`
	UserTag            string                      `json:"user-tag,omitempty"`
}

type Statistics struct {
	Counters1 []SamplingEnableWebCategory `json:"sampling-enable,omitempty"`
	UUID      string                      `json:"uuid,omitempty"`
}

// type WebReputation struct {
// 	UUID BypassedUrls `json:"bypassed-urls,omitempty"`
// 	UUID BypassedUrls `json:"intercepted-urls,omitempty"`
// 	UUID BypassedUrls `json:"url,omitempty"`
// 	UUID string       `json:"uuid,omitempty"`
// }

type SamplingEnableWebCategory struct {
	Counters1 string `json:"counters1,omitempty"`
}

type GreaterThan struct {
	GreaterLowRisk      int `json:"greater-low-risk,omitempty"`
	GreaterMalicious    int `json:"greater-malicious,omitempty"`
	GreaterModerateRisk int `json:"greater-moderate-risk,omitempty"`
	GreaterSuspicious   int `json:"greater-suspicious,omitempty"`
	GreaterThreshold    int `json:"greater-threshold,omitempty"`
	GreaterTrustworthy  int `json:"greater-trustworthy,omitempty"`
}

type LessThan struct {
	LessLowRisk      int `json:"less-low-risk,omitempty"`
	LessMalicious    int `json:"less-malicious,omitempty"`
	LessModerateRisk int `json:"less-moderate-risk,omitempty"`
	LessSuspicious   int `json:"less-suspicious,omitempty"`
	LessThreshold    int `json:"less-threshold,omitempty"`
	LessTrustworthy  int `json:"less-trustworthy,omitempty"`
}

func PostWebCategory(id string, inst WebCategorySelf, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostWebCategory")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/web-category", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategorySelf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostWebCategory", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetWebCategory(id string, host string) (*WebCategorySelf, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetWebCategory")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/web-category", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategorySelf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetWebCategory", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
