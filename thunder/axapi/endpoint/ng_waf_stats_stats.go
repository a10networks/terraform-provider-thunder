

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NgWafStatsStats struct {
    
    Name string `json:"name"`
    Stats NgWafStatsStatsStats `json:"stats"`

}
type DataNgWafStatsStats struct {
    DtNgWafStatsStats NgWafStatsStats `json:"stats"`
}


type NgWafStatsStatsStats struct {
    ReqReceived int `json:"req-received"`
    ReqForward int `json:"req-forward"`
    ReqBlocked int `json:"req-blocked"`
    ReqMarked int `json:"req-marked"`
    ReqBypass int `json:"req-bypass"`
    Error int `json:"error"`
    Useragent int `json:"useragent"`
    AwsSsrf int `json:"aws-ssrf"`
    Backdoor int `json:"backdoor"`
    Cmdexe int `json:"cmdexe"`
    Xss int `json:"xss"`
    Traversal int `json:"traversal"`
    GraphqlDepth int `json:"graphql-depth"`
    Log4jJndi int `json:"log4j-jndi"`
    Sqli int `json:"sqli"`
    Abnormalpath int `json:"abnormalpath"`
    Bhh int `json:"bhh"`
    Block int `json:"block"`
    Codei int `json:"codei"`
    Datacenter int `json:"datacenter"`
    Doubleenc int `json:"doubleenc"`
    DuplicateHeaders int `json:"duplicate-headers"`
    Fbrowing int `json:"fbrowing"`
    GraphqlIde int `json:"graphql-ide"`
    GraphqlIntrospection int `json:"graphql-introspection"`
    GraphqlUnusedVariables int `json:"graphql-unused-variables"`
    Http403 int `json:"http403"`
    Http404 int `json:"http404"`
    Http429 int `json:"http429"`
    Http4xx int `json:"http4xx"`
    Http500 int `json:"http500"`
    Http503 int `json:"http503"`
    Http5xx int `json:"http5xx"`
    Respsplit int `json:"respsplit"`
    Notutf8 int `json:"notutf8"`
    JsonError int `json:"json-error"`
    Malformed int `json:"malformed"`
    Sans int `json:"sans"`
    SigsciIp int `json:"sigsci-ip"`
    NoCtype int `json:"no-ctype"`
    Noua int `json:"noua"`
    Nullbyte int `json:"nullbyte"`
    Privatefile int `json:"privatefile"`
    Scanner int `json:"scanner"`
    Impostor int `json:"impostor"`
    SiteFlaggedIp int `json:"site-flagged-ip"`
    Tornode int `json:"tornode"`
    Weaktls int `json:"weaktls"`
    XmlError int `json:"xml-error"`
    Other int `json:"other"`
    Custom1 int `json:"custom1"`
    Custom2 int `json:"custom2"`
    Custom3 int `json:"custom3"`
    Custom4 int `json:"custom4"`
    Custom5 int `json:"custom5"`
    Custom6 int `json:"custom6"`
    Custom7 int `json:"custom7"`
    Custom8 int `json:"custom8"`
    Custom9 int `json:"custom9"`
    Custom10 int `json:"custom10"`
    Custom11 int `json:"custom11"`
    Custom12 int `json:"custom12"`
    Custom13 int `json:"custom13"`
    Custom14 int `json:"custom14"`
    Custom15 int `json:"custom15"`
    Custom16 int `json:"custom16"`
    Custom17 int `json:"custom17"`
    Custom18 int `json:"custom18"`
    Custom19 int `json:"custom19"`
    Custom20 int `json:"custom20"`
    Custom21 int `json:"custom21"`
    Custom22 int `json:"custom22"`
    Custom23 int `json:"custom23"`
    Custom24 int `json:"custom24"`
    Custom25 int `json:"custom25"`
    Custom26 int `json:"custom26"`
    Custom27 int `json:"custom27"`
    Custom28 int `json:"custom28"`
    Custom29 int `json:"custom29"`
    Custom30 int `json:"custom30"`
    Custom31 int `json:"custom31"`
    Custom32 int `json:"custom32"`
    Custom33 int `json:"custom33"`
    Custom34 int `json:"custom34"`
    Custom35 int `json:"custom35"`
    Custom36 int `json:"custom36"`
    Custom37 int `json:"custom37"`
    Custom38 int `json:"custom38"`
    Custom39 int `json:"custom39"`
    Custom40 int `json:"custom40"`
    Custom41 int `json:"custom41"`
    Custom42 int `json:"custom42"`
    Custom43 int `json:"custom43"`
    Custom44 int `json:"custom44"`
    Custom45 int `json:"custom45"`
    Custom46 int `json:"custom46"`
    Custom47 int `json:"custom47"`
    Custom48 int `json:"custom48"`
    Custom49 int `json:"custom49"`
    Custom50 int `json:"custom50"`
    Custom51 int `json:"custom51"`
    Custom52 int `json:"custom52"`
    Custom53 int `json:"custom53"`
    Custom54 int `json:"custom54"`
    Custom55 int `json:"custom55"`
    Custom56 int `json:"custom56"`
    Custom57 int `json:"custom57"`
    Custom58 int `json:"custom58"`
    Custom59 int `json:"custom59"`
    Custom60 int `json:"custom60"`
    Custom61 int `json:"custom61"`
    Custom62 int `json:"custom62"`
    Custom63 int `json:"custom63"`
    Custom64 int `json:"custom64"`
}

func (p *NgWafStatsStats) GetId() string{
    return "1"
}

func (p *NgWafStatsStats) getPath() string{
    
    return "ng-waf/stats/"+p.Name+"/stats"
}

func (p *NgWafStatsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNgWafStatsStats,error) {
logger.Println("NgWafStatsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNgWafStatsStats
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
