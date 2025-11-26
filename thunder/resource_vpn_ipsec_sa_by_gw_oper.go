package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsecSaByGwOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ipsec_sa_by_gw_oper`: Operational Status for the object ipsec-sa-by-gw\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIpsecSaByGwOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ike_gateway_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"local_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"peer_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipsec_sa_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipsec_sa_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"local_ts": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_ts": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"in_spi": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"out_spi": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"encryption": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hash": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lifetime": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lifebytes": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceVpnIpsecSaByGwOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecSaByGwOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecSaByGwOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIpsecSaByGwOperOper := setObjectVpnIpsecSaByGwOperOper(res)
		d.Set("oper", VpnIpsecSaByGwOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIpsecSaByGwOperOper(ret edpt.DataVpnIpsecSaByGwOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ike_gateway_name": ret.DtVpnIpsecSaByGwOper.Oper.IkeGatewayName,
			"local_ip":         ret.DtVpnIpsecSaByGwOper.Oper.LocalIp,
			"peer_ip":          ret.DtVpnIpsecSaByGwOper.Oper.PeerIp,
			"ipsec_sa_list":    setSliceVpnIpsecSaByGwOperOperIpsecSaList(ret.DtVpnIpsecSaByGwOper.Oper.IpsecSaList),
		},
	}
}

func setSliceVpnIpsecSaByGwOperOperIpsecSaList(d []edpt.VpnIpsecSaByGwOperOperIpsecSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipsec_sa_name"] = item.IpsecSaName
		in["local_ts"] = item.LocalTs
		in["remote_ts"] = item.RemoteTs
		in["in_spi"] = item.InSpi
		in["out_spi"] = item.OutSpi
		in["protocol"] = item.Protocol
		in["mode"] = item.Mode
		in["encryption"] = item.Encryption
		in["hash"] = item.Hash
		in["lifetime"] = item.Lifetime
		in["lifebytes"] = item.Lifebytes
		result = append(result, in)
	}
	return result
}

func getObjectVpnIpsecSaByGwOperOper(d []interface{}) edpt.VpnIpsecSaByGwOperOper {

	count1 := len(d)
	var ret edpt.VpnIpsecSaByGwOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IkeGatewayName = in["ike_gateway_name"].(string)
		ret.LocalIp = in["local_ip"].(string)
		ret.PeerIp = in["peer_ip"].(string)
		ret.IpsecSaList = getSliceVpnIpsecSaByGwOperOperIpsecSaList(in["ipsec_sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnIpsecSaByGwOperOperIpsecSaList(d []interface{}) []edpt.VpnIpsecSaByGwOperOperIpsecSaList {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecSaByGwOperOperIpsecSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecSaByGwOperOperIpsecSaList
		oi.IpsecSaName = in["ipsec_sa_name"].(string)
		oi.LocalTs = in["local_ts"].(string)
		oi.RemoteTs = in["remote_ts"].(string)
		oi.InSpi = in["in_spi"].(string)
		oi.OutSpi = in["out_spi"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Mode = in["mode"].(string)
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.Lifetime = in["lifetime"].(int)
		oi.Lifebytes = in["lifebytes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnIpsecSaByGwOper(d *schema.ResourceData) edpt.VpnIpsecSaByGwOper {
	var ret edpt.VpnIpsecSaByGwOper

	ret.Oper = getObjectVpnIpsecSaByGwOperOper(d.Get("oper").([]interface{}))
	return ret
}
