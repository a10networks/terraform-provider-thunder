package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOneSharedPoolOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_one_to_one_shared_pool_oper`: Operational Status for the object shared-pool\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6OneToOneSharedPoolOperRead,

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

func resourceCgnv6OneToOneSharedPoolOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneSharedPoolOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneSharedPoolOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6OneToOneSharedPoolOperOper := setObjectCgnv6OneToOneSharedPoolOperOper(res)
		d.Set("oper", Cgnv6OneToOneSharedPoolOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6OneToOneSharedPoolOperOper(ret edpt.DataCgnv6OneToOneSharedPoolOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"shared_pool_list": setSliceCgnv6OneToOneSharedPoolOperOperSharedPoolList(ret.DtCgnv6OneToOneSharedPoolOper.Oper.SharedPoolList),
		},
	}
}

func setSliceCgnv6OneToOneSharedPoolOperOperSharedPoolList(d []edpt.Cgnv6OneToOneSharedPoolOperOperSharedPoolList) []map[string]interface{} {
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

func getObjectCgnv6OneToOneSharedPoolOperOper(d []interface{}) edpt.Cgnv6OneToOneSharedPoolOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6OneToOneSharedPoolOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SharedPoolList = getSliceCgnv6OneToOneSharedPoolOperOperSharedPoolList(in["shared_pool_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6OneToOneSharedPoolOperOperSharedPoolList(d []interface{}) []edpt.Cgnv6OneToOneSharedPoolOperOperSharedPoolList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6OneToOneSharedPoolOperOperSharedPoolList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6OneToOneSharedPoolOperOperSharedPoolList
		oi.PoolName = in["pool_name"].(string)
		oi.StartAddress = in["start_address"].(string)
		oi.EndAddress = in["end_address"].(string)
		oi.Netmask = in["netmask"].(string)
		oi.Vird = in["vird"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6OneToOneSharedPoolOper(d *schema.ResourceData) edpt.Cgnv6OneToOneSharedPoolOper {
	var ret edpt.Cgnv6OneToOneSharedPoolOper

	ret.Oper = getObjectCgnv6OneToOneSharedPoolOperOper(d.Get("oper").([]interface{}))
	return ret
}
