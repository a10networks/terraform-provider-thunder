package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutUserGroupAssignmentTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_user_group_assignment_template`: Configure scaleout user group assignment templates\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutUserGroupAssignmentTemplateCreate,
		UpdateContext: resourceScaleoutUserGroupAssignmentTemplateUpdate,
		ReadContext:   resourceScaleoutUserGroupAssignmentTemplateRead,
		DeleteContext: resourceScaleoutUserGroupAssignmentTemplateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Scaleout User Group Assignment Template Name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v4_assignment_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_prefix": {
							Type: schema.TypeString, Required: true, Description: "IPv4 prefix",
						},
						"assignment_prefix_length": {
							Type: schema.TypeInt, Optional: true, Default: 32, Description: "User group assignment prefix length, default is 32",
						},
						"user_group_range_start": {
							Type: schema.TypeInt, Optional: true, Description: "User group range start",
						},
						"user_group_range_end": {
							Type: schema.TypeInt, Optional: true, Description: "User group range end",
						},
						"service_config_template": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"v6_assignment_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_prefix": {
							Type: schema.TypeString, Required: true, Description: "IPv6 prefix",
						},
						"assignment_prefix_length": {
							Type: schema.TypeInt, Optional: true, Default: 128, Description: "User group assignment prefix length, default is 128",
						},
						"user_group_range_start": {
							Type: schema.TypeInt, Optional: true, Description: "User group range start",
						},
						"user_group_range_end": {
							Type: schema.TypeInt, Optional: true, Description: "User group range end",
						},
						"service_config_template": {
							Type: schema.TypeString, Optional: true, Description: "Configure a scaleout service config template to use",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
		},
	}
}
func resourceScaleoutUserGroupAssignmentTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutUserGroupAssignmentTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutUserGroupAssignmentTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutUserGroupAssignmentTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutUserGroupAssignmentTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutUserGroupAssignmentTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutUserGroupAssignmentTemplateV4AssignmentList(d []interface{}) []edpt.ScaleoutUserGroupAssignmentTemplateV4AssignmentList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutUserGroupAssignmentTemplateV4AssignmentList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutUserGroupAssignmentTemplateV4AssignmentList
		oi.Ipv4Prefix = in["ipv4_prefix"].(string)
		oi.AssignmentPrefixLength = in["assignment_prefix_length"].(int)
		oi.UserGroupRangeStart = in["user_group_range_start"].(int)
		oi.UserGroupRangeEnd = in["user_group_range_end"].(int)
		oi.ServiceConfigTemplate = in["service_config_template"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutUserGroupAssignmentTemplateV6AssignmentList(d []interface{}) []edpt.ScaleoutUserGroupAssignmentTemplateV6AssignmentList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutUserGroupAssignmentTemplateV6AssignmentList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutUserGroupAssignmentTemplateV6AssignmentList
		oi.Ipv6Prefix = in["ipv6_prefix"].(string)
		oi.AssignmentPrefixLength = in["assignment_prefix_length"].(int)
		oi.UserGroupRangeStart = in["user_group_range_start"].(int)
		oi.UserGroupRangeEnd = in["user_group_range_end"].(int)
		oi.ServiceConfigTemplate = in["service_config_template"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutUserGroupAssignmentTemplate(d *schema.ResourceData) edpt.ScaleoutUserGroupAssignmentTemplate {
	var ret edpt.ScaleoutUserGroupAssignmentTemplate
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.V4AssignmentList = getSliceScaleoutUserGroupAssignmentTemplateV4AssignmentList(d.Get("v4_assignment_list").([]interface{}))
	ret.Inst.V6AssignmentList = getSliceScaleoutUserGroupAssignmentTemplateV6AssignmentList(d.Get("v6_assignment_list").([]interface{}))
	return ret
}
