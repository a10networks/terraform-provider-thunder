package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ScaleoutAddressMappingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_scaleout_address_mapping_oper`: Operational Status for the object address-mapping\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6ScaleoutAddressMappingOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
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
						"ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_nat_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"end_nat_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6ScaleoutAddressMappingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ScaleoutAddressMappingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ScaleoutAddressMappingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6ScaleoutAddressMappingOperOper := setObjectCgnv6ScaleoutAddressMappingOperOper(res)
		d.Set("oper", Cgnv6ScaleoutAddressMappingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6ScaleoutAddressMappingOperOper(ret edpt.DataCgnv6ScaleoutAddressMappingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"user_group":   ret.DtCgnv6ScaleoutAddressMappingOper.Oper.UserGroup,
			"active_node":  ret.DtCgnv6ScaleoutAddressMappingOper.Oper.ActiveNode,
			"standby_node": ret.DtCgnv6ScaleoutAddressMappingOper.Oper.StandbyNode,
			"ip_list":      setSliceCgnv6ScaleoutAddressMappingOperOperIpList(ret.DtCgnv6ScaleoutAddressMappingOper.Oper.IpList),
			"ip":           ret.DtCgnv6ScaleoutAddressMappingOper.Oper.Ip,
			"ipv6":         ret.DtCgnv6ScaleoutAddressMappingOper.Oper.Ipv6,
			"nat_ip":       ret.DtCgnv6ScaleoutAddressMappingOper.Oper.NatIp,
			"nat_pool":     ret.DtCgnv6ScaleoutAddressMappingOper.Oper.NatPool,
		},
	}
}

func setSliceCgnv6ScaleoutAddressMappingOperOperIpList(d []edpt.Cgnv6ScaleoutAddressMappingOperOperIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["start_nat_ip"] = item.StartNatIp
		in["end_nat_ip"] = item.EndNatIp
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6ScaleoutAddressMappingOperOper(d []interface{}) edpt.Cgnv6ScaleoutAddressMappingOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6ScaleoutAddressMappingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UserGroup = in["user_group"].(int)
		ret.ActiveNode = in["active_node"].(int)
		ret.StandbyNode = in["standby_node"].(int)
		ret.IpList = getSliceCgnv6ScaleoutAddressMappingOperOperIpList(in["ip_list"].([]interface{}))
		ret.Ip = in["ip"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.NatIp = in["nat_ip"].(string)
		ret.NatPool = in["nat_pool"].(string)
	}
	return ret
}

func getSliceCgnv6ScaleoutAddressMappingOperOperIpList(d []interface{}) []edpt.Cgnv6ScaleoutAddressMappingOperOperIpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ScaleoutAddressMappingOperOperIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ScaleoutAddressMappingOperOperIpList
		oi.StartNatIp = in["start_nat_ip"].(string)
		oi.EndNatIp = in["end_nat_ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6ScaleoutAddressMappingOper(d *schema.ResourceData) edpt.Cgnv6ScaleoutAddressMappingOper {
	var ret edpt.Cgnv6ScaleoutAddressMappingOper

	ret.Oper = getObjectCgnv6ScaleoutAddressMappingOperOper(d.Get("oper").([]interface{}))
	return ret
}
