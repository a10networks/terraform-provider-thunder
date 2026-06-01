package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type OverlayTunnelDebugEncapTableOper struct {
	Oper OverlayTunnelDebugEncapTableOperOper `json:"oper"`
}
type DataOverlayTunnelDebugEncapTableOper struct {
	DtOverlayTunnelDebugEncapTableOper OverlayTunnelDebugEncapTableOper `json:"encap-table"`
}

type OverlayTunnelDebugEncapTableOperOper struct {
	Vtep           []OverlayTunnelDebugEncapTableOperOperVtep `json:"vtep"`
	Tot_ip_encap   int                                        `json:"tot_ip_encap"`
	Tot_ipv6_encap int                                        `json:"tot_ipv6_encap"`
}

type OverlayTunnelDebugEncapTableOperOperVtep struct {
	Index         int    `json:"index"`
	Hindex        int    `json:"hindex"`
	Src_vtep_ip   string `json:"src_vtep_ip"`
	Dst_vtep_ip   string `json:"dst_vtep_ip"`
	Src_vtep_ipv6 string `json:"src_vtep_ipv6"`
	Dst_vtep_ipv6 string `json:"dst_vtep_ipv6"`
	Is_ipv6       int    `json:"is_ipv6"`
	Dst_vtep_mac  string `json:"dst_vtep_mac"`
	Encap_type    string `json:"encap_type"`
	Vtep_id       int    `json:"vtep_id"`
	Vtep_vnp_id   int    `json:"vtep_vnp_id"`
	Lifname       string `json:"lifname"`
	Partname      string `json:"partname"`
	Vlan          int    `json:"vlan"`
	Is_static     int    `json:"is_static"`
	Age           int    `json:"age"`
}

func (p *OverlayTunnelDebugEncapTableOper) GetId() string {
	return "1"
}

func (p *OverlayTunnelDebugEncapTableOper) getPath() string {
	return "overlay-tunnel/debug/encap-table/oper"
}

func (p *OverlayTunnelDebugEncapTableOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataOverlayTunnelDebugEncapTableOper, error) {
	logger.Println("OverlayTunnelDebugEncapTableOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataOverlayTunnelDebugEncapTableOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
