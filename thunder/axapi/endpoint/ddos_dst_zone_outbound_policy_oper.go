

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneOutboundPolicyOper struct {
    
    Oper DdosDstZoneOutboundPolicyOperOper `json:"oper"`
    ZoneName string 

}
type DataDdosDstZoneOutboundPolicyOper struct {
    DtDdosDstZoneOutboundPolicyOper DdosDstZoneOutboundPolicyOper `json:"outbound-policy"`
}


type DdosDstZoneOutboundPolicyOperOper struct {
    PolicyName string `json:"policy-name"`
    NoClassListMatch int `json:"no-class-list-match"`
    PolicyClassList []DdosDstZoneOutboundPolicyOperOperPolicyClassList `json:"policy-class-list"`
    GeoTrackingStatistics DdosDstZoneOutboundPolicyOperOperGeoTrackingStatistics `json:"geo-tracking-statistics"`
    TrackingEntryList []DdosDstZoneOutboundPolicyOperOperTrackingEntryList `json:"tracking-entry-list"`
    PolicyRate int `json:"policy-rate"`
    PolicyStatistics int `json:"policy-statistics"`
    TrackingEntryFilter int `json:"tracking-entry-filter"`
}


type DdosDstZoneOutboundPolicyOperOperPolicyClassList struct {
    ClassListName string `json:"class-list-name"`
    CurrentPacketRate string `json:"current-packet-rate"`
    IsPacketRateExceed int `json:"is-packet-rate-exceed"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    IsKbitRateExceed int `json:"is-kBit-rate-exceed"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentConnections string `json:"current-connections"`
    IsConnectionsExceed int `json:"is-connections-exceed"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    IsConnectionRateExceed int `json:"is-connection-rate-exceed"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    IsFragPacketRateExceed int `json:"is-frag-packet-rate-exceed"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    AgeStr string `json:"age-str"`
    LockupTime int `json:"lockup-time"`
    PacketReceived int `json:"packet-received"`
    PacketDropped int `json:"packet-dropped"`
    PacketRateExceed int `json:"packet-rate-exceed"`
    KbitRateExceed int `json:"kBit-rate-exceed"`
    KbitRateExceedCount int `json:"kBit-rate-exceed-count"`
    ConnectionsExceed int `json:"connections-exceed"`
    ConnectionRateExceed int `json:"connection-rate-exceed"`
    FragPacketRate int `json:"frag-packet-rate"`
}


type DdosDstZoneOutboundPolicyOperOperGeoTrackingStatistics struct {
    PacketReceived int `json:"packet-received"`
    PacketDropped int `json:"packet-dropped"`
    PacketRateExceed int `json:"packet-rate-exceed"`
    KbitRateExceed int `json:"kBit-rate-exceed"`
    KbitRateExceedCount int `json:"kBit-rate-exceed-count"`
    ConnectionsExceed int `json:"connections-exceed"`
    ConnectionRateExceed int `json:"connection-rate-exceed"`
    FragPacketRate int `json:"frag-packet-rate"`
    TrackingEntryLearn int `json:"tracking-entry-learn"`
    TrackingEntryAged int `json:"tracking-entry-aged"`
    TrackingEntryLearningThreExceed int `json:"tracking-entry-learning-thre-exceed"`
}


type DdosDstZoneOutboundPolicyOperOperTrackingEntryList struct {
    GeoLocationName string `json:"geo-location-name"`
    CurrentConnections string `json:"current-connections"`
    IsConnectionsExceed int `json:"is-connections-exceed"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    IsConnectionRateExceed int `json:"is-connection-rate-exceed"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    IsPacketRateExceed int `json:"is-packet-rate-exceed"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    IsKbitRateExceed int `json:"is-kBit-rate-exceed"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    IsFragPacketRateExceed int `json:"is-frag-packet-rate-exceed"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    Age int `json:"age"`
}

func (p *DdosDstZoneOutboundPolicyOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneOutboundPolicyOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/outbound-policy/oper"
}

func (p *DdosDstZoneOutboundPolicyOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneOutboundPolicyOper,error) {
logger.Println("DdosDstZoneOutboundPolicyOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneOutboundPolicyOper
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
