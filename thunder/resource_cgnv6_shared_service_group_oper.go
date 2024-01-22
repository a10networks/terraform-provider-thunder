package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SharedServiceGroupOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_shared_service_group_oper`: Operational Status for the object shared-service-group\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6SharedServiceGroupOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"shared_service_group_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_group_name": {
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

func resourceCgnv6SharedServiceGroupOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SharedServiceGroupOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SharedServiceGroupOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6SharedServiceGroupOperOper := setObjectCgnv6SharedServiceGroupOperOper(res)
		d.Set("oper", Cgnv6SharedServiceGroupOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6SharedServiceGroupOperOper(ret edpt.DataCgnv6SharedServiceGroupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"shared_service_group_list": setSliceCgnv6SharedServiceGroupOperOperSharedServiceGroupList(ret.DtCgnv6SharedServiceGroupOper.Oper.SharedServiceGroupList),
		},
	}
}

func setSliceCgnv6SharedServiceGroupOperOperSharedServiceGroupList(d []edpt.Cgnv6SharedServiceGroupOperOperSharedServiceGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["service_group_name"] = item.ServiceGroupName
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6SharedServiceGroupOperOper(d []interface{}) edpt.Cgnv6SharedServiceGroupOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6SharedServiceGroupOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SharedServiceGroupList = getSliceCgnv6SharedServiceGroupOperOperSharedServiceGroupList(in["shared_service_group_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6SharedServiceGroupOperOperSharedServiceGroupList(d []interface{}) []edpt.Cgnv6SharedServiceGroupOperOperSharedServiceGroupList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6SharedServiceGroupOperOperSharedServiceGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6SharedServiceGroupOperOperSharedServiceGroupList
		oi.ServiceGroupName = in["service_group_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6SharedServiceGroupOper(d *schema.ResourceData) edpt.Cgnv6SharedServiceGroupOper {
	var ret edpt.Cgnv6SharedServiceGroupOper

	ret.Oper = getObjectCgnv6SharedServiceGroupOperOper(d.Get("oper").([]interface{}))
	return ret
}
