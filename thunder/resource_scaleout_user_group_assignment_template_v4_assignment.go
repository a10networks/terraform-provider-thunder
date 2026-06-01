package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutUserGroupAssignmentTemplateV4Assignment() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_user_group_assignment_template_v4_assignment`: Configure scaleout user group assignment\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutUserGroupAssignmentTemplateV4AssignmentCreate,
		UpdateContext: resourceScaleoutUserGroupAssignmentTemplateV4AssignmentUpdate,
		ReadContext:   resourceScaleoutUserGroupAssignmentTemplateV4AssignmentRead,
		DeleteContext: resourceScaleoutUserGroupAssignmentTemplateV4AssignmentDelete,

		Schema: map[string]*schema.Schema{
			"assignment_prefix_auto": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Automatically break down the prefix so that each user-group is assigned one and only one subnet",
			},
			"assignment_prefix_length": {
				Type: schema.TypeInt, Optional: true, Default: 32, Description: "User group assignment prefix length, default is 32",
			},
			"ipv4_prefix": {
				Type: schema.TypeString, Required: true, Description: "IPv4 prefix",
			},
			"private_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the assignment as private, and no BGP route will be advertised for it",
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
func resourceScaleoutUserGroupAssignmentTemplateV4AssignmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateV4AssignmentCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplateV4Assignment(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutUserGroupAssignmentTemplateV4AssignmentRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutUserGroupAssignmentTemplateV4AssignmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateV4AssignmentUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplateV4Assignment(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutUserGroupAssignmentTemplateV4AssignmentRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutUserGroupAssignmentTemplateV4AssignmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateV4AssignmentDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplateV4Assignment(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutUserGroupAssignmentTemplateV4AssignmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutUserGroupAssignmentTemplateV4AssignmentRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutUserGroupAssignmentTemplateV4Assignment(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutUserGroupAssignmentTemplateV4Assignment(d *schema.ResourceData) edpt.ScaleoutUserGroupAssignmentTemplateV4Assignment {
	var ret edpt.ScaleoutUserGroupAssignmentTemplateV4Assignment
	ret.Inst.AssignmentPrefixAuto = d.Get("assignment_prefix_auto").(int)
	ret.Inst.AssignmentPrefixLength = d.Get("assignment_prefix_length").(int)
	ret.Inst.Ipv4Prefix = d.Get("ipv4_prefix").(string)
	ret.Inst.PrivateIp = d.Get("private_ip").(int)
	ret.Inst.ServiceConfigTemplate = d.Get("service_config_template").(string)
	ret.Inst.UserGroupRangeEnd = d.Get("user_group_range_end").(int)
	ret.Inst.UserGroupRangeStart = d.Get("user_group_range_start").(int)
	//omit uuid
	ret.Inst.User_group_assignment_template_name = d.Get("user_group_assignment_template_name").(string)
	return ret
}
