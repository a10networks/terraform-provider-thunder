package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatPoolOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_nat_pool_oper`: Operational Status for the object pool\n\n__PLACEHOLDER__",
		ReadContext: resourceIpNatPoolOperRead,

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
				Type: schema.TypeString, Required: true, Description: "Specify pool name or pool group",
			},
		},
	}
}

func resourceIpNatPoolOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpNatPoolOperOper := setObjectIpNatPoolOperOper(res)
		d.Set("oper", IpNatPoolOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpNatPoolOperOper(ret edpt.DataIpNatPoolOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"nat_pool_addr_list": setSliceIpNatPoolOperOperNatPoolAddrList(ret.DtIpNatPoolOper.Oper.NatPoolAddrList),
		},
	}
}

func setSliceIpNatPoolOperOperNatPoolAddrList(d []edpt.IpNatPoolOperOperNatPoolAddrList) []map[string]interface{} {
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

func getObjectIpNatPoolOperOper(d []interface{}) edpt.IpNatPoolOperOper {

	count1 := len(d)
	var ret edpt.IpNatPoolOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatPoolAddrList = getSliceIpNatPoolOperOperNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
	}
	return ret
}

func getSliceIpNatPoolOperOperNatPoolAddrList(d []interface{}) []edpt.IpNatPoolOperOperNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.IpNatPoolOperOperNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpNatPoolOperOperNatPoolAddrList
		oi.Address = in["address"].(string)
		oi.PortUsage = in["port_usage"].(int)
		oi.TotalUsed = in["total_used"].(int)
		oi.TotalFreed = in["total_freed"].(int)
		oi.Failed = in["failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpNatPoolOper(d *schema.ResourceData) edpt.IpNatPoolOper {
	var ret edpt.IpNatPoolOper

	ret.Oper = getObjectIpNatPoolOperOper(d.Get("oper").([]interface{}))

	ret.PoolName = d.Get("pool_name").(string)
	return ret
}
