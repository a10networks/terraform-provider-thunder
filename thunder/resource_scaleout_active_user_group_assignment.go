package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutActiveUserGroupAssignment() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_active_user_group_assignment`: Active scaleout user group assignment template\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutActiveUserGroupAssignmentCreate,
		UpdateContext: resourceScaleoutActiveUserGroupAssignmentUpdate,
		ReadContext:   resourceScaleoutActiveUserGroupAssignmentRead,
		DeleteContext: resourceScaleoutActiveUserGroupAssignmentDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Scaleout user group assignment template",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceScaleoutActiveUserGroupAssignmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutActiveUserGroupAssignmentCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutActiveUserGroupAssignment(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutActiveUserGroupAssignmentRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutActiveUserGroupAssignmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutActiveUserGroupAssignmentUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutActiveUserGroupAssignment(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutActiveUserGroupAssignmentRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutActiveUserGroupAssignmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutActiveUserGroupAssignmentDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutActiveUserGroupAssignment(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutActiveUserGroupAssignmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutActiveUserGroupAssignmentRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutActiveUserGroupAssignment(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutActiveUserGroupAssignment(d *schema.ResourceData) edpt.ScaleoutActiveUserGroupAssignment {
	var ret edpt.ScaleoutActiveUserGroupAssignment
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	return ret
}
