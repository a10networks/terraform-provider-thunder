package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutUserGroupAssignmentAddressOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_user_group_assignment_address_oper`: Operational Status for the object address\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutUserGroupAssignmentAddressOperRead,

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
					},
				},
			},
		},
	}
}

func resourceScaleoutUserGroupAssignmentAddressOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentAddressOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentAddressOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutUserGroupAssignmentAddressOperOper := setObjectScaleoutUserGroupAssignmentAddressOperOper(res)
		d.Set("oper", ScaleoutUserGroupAssignmentAddressOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutUserGroupAssignmentAddressOperOper(ret edpt.DataScaleoutUserGroupAssignmentAddressOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"user_group":       ret.DtScaleoutUserGroupAssignmentAddressOper.Oper.UserGroup,
			"active_node":      ret.DtScaleoutUserGroupAssignmentAddressOper.Oper.ActiveNode,
			"standby_node":     ret.DtScaleoutUserGroupAssignmentAddressOper.Oper.StandbyNode,
			"service_template": ret.DtScaleoutUserGroupAssignmentAddressOper.Oper.ServiceTemplate,
			"application_type": ret.DtScaleoutUserGroupAssignmentAddressOper.Oper.Application_type,
			"ip":               ret.DtScaleoutUserGroupAssignmentAddressOper.Oper.Ip,
			"ipv6":             ret.DtScaleoutUserGroupAssignmentAddressOper.Oper.Ipv6,
		},
	}
}

func getObjectScaleoutUserGroupAssignmentAddressOperOper(d []interface{}) edpt.ScaleoutUserGroupAssignmentAddressOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutUserGroupAssignmentAddressOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UserGroup = in["user_group"].(int)
		ret.ActiveNode = in["active_node"].(int)
		ret.StandbyNode = in["standby_node"].(int)
		ret.ServiceTemplate = in["service_template"].(string)
		ret.Application_type = in["application_type"].(string)
		ret.Ip = in["ip"].(string)
		ret.Ipv6 = in["ipv6"].(string)
	}
	return ret
}

func dataToEndpointScaleoutUserGroupAssignmentAddressOper(d *schema.ResourceData) edpt.ScaleoutUserGroupAssignmentAddressOper {
	var ret edpt.ScaleoutUserGroupAssignmentAddressOper

	ret.Oper = getObjectScaleoutUserGroupAssignmentAddressOperOper(d.Get("oper").([]interface{}))
	return ret
}
