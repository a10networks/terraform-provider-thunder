package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutUserGroupAssignmentTemplateV6Assignment() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_user_group_assignment_template_v6_assignment`: Configure scaleout user group assignment\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutUserGroupAssignmentTemplateV6AssignmentCreate,
		UpdateContext: resourceScaleoutUserGroupAssignmentTemplateV6AssignmentUpdate,
		ReadContext:   resourceScaleoutUserGroupAssignmentTemplateV6AssignmentRead,
		DeleteContext: resourceScaleoutUserGroupAssignmentTemplateV6AssignmentDelete,

		Schema: map[string]*schema.Schema{
			"assignment_prefix_length": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "User group assignment prefix length, default is 128",
			},
			"ipv6_prefix": {
				Type: schema.TypeString, Required: true, Description: "IPv6 prefix",
			},
			"service_config_template": {
				Type: schema.TypeString, Optional: true, Description: "Configure a scaleout service config template to use",
			},
			"user_group_range_end": {
				Type: schema.TypeInt, Optional: true, Description: "User group range end",
			},
			"user_group_range_start": {
				Type: schema.TypeInt, Optional: true, Description: "User group range start",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"user_group_assignment_template_name": {
				Type: schema.TypeString, Required: true, Description: "User_group_assignment_template_name",
			},
		},
	}
}
func resourceScaleoutUserGroupAssignmentTemplateV6AssignmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateV6AssignmentCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplateV6Assignment(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutUserGroupAssignmentTemplateV6AssignmentRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutUserGroupAssignmentTemplateV6AssignmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateV6AssignmentUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplateV6Assignment(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutUserGroupAssignmentTemplateV6AssignmentRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutUserGroupAssignmentTemplateV6AssignmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateV6AssignmentDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplateV6Assignment(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutUserGroupAssignmentTemplateV6AssignmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateV6AssignmentRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplateV6Assignment(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutUserGroupAssignmentTemplateV6Assignment(d *schema.ResourceData) edpt.ScaleoutUserGroupAssignmentTemplateV6Assignment {
	var ret edpt.ScaleoutUserGroupAssignmentTemplateV6Assignment
	ret.Inst.AssignmentPrefixLength = d.Get("assignment_prefix_length").(int)
	ret.Inst.Ipv6Prefix = d.Get("ipv6_prefix").(string)
	ret.Inst.ServiceConfigTemplate = d.Get("service_config_template").(string)
	ret.Inst.UserGroupRangeEnd = d.Get("user_group_range_end").(int)
	ret.Inst.UserGroupRangeStart = d.Get("user_group_range_start").(int)
	//omit uuid
	ret.Inst.User_group_assignment_template_name = d.Get("user_group_assignment_template_name").(string)
	return ret
}
