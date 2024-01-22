

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type RouteMapSet struct {
	Inst struct {

    Aggregator RouteMapSetAggregator `json:"aggregator"`
    AsPath RouteMapSetAsPath `json:"as-path"`
    AtomicAggregate int `json:"atomic-aggregate"`
    CommList RouteMapSetCommList `json:"comm-list"`
    Community string `json:"community"`
    DampeningCfg RouteMapSetDampeningCfg `json:"dampening-cfg"`
    Ddos RouteMapSetDdos `json:"ddos"`
    Extcommunity RouteMapSetExtcommunity `json:"extcommunity"`
    Ip RouteMapSetIp `json:"ip"`
    Ipv6 RouteMapSetIpv6 `json:"ipv6"`
    LargeCommList RouteMapSetLargeCommList `json:"large-comm-list"`
    LargeCommunity string `json:"large-community"`
    Level RouteMapSetLevel `json:"level"`
    LocalPreference RouteMapSetLocalPreference `json:"local-preference"`
    Metric RouteMapSetMetric `json:"metric"`
    MetricType RouteMapSetMetricType `json:"metric-type"`
    Origin RouteMapSetOrigin `json:"origin"`
    OriginatorId RouteMapSetOriginatorId `json:"originator-id"`
    Tag RouteMapSetTag `json:"tag"`
    Uuid string `json:"uuid"`
    Weight RouteMapSetWeight `json:"weight"`
    Sequence string 
    Action string 

	} `json:"set"`
}


type RouteMapSetAggregator struct {
    AggregatorAs RouteMapSetAggregatorAggregatorAs `json:"aggregator-as"`
}


type RouteMapSetAggregatorAggregatorAs struct {
    Asn int `json:"asn"`
    Ip string `json:"ip"`
}


type RouteMapSetAsPath struct {
    Prepend string `json:"prepend"`
    Num string `json:"num"`
    Num2 string `json:"num2"`
}


type RouteMapSetCommList struct {
    VStd int `json:"v-std"`
    Delete int `json:"delete"`
    VExp int `json:"v-exp"`
    VExpDelete int `json:"v-exp-delete"`
    Name string `json:"name"`
    NameDelete int `json:"name-delete"`
}


type RouteMapSetDampeningCfg struct {
    Dampening int `json:"dampening"`
    DampeningHalfTime int `json:"dampening-half-time"`
    DampeningReuse int `json:"dampening-reuse"`
    DampeningSupress int `json:"dampening-supress"`
    DampeningMaxSupress int `json:"dampening-max-supress"`
    DampeningPenalty int `json:"dampening-penalty"`
}


type RouteMapSetDdos struct {
    ClassListName string `json:"class-list-name"`
    ClassListCid int `json:"class-list-cid"`
    Zone string `json:"zone"`
}


type RouteMapSetExtcommunity struct {
    Rt RouteMapSetExtcommunityRt `json:"rt"`
    Soo RouteMapSetExtcommunitySoo `json:"soo"`
}


type RouteMapSetExtcommunityRt struct {
    Value string `json:"value"`
}


type RouteMapSetExtcommunitySoo struct {
    Value string `json:"value"`
}


type RouteMapSetIp struct {
    NextHop RouteMapSetIpNextHop `json:"next-hop"`
}


type RouteMapSetIpNextHop struct {
    Address string `json:"address"`
}


type RouteMapSetIpv6 struct {
    NextHop1 RouteMapSetIpv6NextHop1 `json:"next-hop-1"`
}


type RouteMapSetIpv6NextHop1 struct {
    Address string `json:"address"`
    Local RouteMapSetIpv6NextHop1Local `json:"local"`
}


type RouteMapSetIpv6NextHop1Local struct {
    Address string `json:"address"`
}


type RouteMapSetLargeCommList struct {
    LVStd int `json:"l-v-std"`
    LVStdDelete int `json:"l-v-std-delete"`
    LVExp int `json:"l-v-exp"`
    LVExpDelete int `json:"l-v-exp-delete"`
    LName string `json:"l-name"`
    LargeNameDelete int `json:"large-name-delete"`
}


type RouteMapSetLevel struct {
    Value string `json:"value"`
}


type RouteMapSetLocalPreference struct {
    Val int `json:"val"`
}


type RouteMapSetMetric struct {
    Value string `json:"value"`
}


type RouteMapSetMetricType struct {
    Value string `json:"value"`
}


type RouteMapSetOrigin struct {
    Egp int `json:"egp"`
    Igp int `json:"igp"`
    Incomplete int `json:"incomplete"`
}


type RouteMapSetOriginatorId struct {
    OriginatorIp string `json:"originator-ip"`
}


type RouteMapSetTag struct {
    Value int `json:"value"`
}


type RouteMapSetWeight struct {
    WeightVal int `json:"weight-val"`
}

func (p *RouteMapSet) GetId() string{
    return "1"
}

func (p *RouteMapSet) getPath() string{
    return "route-map/"+strconv.Itoa(p.Inst.Tag.Value)+"+" +p.Inst.Action + "+" +p.Inst.Sequence + "/set"
}

func (p *RouteMapSet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMapSet::Post")
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

func (p *RouteMapSet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMapSet::Get")
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
func (p *RouteMapSet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMapSet::Put")
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

func (p *RouteMapSet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMapSet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
