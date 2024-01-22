package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIkeSaBriefOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ike_sa_brief_oper`: Operational Status for the object ike-sa-brief\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIkeSaBriefOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"local_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_sa_brief_remote_gw": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ike_sa_brief_remote_gw_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_brief_remote_gw_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_brief_remote_gw_lifetime": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_brief_remote_gw_status": {
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

func resourceVpnIkeSaBriefOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeSaBriefOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeSaBriefOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIkeSaBriefOperOper := setObjectVpnIkeSaBriefOperOper(res)
		d.Set("oper", VpnIkeSaBriefOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIkeSaBriefOperOper(ret edpt.DataVpnIkeSaBriefOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"name":                   ret.DtVpnIkeSaBriefOper.Oper.Name,
			"local_ip":               ret.DtVpnIkeSaBriefOper.Oper.LocalIp,
			"ike_sa_brief_remote_gw": setSliceVpnIkeSaBriefOperOperIkeSaBriefRemoteGw(ret.DtVpnIkeSaBriefOper.Oper.IkeSaBriefRemoteGw),
		},
	}
}

func setSliceVpnIkeSaBriefOperOperIkeSaBriefRemoteGw(d []edpt.VpnIkeSaBriefOperOperIkeSaBriefRemoteGw) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ike_sa_brief_remote_gw_ip"] = item.IkeSaBriefRemoteGwIp
		in["ike_sa_brief_remote_gw_id"] = item.IkeSaBriefRemoteGwId
		in["ike_sa_brief_remote_gw_lifetime"] = item.IkeSaBriefRemoteGwLifetime
		in["ike_sa_brief_remote_gw_status"] = item.IkeSaBriefRemoteGwStatus
		result = append(result, in)
	}
	return result
}

func getObjectVpnIkeSaBriefOperOper(d []interface{}) edpt.VpnIkeSaBriefOperOper {

	count1 := len(d)
	var ret edpt.VpnIkeSaBriefOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.LocalIp = in["local_ip"].(string)
		ret.IkeSaBriefRemoteGw = getSliceVpnIkeSaBriefOperOperIkeSaBriefRemoteGw(in["ike_sa_brief_remote_gw"].([]interface{}))
	}
	return ret
}

func getSliceVpnIkeSaBriefOperOperIkeSaBriefRemoteGw(d []interface{}) []edpt.VpnIkeSaBriefOperOperIkeSaBriefRemoteGw {

	count1 := len(d)
	ret := make([]edpt.VpnIkeSaBriefOperOperIkeSaBriefRemoteGw, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeSaBriefOperOperIkeSaBriefRemoteGw
		oi.IkeSaBriefRemoteGwIp = in["ike_sa_brief_remote_gw_ip"].(string)
		oi.IkeSaBriefRemoteGwId = in["ike_sa_brief_remote_gw_id"].(string)
		oi.IkeSaBriefRemoteGwLifetime = in["ike_sa_brief_remote_gw_lifetime"].(string)
		oi.IkeSaBriefRemoteGwStatus = in["ike_sa_brief_remote_gw_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnIkeSaBriefOper(d *schema.ResourceData) edpt.VpnIkeSaBriefOper {
	var ret edpt.VpnIkeSaBriefOper

	ret.Oper = getObjectVpnIkeSaBriefOperOper(d.Get("oper").([]interface{}))
	return ret
}
