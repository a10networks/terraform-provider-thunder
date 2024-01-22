package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwHelperSessions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_helper_sessions`: Configure firewall helper-session (TAC use only)\n\n__PLACEHOLDER__",
		CreateContext: resourceFwHelperSessionsCreate,
		UpdateContext: resourceFwHelperSessionsUpdate,
		ReadContext:   resourceFwHelperSessionsRead,
		DeleteContext: resourceFwHelperSessionsDelete,

		Schema: map[string]*schema.Schema{
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "helper-sessions idle-timeout time (Idle-timeout in minutes (default: 1 minute))",
			},
			"limit": {
				Type: schema.TypeInt, Optional: true, Description: "Limit number of helper-sessions (Limit helper-sessions number)",
			},
			"mode": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable helper-sessions;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwHelperSessionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwHelperSessionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwHelperSessions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwHelperSessionsRead(ctx, d, meta)
	}
	return diags
}

func resourceFwHelperSessionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwHelperSessionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwHelperSessions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwHelperSessionsRead(ctx, d, meta)
	}
	return diags
}
func resourceFwHelperSessionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwHelperSessionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwHelperSessions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwHelperSessionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwHelperSessionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwHelperSessions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwHelperSessions(d *schema.ResourceData) edpt.FwHelperSessions {
	var ret edpt.FwHelperSessions
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.Limit = d.Get("limit").(int)
	ret.Inst.Mode = d.Get("mode").(string)
	//omit uuid
	return ret
}
