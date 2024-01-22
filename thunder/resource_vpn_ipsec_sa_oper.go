package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsecSaOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ipsec_sa_oper`: Operational Status for the object ipsec-sa\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIpsecSaOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipsec_sa_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipsec_sa_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_gateway_name": {
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

func resourceVpnIpsecSaOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecSaOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecSaOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIpsecSaOperOper := setObjectVpnIpsecSaOperOper(res)
		d.Set("oper", VpnIpsecSaOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIpsecSaOperOper(ret edpt.DataVpnIpsecSaOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ipsec_sa_list": setSliceVpnIpsecSaOperOperIpsecSaList(ret.DtVpnIpsecSaOper.Oper.IpsecSaList),
		},
	}
}

func setSliceVpnIpsecSaOperOperIpsecSaList(d []edpt.VpnIpsecSaOperOperIpsecSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipsec_sa_name"] = item.IpsecSaName
		in["ike_gateway_name"] = item.IkeGatewayName
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

func getObjectVpnIpsecSaOperOper(d []interface{}) edpt.VpnIpsecSaOperOper {

	count1 := len(d)
	var ret edpt.VpnIpsecSaOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpsecSaList = getSliceVpnIpsecSaOperOperIpsecSaList(in["ipsec_sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnIpsecSaOperOperIpsecSaList(d []interface{}) []edpt.VpnIpsecSaOperOperIpsecSaList {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecSaOperOperIpsecSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecSaOperOperIpsecSaList
		oi.IpsecSaName = in["ipsec_sa_name"].(string)
		oi.IkeGatewayName = in["ike_gateway_name"].(string)
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

func dataToEndpointVpnIpsecSaOper(d *schema.ResourceData) edpt.VpnIpsecSaOper {
	var ret edpt.VpnIpsecSaOper

	ret.Oper = getObjectVpnIpsecSaOperOper(d.Get("oper").([]interface{}))
	return ret
}
