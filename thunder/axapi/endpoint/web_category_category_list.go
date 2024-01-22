

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryCategoryList struct {
	Inst struct {

    Abortion int `json:"abortion"`
    AdultAndPornography int `json:"adult-and-pornography"`
    AlcoholAndTobacco int `json:"alcohol-and-tobacco"`
    Auctions int `json:"auctions"`
    BotNets int `json:"bot-nets"`
    BusinessAndEconomy int `json:"business-and-economy"`
    Cdns int `json:"cdns"`
    Cheating int `json:"cheating"`
    ComputerAndInternetInfo int `json:"computer-and-internet-info"`
    ComputerAndInternetSecurity int `json:"computer-and-internet-security"`
    CultAndOccult int `json:"cult-and-occult"`
    Dating int `json:"dating"`
    DeadSites int `json:"dead-sites"`
    Drugs int `json:"drugs"`
    DynamicallyGeneratedContent int `json:"dynamically-generated-content"`
    EducationalInstitutions int `json:"educational-institutions"`
    EntertainmentAndArts int `json:"entertainment-and-arts"`
    FashionAndBeauty int `json:"fashion-and-beauty"`
    FinancialServices int `json:"financial-services"`
    Gambling int `json:"gambling"`
    Games int `json:"games"`
    Government int `json:"government"`
    Gross int `json:"gross"`
    Hacking int `json:"hacking"`
    HateAndRacism int `json:"hate-and-racism"`
    HealthAndMedicine int `json:"health-and-medicine"`
    HomeAndGarden int `json:"home-and-garden"`
    HuntingAndFishing int `json:"hunting-and-fishing"`
    Illegal int `json:"illegal"`
    IllegalPornography int `json:"illegal-pornography"`
    ImageAndVideoSearch int `json:"image-and-video-search"`
    InternetCommunications int `json:"internet-communications"`
    InternetPortals int `json:"internet-portals"`
    JobSearch int `json:"job-search"`
    KeyloggersAndMonitoring int `json:"keyloggers-and-monitoring"`
    Kids int `json:"kids"`
    Legal int `json:"legal"`
    LocalInformation int `json:"local-information"`
    MalwareSites int `json:"malware-sites"`
    Marijuana int `json:"marijuana"`
    Military int `json:"military"`
    MotorVehicles int `json:"motor-vehicles"`
    Music int `json:"music"`
    Name string `json:"name"`
    NewsAndMedia int `json:"news-and-media"`
    Nudity int `json:"nudity"`
    NudityArtistic int `json:"nudity-artistic"`
    OnlineGreetingCards int `json:"online-greeting-cards"`
    ParkedDomains int `json:"parked-domains"`
    PayToSurf int `json:"pay-to-surf"`
    PeerToPeer int `json:"peer-to-peer"`
    PersonalSitesAndBlogs int `json:"personal-sites-and-blogs"`
    PersonalStorage int `json:"personal-storage"`
    PhilosophyAndPolitics int `json:"philosophy-and-politics"`
    PhishingAndOtherFraud int `json:"phishing-and-other-fraud"`
    ProxyAvoidAndAnonymizers int `json:"proxy-avoid-and-anonymizers"`
    Questionable int `json:"questionable"`
    RealEstate int `json:"real-estate"`
    RecreationAndHobbies int `json:"recreation-and-hobbies"`
    ReferenceAndResearch int `json:"reference-and-research"`
    Religion int `json:"religion"`
    SamplingEnable []WebCategoryCategoryListSamplingEnable `json:"sampling-enable"`
    SearchEngines int `json:"search-engines"`
    SexEducation int `json:"sex-education"`
    SharewareAndFreeware int `json:"shareware-and-freeware"`
    Shopping int `json:"shopping"`
    SocialNetwork int `json:"social-network"`
    Society int `json:"society"`
    SpamUrls int `json:"spam-urls"`
    Sports int `json:"sports"`
    SpywareAndAdware int `json:"spyware-and-adware"`
    StockAdviceAndTools int `json:"stock-advice-and-tools"`
    StreamingMedia int `json:"streaming-media"`
    SwimsuitsAndIntimateApparel int `json:"swimsuits-and-intimate-apparel"`
    TrainingAndTools int `json:"training-and-tools"`
    Translation int `json:"translation"`
    Travel int `json:"travel"`
    Uncategorized int `json:"uncategorized"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Violence int `json:"violence"`
    Weapons int `json:"weapons"`
    WebAdvertisements int `json:"web-advertisements"`
    WebBasedEmail int `json:"web-based-email"`
    WebHostingSites int `json:"web-hosting-sites"`

	} `json:"category-list"`
}


type WebCategoryCategoryListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *WebCategoryCategoryList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *WebCategoryCategoryList) getPath() string{
    return "web-category/category-list"
}

func (p *WebCategoryCategoryList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryCategoryList::Post")
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

func (p *WebCategoryCategoryList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryCategoryList::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *WebCategoryCategoryList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryCategoryList::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *WebCategoryCategoryList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryCategoryList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
