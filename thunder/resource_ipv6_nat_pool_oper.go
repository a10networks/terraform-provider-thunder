package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NatPoolOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ipv6_nat_pool_oper`: Operational Status for the object pool\n\n__PLACEHOLDER__",
		ReadContext: resourceIpv6NatPoolOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nat_pool_addr_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_usage": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name",
			},
		},
	}
}

func resourceIpv6NatPoolOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Ipv6NatPoolOperOper := setObjectIpv6NatPoolOperOper(res)
		d.Set("oper", Ipv6NatPoolOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpv6NatPoolOperOper(ret edpt.DataIpv6NatPoolOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"nat_pool_addr_list": setSliceIpv6NatPoolOperOperNatPoolAddrList(ret.DtIpv6NatPoolOper.Oper.NatPoolAddrList),
		},
	}
}

func setSliceIpv6NatPoolOperOperNatPoolAddrList(d []edpt.Ipv6NatPoolOperOperNatPoolAddrList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["port_usage"] = item.PortUsage
		in["total_used"] = item.TotalUsed
		in["total_freed"] = item.TotalFreed
		in["failed"] = item.Failed
		result = append(result, in)
	}
	return result
}

func getObjectIpv6NatPoolOperOper(d []interface{}) edpt.Ipv6NatPoolOperOper {

	count1 := len(d)
	var ret edpt.Ipv6NatPoolOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatPoolAddrList = getSliceIpv6NatPoolOperOperNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
	}
	return ret
}

func getSliceIpv6NatPoolOperOperNatPoolAddrList(d []interface{}) []edpt.Ipv6NatPoolOperOperNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.Ipv6NatPoolOperOperNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6NatPoolOperOperNatPoolAddrList
		oi.Address = in["address"].(string)
		oi.PortUsage = in["port_usage"].(int)
		oi.TotalUsed = in["total_used"].(int)
		oi.TotalFreed = in["total_freed"].(int)
		oi.Failed = in["failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6NatPoolOper(d *schema.ResourceData) edpt.Ipv6NatPoolOper {
	var ret edpt.Ipv6NatPoolOper

	ret.Oper = getObjectIpv6NatPoolOperOper(d.Get("oper").([]interface{}))

	ret.PoolName = d.Get("pool_name").(string)
	return ret
}
