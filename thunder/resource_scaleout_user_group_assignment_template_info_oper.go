package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutUserGroupAssignmentTemplateInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_user_group_assignment_template_info_oper`: Operational Status for the object template-info\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutUserGroupAssignmentTemplateInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assignment_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_prefix": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"prefixes_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip_count_per_prefix": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"prefixes_per_user_group": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"user_group_range_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"user_group_range_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceScaleoutUserGroupAssignmentTemplateInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplateInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutUserGroupAssignmentTemplateInfoOperOper := setObjectScaleoutUserGroupAssignmentTemplateInfoOperOper(res)
		d.Set("oper", ScaleoutUserGroupAssignmentTemplateInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutUserGroupAssignmentTemplateInfoOperOper(ret edpt.DataScaleoutUserGroupAssignmentTemplateInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"assignment_list": setSliceScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList(ret.DtScaleoutUserGroupAssignmentTemplateInfoOper.Oper.AssignmentList),
			"name":            ret.DtScaleoutUserGroupAssignmentTemplateInfoOper.Oper.Name,
		},
	}
}

func setSliceScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList(d []edpt.ScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_prefix"] = item.AddressPrefix
		in["prefixes_count"] = item.PrefixesCount
		in["ip_count_per_prefix"] = item.IpCountPerPrefix
		in["prefixes_per_user_group"] = item.PrefixesPerUserGroup
		in["user_group_range_start"] = item.UserGroupRangeStart
		in["user_group_range_end"] = item.UserGroupRangeEnd
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutUserGroupAssignmentTemplateInfoOperOper(d []interface{}) edpt.ScaleoutUserGroupAssignmentTemplateInfoOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutUserGroupAssignmentTemplateInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AssignmentList = getSliceScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList(in["assignment_list"].([]interface{}))
		ret.Name = in["name"].(string)
	}
	return ret
}

func getSliceScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList(d []interface{}) []edpt.ScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList
		oi.AddressPrefix = in["address_prefix"].(string)
		oi.PrefixesCount = in["prefixes_count"].(int)
		oi.IpCountPerPrefix = in["ip_count_per_prefix"].(int)
		oi.PrefixesPerUserGroup = in["prefixes_per_user_group"].(int)
		oi.UserGroupRangeStart = in["user_group_range_start"].(int)
		oi.UserGroupRangeEnd = in["user_group_range_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutUserGroupAssignmentTemplateInfoOper(d *schema.ResourceData) edpt.ScaleoutUserGroupAssignmentTemplateInfoOper {
	var ret edpt.ScaleoutUserGroupAssignmentTemplateInfoOper

	ret.Oper = getObjectScaleoutUserGroupAssignmentTemplateInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
