package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_scaleout_address_mapping_nat_pool_or_fixed_nat_oper`: Operational Status for the object nat-pool-or-fixed-nat\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_template": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_group": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_node": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"standby_node": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"index": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper := setObjectCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper(res)
		d.Set("oper", Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper(ret edpt.DataCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"service_template": ret.DtCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper.Oper.ServiceTemplate,
			"ip_list":          setSliceCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList(ret.DtCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper.Oper.IpList),
			"nat_pool":         ret.DtCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper.Oper.NatPool,
			"index":            ret.DtCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper.Oper.Index,
		},
	}
}

func setSliceCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList(d []edpt.Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["user_group"] = item.UserGroup
		in["active_node"] = item.ActiveNode
		in["standby_node"] = item.StandbyNode
		in["nat_ip"] = item.NatIp
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper(d []interface{}) edpt.Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceTemplate = in["service_template"].(string)
		ret.IpList = getSliceCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList(in["ip_list"].([]interface{}))
		ret.NatPool = in["nat_pool"].(string)
		ret.Index = in["index"].(int)
	}
	return ret
}

func getSliceCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList(d []interface{}) []edpt.Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList
		oi.UserGroup = in["user_group"].(int)
		oi.ActiveNode = in["active_node"].(int)
		oi.StandbyNode = in["standby_node"].(int)
		oi.NatIp = in["nat_ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper(d *schema.ResourceData) edpt.Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper {
	var ret edpt.Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper

	ret.Oper = getObjectCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper(d.Get("oper").([]interface{}))
	return ret
}
