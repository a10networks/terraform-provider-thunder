package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatSharedPoolOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat_shared_pool_oper`: Operational Status for the object shared-pool\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6NatSharedPoolOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"shared_pool_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pool_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"start_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"end_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"netmask": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vird": {
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

func resourceCgnv6NatSharedPoolOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatSharedPoolOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatSharedPoolOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6NatSharedPoolOperOper := setObjectCgnv6NatSharedPoolOperOper(res)
		d.Set("oper", Cgnv6NatSharedPoolOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6NatSharedPoolOperOper(ret edpt.DataCgnv6NatSharedPoolOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"shared_pool_list": setSliceCgnv6NatSharedPoolOperOperSharedPoolList(ret.DtCgnv6NatSharedPoolOper.Oper.SharedPoolList),
		},
	}
}

func setSliceCgnv6NatSharedPoolOperOperSharedPoolList(d []edpt.Cgnv6NatSharedPoolOperOperSharedPoolList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["pool_name"] = item.PoolName
		in["start_address"] = item.StartAddress
		in["end_address"] = item.EndAddress
		in["netmask"] = item.Netmask
		in["vird"] = item.Vird
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6NatSharedPoolOperOper(d []interface{}) edpt.Cgnv6NatSharedPoolOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6NatSharedPoolOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SharedPoolList = getSliceCgnv6NatSharedPoolOperOperSharedPoolList(in["shared_pool_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6NatSharedPoolOperOperSharedPoolList(d []interface{}) []edpt.Cgnv6NatSharedPoolOperOperSharedPoolList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6NatSharedPoolOperOperSharedPoolList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6NatSharedPoolOperOperSharedPoolList
		oi.PoolName = in["pool_name"].(string)
		oi.StartAddress = in["start_address"].(string)
		oi.EndAddress = in["end_address"].(string)
		oi.Netmask = in["netmask"].(string)
		oi.Vird = in["vird"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6NatSharedPoolOper(d *schema.ResourceData) edpt.Cgnv6NatSharedPoolOper {
	var ret edpt.Cgnv6NatSharedPoolOper

	ret.Oper = getObjectCgnv6NatSharedPoolOperOper(d.Get("oper").([]interface{}))
	return ret
}
