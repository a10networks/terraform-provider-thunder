

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemDpdkStatsStats struct {
    
    Stats SystemDpdkStatsStatsStats `json:"stats"`

}
type DataSystemDpdkStatsStats struct {
    DtSystemDpdkStatsStats SystemDpdkStatsStats `json:"dpdk-stats"`
}


type SystemDpdkStatsStatsStats struct {
    PktDrop int `json:"pkt-drop"`
    PktLnkDownDrop int `json:"pkt-lnk-down-drop"`
    ErrPktDrop int `json:"err-pkt-drop"`
    RxErr int `json:"rx-err"`
    TxErr int `json:"tx-err"`
    TxDrop int `json:"tx-drop"`
    RxLenErr int `json:"rx-len-err"`
    RxOverErr int `json:"rx-over-err"`
    RxCrcErr int `json:"rx-crc-err"`
    RxFrameErr int `json:"rx-frame-err"`
    RxNoBuffErr int `json:"rx-no-buff-err"`
    RxMissErr int `json:"rx-miss-err"`
    TxAbortErr int `json:"tx-abort-err"`
    TxCarrierErr int `json:"tx-carrier-err"`
    TxFifoErr int `json:"tx-fifo-err"`
    TxHbeatErr int `json:"tx-hbeat-err"`
    TxWindowsErr int `json:"tx-windows-err"`
    RxLongLenErr int `json:"rx-long-len-err"`
    RxShortLenErr int `json:"rx-short-len-err"`
    RxAlignErr int `json:"rx-align-err"`
    RxCsumOffloadErr int `json:"rx-csum-offload-err"`
    IoRxQueDrop int `json:"io-rx-que-drop"`
    IoTxQueDrop int `json:"io-tx-que-drop"`
    IoRingDrop int `json:"io-ring-drop"`
    WTxQueDrop int `json:"w-tx-que-drop"`
    WLinkDownDrop int `json:"w-link-down-drop"`
    WRingDrop int `json:"w-ring-drop"`
}

func (p *SystemDpdkStatsStats) GetId() string{
    return "1"
}

func (p *SystemDpdkStatsStats) getPath() string{
    return "system/dpdk-stats/stats"
}

func (p *SystemDpdkStatsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemDpdkStatsStats,error) {
logger.Println("SystemDpdkStatsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemDpdkStatsStats
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
