package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsAction() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_action`: action for aVCS\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsActionCreate,
		UpdateContext: resourceVcsActionUpdate,
		ReadContext:   resourceVcsActionRead,
		DeleteContext: resourceVcsActionDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'disable': disable VCS; 'enable': enable VCS;",
			},
			"database_distribution": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable database cluster distribution; 'disable': disable database cluster distribution;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVcsActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsActionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsAction(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsActionRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsActionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsAction(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsActionRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsActionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsAction(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsActionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsAction(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVcsAction(d *schema.ResourceData) edpt.VcsAction {
	var ret edpt.VcsAction
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DatabaseDistribution = d.Get("database_distribution").(string)
	//omit uuid
	return ret
}
