package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_scaleout_address_mapping_inside_user_or_nat_ip_oper`: Operational Status for the object inside-user-or-nat-ip\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6ScaleoutAddressMappingInsideUserOrNatIpOperRead,

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
						"service_template": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"application_type": {
							Type: schema.TypeString, Optional: true, Description: "",
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
						"application": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6ScaleoutAddressMappingInsideUserOrNatIpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ScaleoutAddressMappingInsideUserOrNatIpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper := setObjectCgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper(res)
		d.Set("oper", Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper(ret edpt.DataCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"user_group":       ret.DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper.Oper.UserGroup,
			"active_node":      ret.DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper.Oper.ActiveNode,
			"standby_node":     ret.DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper.Oper.StandbyNode,
			"service_template": ret.DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper.Oper.ServiceTemplate,
			"application_type": ret.DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper.Oper.Application_type,
			"ip":               ret.DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper.Oper.Ip,
			"ipv6":             ret.DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper.Oper.Ipv6,
			"nat_ip":           ret.DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper.Oper.NatIp,
			"application":      ret.DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper.Oper.Application,
		},
	}
}

func getObjectCgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper(d []interface{}) edpt.Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UserGroup = in["user_group"].(int)
		ret.ActiveNode = in["active_node"].(int)
		ret.StandbyNode = in["standby_node"].(int)
		ret.ServiceTemplate = in["service_template"].(string)
		ret.Application_type = in["application_type"].(string)
		ret.Ip = in["ip"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.NatIp = in["nat_ip"].(string)
		ret.Application = in["application"].(string)
	}
	return ret
}

func dataToEndpointCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper(d *schema.ResourceData) edpt.Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOper {
	var ret edpt.Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOper

	ret.Oper = getObjectCgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper(d.Get("oper").([]interface{}))
	return ret
}
