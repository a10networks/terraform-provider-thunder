

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategory struct {
	Inst struct {

    BypassedUrls WebCategoryBypassedUrls3661 `json:"bypassed-urls"`
    CategoryListList []WebCategoryCategoryListList `json:"category-list-list"`
    CloudQueryCacheSize int `json:"cloud-query-cache-size"`
    CloudQueryDisable int `json:"cloud-query-disable"`
    DatabaseServer string `json:"database-server"`
    DbUpdateTime string `json:"db-update-time"`
    Enable int `json:"enable"`
    InterceptedUrls WebCategoryInterceptedUrls3662 `json:"intercepted-urls"`
    License WebCategoryLicense3663 `json:"license"`
    OnlineCheckDisable int `json:"online-check-disable"`
    Port int `json:"port" dval:"80"`
    ProxyServer WebCategoryProxyServer3664 `json:"proxy-server"`
    RemoteSyslogEnable int `json:"remote-syslog-enable"`
    ReputationScopeList []WebCategoryReputationScopeList `json:"reputation-scope-list"`
    RtuCacheSize int `json:"rtu-cache-size"`
    RtuUpdateDisable int `json:"rtu-update-disable"`
    RtuUpdateInterval int `json:"rtu-update-interval" dval:"60"`
    Server string `json:"server"`
    ServerTimeout int `json:"server-timeout" dval:"15"`
    SslPort int `json:"ssl-port" dval:"443"`
    Statistics WebCategoryStatistics3665 `json:"statistics"`
    Url WebCategoryUrl3667 `json:"url"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`
    WebReputation WebCategoryWebReputation3668 `json:"web-reputation"`

	} `json:"web-category"`
}


type WebCategoryBypassedUrls3661 struct {
    Uuid string `json:"uuid"`
}


type WebCategoryCategoryListList struct {
    Name string `json:"name"`
    Uncategorized int `json:"uncategorized"`
    RealEstate int `json:"real-estate"`
    ComputerAndInternetSecurity int `json:"computer-and-internet-security"`
    FinancialServices int `json:"financial-services"`
    BusinessAndEconomy int `json:"business-and-economy"`
    ComputerAndInternetInfo int `json:"computer-and-internet-info"`
    Auctions int `json:"auctions"`
    Shopping int `json:"shopping"`
    CultAndOccult int `json:"cult-and-occult"`
    Travel int `json:"travel"`
    Drugs int `json:"drugs"`
    AdultAndPornography int `json:"adult-and-pornography"`
    HomeAndGarden int `json:"home-and-garden"`
    Military int `json:"military"`
    SocialNetwork int `json:"social-network"`
    DeadSites int `json:"dead-sites"`
    StockAdviceAndTools int `json:"stock-advice-and-tools"`
    TrainingAndTools int `json:"training-and-tools"`
    Dating int `json:"dating"`
    SexEducation int `json:"sex-education"`
    Religion int `json:"religion"`
    EntertainmentAndArts int `json:"entertainment-and-arts"`
    PersonalSitesAndBlogs int `json:"personal-sites-and-blogs"`
    Legal int `json:"legal"`
    LocalInformation int `json:"local-information"`
    StreamingMedia int `json:"streaming-media"`
    JobSearch int `json:"job-search"`
    Gambling int `json:"gambling"`
    Translation int `json:"translation"`
    ReferenceAndResearch int `json:"reference-and-research"`
    SharewareAndFreeware int `json:"shareware-and-freeware"`
    PeerToPeer int `json:"peer-to-peer"`
    Marijuana int `json:"marijuana"`
    Hacking int `json:"hacking"`
    Games int `json:"games"`
    PhilosophyAndPolitics int `json:"philosophy-and-politics"`
    Weapons int `json:"weapons"`
    PayToSurf int `json:"pay-to-surf"`
    HuntingAndFishing int `json:"hunting-and-fishing"`
    Society int `json:"society"`
    EducationalInstitutions int `json:"educational-institutions"`
    OnlineGreetingCards int `json:"online-greeting-cards"`
    Sports int `json:"sports"`
    SwimsuitsAndIntimateApparel int `json:"swimsuits-and-intimate-apparel"`
    Questionable int `json:"questionable"`
    Kids int `json:"kids"`
    HateAndRacism int `json:"hate-and-racism"`
    PersonalStorage int `json:"personal-storage"`
    Violence int `json:"violence"`
    KeyloggersAndMonitoring int `json:"keyloggers-and-monitoring"`
    SearchEngines int `json:"search-engines"`
    InternetPortals int `json:"internet-portals"`
    WebAdvertisements int `json:"web-advertisements"`
    Cheating int `json:"cheating"`
    Gross int `json:"gross"`
    WebBasedEmail int `json:"web-based-email"`
    MalwareSites int `json:"malware-sites"`
    PhishingAndOtherFraud int `json:"phishing-and-other-fraud"`
    ProxyAvoidAndAnonymizers int `json:"proxy-avoid-and-anonymizers"`
    SpywareAndAdware int `json:"spyware-and-adware"`
    Music int `json:"music"`
    Government int `json:"government"`
    Nudity int `json:"nudity"`
    NewsAndMedia int `json:"news-and-media"`
    Illegal int `json:"illegal"`
    Cdns int `json:"cdns"`
    InternetCommunications int `json:"internet-communications"`
    BotNets int `json:"bot-nets"`
    Abortion int `json:"abortion"`
    HealthAndMedicine int `json:"health-and-medicine"`
    SpamUrls int `json:"spam-urls"`
    DynamicallyGeneratedContent int `json:"dynamically-generated-content"`
    ParkedDomains int `json:"parked-domains"`
    AlcoholAndTobacco int `json:"alcohol-and-tobacco"`
    ImageAndVideoSearch int `json:"image-and-video-search"`
    FashionAndBeauty int `json:"fashion-and-beauty"`
    RecreationAndHobbies int `json:"recreation-and-hobbies"`
    MotorVehicles int `json:"motor-vehicles"`
    WebHostingSites int `json:"web-hosting-sites"`
    NudityArtistic int `json:"nudity-artistic"`
    IllegalPornography int `json:"illegal-pornography"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []WebCategoryCategoryListListSamplingEnable `json:"sampling-enable"`
}


type WebCategoryCategoryListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type WebCategoryInterceptedUrls3662 struct {
    Uuid string `json:"uuid"`
}


type WebCategoryLicense3663 struct {
    Uuid string `json:"uuid"`
}


type WebCategoryProxyServer3664 struct {
    ProxyHost string `json:"proxy-host"`
    HttpPort int `json:"http-port"`
    HttpsPort int `json:"https-port"`
    AuthType string `json:"auth-type" dval:"ntlm"`
    Domain string `json:"domain"`
    Username string `json:"username"`
    Password int `json:"password"`
    SecretString string `json:"secret-string"`
    Encrypted string `json:"encrypted"`
    Uuid string `json:"uuid"`
}


type WebCategoryReputationScopeList struct {
    Name string `json:"name"`
    GreaterThan WebCategoryReputationScopeListGreaterThan `json:"greater-than"`
    LessThan WebCategoryReputationScopeListLessThan `json:"less-than"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []WebCategoryReputationScopeListSamplingEnable `json:"sampling-enable"`
}


type WebCategoryReputationScopeListGreaterThan struct {
    GreaterTrustworthy int `json:"greater-trustworthy"`
    GreaterLowRisk int `json:"greater-low-risk"`
    GreaterModerateRisk int `json:"greater-moderate-risk"`
    GreaterSuspicious int `json:"greater-suspicious"`
    GreaterMalicious int `json:"greater-malicious"`
    GreaterThreshold int `json:"greater-threshold"`
}


type WebCategoryReputationScopeListLessThan struct {
    LessTrustworthy int `json:"less-trustworthy"`
    LessLowRisk int `json:"less-low-risk"`
    LessModerateRisk int `json:"less-moderate-risk"`
    LessSuspicious int `json:"less-suspicious"`
    LessMalicious int `json:"less-malicious"`
    LessThreshold int `json:"less-threshold"`
}


type WebCategoryReputationScopeListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type WebCategoryStatistics3665 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []WebCategoryStatisticsSamplingEnable3666 `json:"sampling-enable"`
}


type WebCategoryStatisticsSamplingEnable3666 struct {
    Counters1 string `json:"counters1"`
}


type WebCategoryUrl3667 struct {
    Uuid string `json:"uuid"`
}


type WebCategoryWebReputation3668 struct {
    Uuid string `json:"uuid"`
    InterceptedUrls WebCategoryWebReputationInterceptedUrls3669 `json:"intercepted-urls"`
    BypassedUrls WebCategoryWebReputationBypassedUrls3670 `json:"bypassed-urls"`
    Url WebCategoryWebReputationUrl3671 `json:"url"`
}


type WebCategoryWebReputationInterceptedUrls3669 struct {
    Uuid string `json:"uuid"`
}


type WebCategoryWebReputationBypassedUrls3670 struct {
    Uuid string `json:"uuid"`
}


type WebCategoryWebReputationUrl3671 struct {
    Uuid string `json:"uuid"`
}

func (p *WebCategory) GetId() string{
    return "1"
}

func (p *WebCategory) getPath() string{
    return "web-category"
}

func (p *WebCategory) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategory::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *WebCategory) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategory::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *WebCategory) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategory::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *WebCategory) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategory::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
