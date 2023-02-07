package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type WebCategoryCategoryList struct {
	Name WebCategoryCategoryListInstance `json:"category-list,omitempty"`
}

type WebCategoryCategoryListInstance struct {
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
	Counters1                   []WebCategorySamplingEnable `json:"sampling-enable,omitempty"`
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

type WebCategorySamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostWebCategoryCategoryList(id string, inst WebCategoryCategoryList, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostWebCategoryCategoryList")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/web-category/category-list", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryCategoryList
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostWebCategoryCategoryList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetWebCategoryCategoryList(id string, name1 string, host string) (*WebCategoryCategoryList, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetWebCategoryCategoryList")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/web-category/category-list/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryCategoryList
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetWebCategoryCategoryList", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutWebCategoryCategoryList(id string, name1 string, inst WebCategoryCategoryList, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutWebCategoryCategoryList")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/web-category/category-list/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryCategoryList
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutWebCategoryCategoryList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteWebCategoryCategoryList(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteWebCategoryCategoryList")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/web-category/category-list/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryCategoryList
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteWebCategoryCategoryList", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
