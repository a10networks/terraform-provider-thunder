

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SctpGlobalStats struct {
    
    Stats SctpGlobalStatsStats `json:"stats"`

}
type DataSctpGlobalStats struct {
    DtSctpGlobalStats SctpGlobalStats `json:"global"`
}


type SctpGlobalStatsStats struct {
    SctpStaticNatSessionCreated int `json:"sctp-static-nat-session-created"`
    SctpStaticNatSessionDeleted int `json:"sctp-static-nat-session-deleted"`
    SctpFwSessionCreated int `json:"sctp-fw-session-created"`
    SctpFwSessionDeleted int `json:"sctp-fw-session-deleted"`
    PktErrDrop int `json:"pkt-err-drop"`
    BadCsum int `json:"bad-csum"`
    BadPayloadDrop int `json:"bad-payload-drop"`
    BadAlignmentDrop int `json:"bad-alignment-drop"`
    OosPktDrop int `json:"oos-pkt-drop"`
    MaxMultiHomeDrop int `json:"max-multi-home-drop"`
    MultiHomeRemoveIpSkip int `json:"multi-home-remove-ip-skip"`
    MultiHomeAddrNotFoundDrop int `json:"multi-home-addr-not-found-drop"`
    StaticNatCfgNotFound int `json:"static-nat-cfg-not-found"`
    CfgErrDrop int `json:"cfg-err-drop"`
    VrrpStandbyDrop int `json:"vrrp-standby-drop"`
    InvalidFragChunkDrop int `json:"invalid-frag-chunk-drop"`
    DisallowedChunkFiltered int `json:"disallowed-chunk-filtered"`
    DisallowedPktDrop int `json:"disallowed-pkt-drop"`
    RateLimitDrop int `json:"rate-limit-drop"`
    SbySessionCreated int `json:"sby-session-created"`
    SbySessionCreateFail int `json:"sby-session-create-fail"`
    SbySessionUpdated int `json:"sby-session-updated"`
    SbySessionUpdateFail int `json:"sby-session-update-fail"`
    SbyStaticNatCfgNotFound int `json:"sby-static-nat-cfg-not-found"`
    SctpChunkTypeInit int `json:"sctp-chunk-type-init"`
    SctpChunkTypeInitAck int `json:"sctp-chunk-type-init-ack"`
    SctpChunkTypeCookieEcho int `json:"sctp-chunk-type-cookie-echo"`
    SctpChunkTypeCookieAck int `json:"sctp-chunk-type-cookie-ack"`
    SctpChunkTypeSack int `json:"sctp-chunk-type-sack"`
    SctpChunkTypeAsconf int `json:"sctp-chunk-type-asconf"`
    SctpChunkTypeAsconfAck int `json:"sctp-chunk-type-asconf-ack"`
    SctpChunkTypeData int `json:"sctp-chunk-type-data"`
    SctpChunkTypeAbort int `json:"sctp-chunk-type-abort"`
    SctpChunkTypeShutdown int `json:"sctp-chunk-type-shutdown"`
    SctpChunkTypeShutdownAck int `json:"sctp-chunk-type-shutdown-ack"`
    SctpChunkTypeShutdownComplete int `json:"sctp-chunk-type-shutdown-complete"`
    SctpChunkTypeErrorOp int `json:"sctp-chunk-type-error-op"`
    SctpChunkTypeHeartbeat int `json:"sctp-chunk-type-heartbeat"`
    SctpChunkTypeHeartbeatAck int `json:"sctp-chunk-type-heartbeat-ack"`
    SctpChunkTypeOther int `json:"sctp-chunk-type-other"`
    SctpHeartbeatNoTuple int `json:"sctp-heartbeat-no-tuple"`
    SctpDataNoTuple int `json:"sctp-data-no-tuple"`
    SctpDataNoExtMatch int `json:"sctp-data-no-ext-match"`
    SctpChunkTypeInitDrop int `json:"sctp-chunk-type-init-drop"`
    SctpChunkTypeInitAckDrop int `json:"sctp-chunk-type-init-ack-drop"`
    SctpChunkTypeShutdownCompleteDrop int `json:"sctp-chunk-type-shutdown-complete-drop"`
    SctpChunkTypeAbortDataDrop int `json:"sctp-chunk-type-abort-data-drop"`
    SctpChunkHeartBeatClubbed int `json:"sctp-chunk-heart-beat-clubbed"`
    SctpRetxInitAckDrop int `json:"sctp-retx-init-ack-drop"`
    SctpRouteErrHeartbeatDrop int `json:"sctp-route-err-heartbeat-drop"`
    SctpRerouteFailover int `json:"sctp-reroute-failover"`
    SctpRouteErrDrop int `json:"sctp-route-err-drop"`
    SctpNoExtMatch int `json:"sctp-no-ext-match"`
    SctpRetxInitAck int `json:"sctp-retx-init-ack"`
    SctpRetxInitDrop int `json:"sctp-retx-init-drop"`
    SctpRetxInit int `json:"sctp-retx-init"`
    SctpAsconfProcessDrop int `json:"sctp-asconf-process-drop"`
    SctpInitVtagZeroDrop int `json:"sctp-init-vtag-zero-drop"`
    PktLenErrDrop int `json:"pkt-len-err-drop"`
    PktChunkLenErrDrop int `json:"pkt-chunk-len-err-drop"`
    PktAsconfParamLenErrDrop int `json:"pkt-asconf-param-len-err-drop"`
}

func (p *SctpGlobalStats) GetId() string{
    return "1"
}

func (p *SctpGlobalStats) getPath() string{
    return "sctp/global/stats"
}

func (p *SctpGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSctpGlobalStats,error) {
logger.Println("SctpGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSctpGlobalStats
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
