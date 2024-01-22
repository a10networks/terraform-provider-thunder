

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTcpStatsOper struct {
    
    Oper SystemTcpStatsOperOper `json:"oper"`

}
type DataSystemTcpStatsOper struct {
    DtSystemTcpStatsOper SystemTcpStatsOper `json:"tcp-stats"`
}


type SystemTcpStatsOperOper struct {
    TcpstatsCpuList []SystemTcpStatsOperOperTcpstatsCpuList `json:"tcpstats-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SystemTcpStatsOperOperTcpstatsCpuList struct {
    Connattempt int `json:"connattempt"`
    Connects int `json:"connects"`
    Drops int `json:"drops"`
    Conndrops int `json:"conndrops"`
    Closed int `json:"closed"`
    Segstimed int `json:"segstimed"`
    Rttupdated int `json:"rttupdated"`
    Delack int `json:"delack"`
    Timeoutdrop int `json:"timeoutdrop"`
    Rexmttimeo int `json:"rexmttimeo"`
    Persisttimeo int `json:"persisttimeo"`
    Keeptimeo int `json:"keeptimeo"`
    Keepprobe int `json:"keepprobe"`
    Keepdrops int `json:"keepdrops"`
    Sndtotal int `json:"sndtotal"`
    Sndpack int `json:"sndpack"`
    Sndbyte int `json:"sndbyte"`
    Sndrexmitpack int `json:"sndrexmitpack"`
    Sndrexmitbyte int `json:"sndrexmitbyte"`
    Sndrexmitbad int `json:"sndrexmitbad"`
    Sndacks int `json:"sndacks"`
    Sndprobe int `json:"sndprobe"`
    Sndurg int `json:"sndurg"`
    Sndwinup int `json:"sndwinup"`
    Sndctrl int `json:"sndctrl"`
    Rcvtotal int `json:"rcvtotal"`
    Rcvpack int `json:"rcvpack"`
    Rcvbyte int `json:"rcvbyte"`
    Rcvbadoff int `json:"rcvbadoff"`
    Rcvmemdrop int `json:"rcvmemdrop"`
    Reassmemdrop int `json:"reassmemdrop"`
    Reasstimeout int `json:"reasstimeout"`
    Rcvduppack int `json:"rcvduppack"`
    Rcvdupbyte int `json:"rcvdupbyte"`
    Rcvpartduppack int `json:"rcvpartduppack"`
    Rcvpartdupbyte int `json:"rcvpartdupbyte"`
    Rcvoopack int `json:"rcvoopack"`
    Rcvoobyte int `json:"rcvoobyte"`
    Rcvpackafterwin int `json:"rcvpackafterwin"`
    Rcvbyteafterwin int `json:"rcvbyteafterwin"`
    Rcvwinprobe int `json:"rcvwinprobe"`
    Rcvdupack int `json:"rcvdupack"`
    Rcvacktoomuch int `json:"rcvacktoomuch"`
    Rcvackpack int `json:"rcvackpack"`
    Rcvackbyte int `json:"rcvackbyte"`
    Rcvwinupd int `json:"rcvwinupd"`
    Pawsdrop int `json:"pawsdrop"`
    Predack int `json:"predack"`
    Preddat int `json:"preddat"`
    Persistdrop int `json:"persistdrop"`
    Badrst int `json:"badrst"`
    Finwait2_drops int `json:"finwait2_drops"`
    Rcvdsack int `json:"rcvdsack"`
    Sack_recovery_episode int `json:"sack_recovery_episode"`
    Sack_rexmits int `json:"sack_rexmits"`
    Sack_rexmit_bytes int `json:"sack_rexmit_bytes"`
    Sack_rcv_blocks int `json:"sack_rcv_blocks"`
    Sack_send_blocks int `json:"sack_send_blocks"`
    Sndrst int `json:"sndrst"`
    Sndfin int `json:"sndfin"`
    Sndsyn int `json:"sndsyn"`
    Sndcack int `json:"sndcack"`
    Cacklim int `json:"cacklim"`
    Cc_idle int `json:"cc_idle"`
    Cc_reduce int `json:"cc_reduce"`
    Bad_iochan int `json:"bad_iochan"`
    A2btcpoptions int `json:"a2btcpoptions"`
    A2brcvwnd int `json:"a2brcvwnd"`
    A2bsackpresent int `json:"a2bsackpresent"`
    A2bdupack int `json:"a2bdupack"`
    A2brxdata int `json:"a2brxdata"`
    A2boodata int `json:"a2boodata"`
    A2bpartialack int `json:"a2bpartialack"`
    A2bfsmtransition int `json:"a2bfsmtransition"`
    A2btransitionnum int `json:"a2btransitionnum"`
    B2atransitionnum int `json:"b2atransitionnum"`
    Atcpforward int `json:"atcpforward"`
    Atcpsent int `json:"atcpsent"`
    Atcpsendbackack int `json:"atcpsendbackack"`
    Atcprexmit int `json:"atcprexmit"`
    Atcpbuffallocfail int `json:"atcpbuffallocfail"`
    A2bappbuffering int `json:"a2bappbuffering"`
    Atcpsendfail int `json:"atcpsendfail"`
    Earlyrexmit int `json:"earlyrexmit"`
    Mburstlim int `json:"mburstlim"`
    A2bsndwnd int `json:"a2bsndwnd"`
    Proxyheaderv1 int `json:"proxyheaderv1"`
    Proxyheaderv2 int `json:"proxyheaderv2"`
}

func (p *SystemTcpStatsOper) GetId() string{
    return "1"
}

func (p *SystemTcpStatsOper) getPath() string{
    return "system/tcp-stats/oper"
}

func (p *SystemTcpStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTcpStatsOper,error) {
logger.Println("SystemTcpStatsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTcpStatsOper
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
