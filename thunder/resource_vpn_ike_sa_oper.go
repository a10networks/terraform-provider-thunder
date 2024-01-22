package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIkeSaOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ike_sa_oper`: Operational Status for the object ike-sa\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIkeSaOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ike_sa_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"initiator_spi": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"responder_spi": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"local_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_ip": {
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
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_traversal": {
										Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceVpnIkeSaOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeSaOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeSaOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIkeSaOperOper := setObjectVpnIkeSaOperOper(res)
		d.Set("oper", VpnIkeSaOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIkeSaOperOper(ret edpt.DataVpnIkeSaOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ike_sa_list": setSliceVpnIkeSaOperOperIkeSaList(ret.DtVpnIkeSaOper.Oper.IkeSaList),
		},
	}
}

func setSliceVpnIkeSaOperOperIkeSaList(d []edpt.VpnIkeSaOperOperIkeSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["initiator_spi"] = item.InitiatorSpi
		in["responder_spi"] = item.ResponderSpi
		in["local_ip"] = item.LocalIp
		in["remote_ip"] = item.RemoteIp
		in["encryption"] = item.Encryption
		in["hash"] = item.Hash
		in["lifetime"] = item.Lifetime
		in["status"] = item.Status
		in["nat_traversal"] = item.NatTraversal
		result = append(result, in)
	}
	return result
}

func getObjectVpnIkeSaOperOper(d []interface{}) edpt.VpnIkeSaOperOper {

	count1 := len(d)
	var ret edpt.VpnIkeSaOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IkeSaList = getSliceVpnIkeSaOperOperIkeSaList(in["ike_sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnIkeSaOperOperIkeSaList(d []interface{}) []edpt.VpnIkeSaOperOperIkeSaList {

	count1 := len(d)
	ret := make([]edpt.VpnIkeSaOperOperIkeSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeSaOperOperIkeSaList
		oi.Name = in["name"].(string)
		oi.InitiatorSpi = in["initiator_spi"].(string)
		oi.ResponderSpi = in["responder_spi"].(string)
		oi.LocalIp = in["local_ip"].(string)
		oi.RemoteIp = in["remote_ip"].(string)
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.Lifetime = in["lifetime"].(int)
		oi.Status = in["status"].(string)
		oi.NatTraversal = in["nat_traversal"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnIkeSaOper(d *schema.ResourceData) edpt.VpnIkeSaOper {
	var ret edpt.VpnIkeSaOper

	ret.Oper = getObjectVpnIkeSaOperOper(d.Get("oper").([]interface{}))
	return ret
}
