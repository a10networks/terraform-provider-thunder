

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfacePacketSizeBucketOper struct {
    
    Oper InterfacePacketSizeBucketOperOper `json:"oper"`

}
type DataInterfacePacketSizeBucketOper struct {
    DtInterfacePacketSizeBucketOper InterfacePacketSizeBucketOper `json:"packet-size-bucket"`
}


type InterfacePacketSizeBucketOperOper struct {
    Interfaces []InterfacePacketSizeBucketOperOperInterfaces `json:"interfaces"`
}


type InterfacePacketSizeBucketOperOperInterfaces struct {
    IfType string `json:"IF-Type"`
    Port_num int `json:"port_num"`
    Rxpkts64_counts int `json:"rxpkts64_counts"`
    Rxpkts65to127_counts int `json:"rxpkts65to127_counts"`
    Rxpkts128to255_counts int `json:"rxpkts128to255_counts"`
    Rxpkts256to511_counts int `json:"rxpkts256to511_counts"`
    Rxpkts512to1023_counts int `json:"rxpkts512to1023_counts"`
    Rxpkts1024to1518_counts int `json:"rxpkts1024to1518_counts"`
    Rxpkts1519to2047_counts int `json:"rxpkts1519to2047_counts"`
    Rxpkts2048to4095_counts int `json:"rxpkts2048to4095_counts"`
    Rxpkts4096to9216_counts int `json:"rxpkts4096to9216_counts"`
    Txpkts64_counts int `json:"txpkts64_counts"`
    Txpkts65to127_counts int `json:"txpkts65to127_counts"`
    Txpkts128to255_counts int `json:"txpkts128to255_counts"`
    Txpkts256to511_counts int `json:"txpkts256to511_counts"`
    Txpkts512to1023_counts int `json:"txpkts512to1023_counts"`
    Txpkts1024to1518_counts int `json:"txpkts1024to1518_counts"`
    Txpkts1519to2047_counts int `json:"txpkts1519to2047_counts"`
    Txpkts2048to4095_counts int `json:"txpkts2048to4095_counts"`
    Txpkts4096to9216_counts int `json:"txpkts4096to9216_counts"`
}

func (p *InterfacePacketSizeBucketOper) GetId() string{
    return "1"
}

func (p *InterfacePacketSizeBucketOper) getPath() string{
    return "interface/packet-size-bucket/oper"
}

func (p *InterfacePacketSizeBucketOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfacePacketSizeBucketOper,error) {
logger.Println("InterfacePacketSizeBucketOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfacePacketSizeBucketOper
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
