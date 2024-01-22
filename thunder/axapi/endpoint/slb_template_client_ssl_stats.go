

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateClientSslStats struct {
    
    Name string `json:"name"`
    Stats SlbTemplateClientSslStatsStats `json:"stats"`

}
type DataSlbTemplateClientSslStats struct {
    DtSlbTemplateClientSslStats SlbTemplateClientSslStats `json:"client-ssl"`
}


type SlbTemplateClientSslStatsStats struct {
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
    Cdns int `json:"CDNs"`
    InternetCommunications int `json:"internet-communications"`
    BotNets int `json:"bot-nets"`
    Abortion int `json:"abortion"`
    HealthAndMedicine int `json:"health-and-medicine"`
    ConfirmedSpamSources int `json:"confirmed-SPAM-sources"`
    SpamUrls int `json:"SPAM-URLs"`
    UnconfirmedSpamSources int `json:"unconfirmed-SPAM-sources"`
    OpenHttpProxies int `json:"open-HTTP-proxies"`
    DynamicallyGeneratedContent int `json:"dynamically-generated-content"`
    ParkedDomains int `json:"parked-domains"`
    AlcoholAndTobacco int `json:"alcohol-and-tobacco"`
    PrivateIpAddresses int `json:"private-IP-addresses"`
    ImageAndVideoSearch int `json:"image-and-video-search"`
    FashionAndBeauty int `json:"fashion-and-beauty"`
    RecreationAndHobbies int `json:"recreation-and-hobbies"`
    MotorVehicles int `json:"motor-vehicles"`
    WebHostingSites int `json:"web-hosting-sites"`
    FoodAndDining int `json:"food-and-dining"`
    NudityArtistic int `json:"nudity-artistic"`
    IllegalPornography int `json:"illegal-pornography"`
    Uncategorised int `json:"uncategorised"`
    OtherCategory int `json:"other-category"`
    Trustworthy int `json:"trustworthy"`
    LowRisk int `json:"low-risk"`
    ModerateRisk int `json:"moderate-risk"`
    Suspicious int `json:"suspicious"`
    Malicious int `json:"malicious"`
}

func (p *SlbTemplateClientSslStats) GetId() string{
    return "1"
}

func (p *SlbTemplateClientSslStats) getPath() string{
    
    return "slb/template/client-ssl/"+p.Name+"/stats"
}

func (p *SlbTemplateClientSslStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplateClientSslStats,error) {
logger.Println("SlbTemplateClientSslStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplateClientSslStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
