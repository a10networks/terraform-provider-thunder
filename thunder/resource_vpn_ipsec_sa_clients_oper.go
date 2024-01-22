package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsecSaClientsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ipsec_sa_clients_oper`: Operational Status for the object ipsec-sa-clients\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIpsecSaClientsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipsec_clients": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipsec_clients_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sa_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"local_ts": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"in_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"out_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lifetime": {
													Type: schema.TypeString, Optional: true, Description: "",
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
				},
			},
		},
	}
}

func resourceVpnIpsecSaClientsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecSaClientsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecSaClientsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIpsecSaClientsOperOper := setObjectVpnIpsecSaClientsOperOper(res)
		d.Set("oper", VpnIpsecSaClientsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIpsecSaClientsOperOper(ret edpt.DataVpnIpsecSaClientsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ipsec_clients": setSliceVpnIpsecSaClientsOperOperIpsecClients(ret.DtVpnIpsecSaClientsOper.Oper.IpsecClients),
		},
	}
}

func setSliceVpnIpsecSaClientsOperOperIpsecClients(d []edpt.VpnIpsecSaClientsOperOperIpsecClients) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipsec_clients_ip"] = item.IpsecClientsIp
		in["sa_list"] = setSliceVpnIpsecSaClientsOperOperIpsecClientsSaList(item.SaList)
		result = append(result, in)
	}
	return result
}

func setSliceVpnIpsecSaClientsOperOperIpsecClientsSaList(d []edpt.VpnIpsecSaClientsOperOperIpsecClientsSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["local_ts"] = item.LocalTs
		in["in_spi"] = item.InSpi
		in["out_spi"] = item.OutSpi
		in["lifetime"] = item.Lifetime
		in["lifebytes"] = item.Lifebytes
		result = append(result, in)
	}
	return result
}

func getObjectVpnIpsecSaClientsOperOper(d []interface{}) edpt.VpnIpsecSaClientsOperOper {

	count1 := len(d)
	var ret edpt.VpnIpsecSaClientsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpsecClients = getSliceVpnIpsecSaClientsOperOperIpsecClients(in["ipsec_clients"].([]interface{}))
	}
	return ret
}

func getSliceVpnIpsecSaClientsOperOperIpsecClients(d []interface{}) []edpt.VpnIpsecSaClientsOperOperIpsecClients {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecSaClientsOperOperIpsecClients, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecSaClientsOperOperIpsecClients
		oi.IpsecClientsIp = in["ipsec_clients_ip"].(string)
		oi.SaList = getSliceVpnIpsecSaClientsOperOperIpsecClientsSaList(in["sa_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnIpsecSaClientsOperOperIpsecClientsSaList(d []interface{}) []edpt.VpnIpsecSaClientsOperOperIpsecClientsSaList {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecSaClientsOperOperIpsecClientsSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecSaClientsOperOperIpsecClientsSaList
		oi.Name = in["name"].(string)
		oi.LocalTs = in["local_ts"].(string)
		oi.InSpi = in["in_spi"].(string)
		oi.OutSpi = in["out_spi"].(string)
		oi.Lifetime = in["lifetime"].(string)
		oi.Lifebytes = in["lifebytes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnIpsecSaClientsOper(d *schema.ResourceData) edpt.VpnIpsecSaClientsOper {
	var ret edpt.VpnIpsecSaClientsOper

	ret.Oper = getObjectVpnIpsecSaClientsOperOper(d.Get("oper").([]interface{}))
	return ret
}
