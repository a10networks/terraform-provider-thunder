

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIkeGatewayStats struct {
    
    Name string `json:"name"`
    Stats VpnIkeGatewayStatsStats `json:"stats"`

}
type DataVpnIkeGatewayStats struct {
    DtVpnIkeGatewayStats VpnIkeGatewayStats `json:"ike-gateway"`
}


type VpnIkeGatewayStatsStats struct {
    V2InitRekey int `json:"v2-init-rekey"`
    V2RspRekey int `json:"v2-rsp-rekey"`
    V2ChildSaRekey int `json:"v2-child-sa-rekey"`
    V2InInvalid int `json:"v2-in-invalid"`
    V2InInvalidSpi int `json:"v2-in-invalid-spi"`
    V2InInitReq int `json:"v2-in-init-req"`
    V2InInitRsp int `json:"v2-in-init-rsp"`
    V2OutInitReq int `json:"v2-out-init-req"`
    V2OutInitRsp int `json:"v2-out-init-rsp"`
    V2InAuthReq int `json:"v2-in-auth-req"`
    V2InAuthRsp int `json:"v2-in-auth-rsp"`
    V2OutAuthReq int `json:"v2-out-auth-req"`
    V2OutAuthRsp int `json:"v2-out-auth-rsp"`
    V2InCreateChildReq int `json:"v2-in-create-child-req"`
    V2InCreateChildRsp int `json:"v2-in-create-child-rsp"`
    V2OutCreateChildReq int `json:"v2-out-create-child-req"`
    V2OutCreateChildRsp int `json:"v2-out-create-child-rsp"`
    V2InInfoReq int `json:"v2-in-info-req"`
    V2InInfoRsp int `json:"v2-in-info-rsp"`
    V2OutInfoReq int `json:"v2-out-info-req"`
    V2OutInfoRsp int `json:"v2-out-info-rsp"`
    V1InIdProtReq int `json:"v1-in-id-prot-req"`
    V1InIdProtRsp int `json:"v1-in-id-prot-rsp"`
    V1OutIdProtReq int `json:"v1-out-id-prot-req"`
    V1OutIdProtRsp int `json:"v1-out-id-prot-rsp"`
    V1InAuthOnlyReq int `json:"v1-in-auth-only-req"`
    V1InAuthOnlyRsp int `json:"v1-in-auth-only-rsp"`
    V1OutAuthOnlyReq int `json:"v1-out-auth-only-req"`
    V1OutAuthOnlyRsp int `json:"v1-out-auth-only-rsp"`
    V1InAggressiveReq int `json:"v1-in-aggressive-req"`
    V1InAggressiveRsp int `json:"v1-in-aggressive-rsp"`
    V1OutAggressiveReq int `json:"v1-out-aggressive-req"`
    V1OutAggressiveRsp int `json:"v1-out-aggressive-rsp"`
    V1InInfoV1Req int `json:"v1-in-info-v1-req"`
    V1InInfoV1Rsp int `json:"v1-in-info-v1-rsp"`
    V1OutInfoV1Req int `json:"v1-out-info-v1-req"`
    V1OutInfoV1Rsp int `json:"v1-out-info-v1-rsp"`
    V1InTransactionReq int `json:"v1-in-transaction-req"`
    V1InTransactionRsp int `json:"v1-in-transaction-rsp"`
    V1OutTransactionReq int `json:"v1-out-transaction-req"`
    V1OutTransactionRsp int `json:"v1-out-transaction-rsp"`
    V1InQuickModeReq int `json:"v1-in-quick-mode-req"`
    V1InQuickModeRsp int `json:"v1-in-quick-mode-rsp"`
    V1OutQuickModeReq int `json:"v1-out-quick-mode-req"`
    V1OutQuickModeRsp int `json:"v1-out-quick-mode-rsp"`
    V1InNewGroupModeReq int `json:"v1-in-new-group-mode-req"`
    V1InNewGroupModeRsp int `json:"v1-in-new-group-mode-rsp"`
    V1OutNewGroupModeReq int `json:"v1-out-new-group-mode-req"`
    V1OutNewGroupModeRsp int `json:"v1-out-new-group-mode-rsp"`
    V1ChildSaInvalidSpi int `json:"v1-child-sa-invalid-spi"`
    V2ChildSaInvalidSpi int `json:"v2-child-sa-invalid-spi"`
    IkeCurrentVersion int `json:"ike-current-version"`
}

func (p *VpnIkeGatewayStats) GetId() string{
    return "1"
}

func (p *VpnIkeGatewayStats) getPath() string{
    
    return "vpn/ike-gateway/"+p.Name+"/stats"
}

func (p *VpnIkeGatewayStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIkeGatewayStats,error) {
logger.Println("VpnIkeGatewayStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIkeGatewayStats
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
