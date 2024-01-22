package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFailSafeDisableFailsafe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fail_safe_disable_failsafe`: Disable HA failover when resource depletion is detected\n\n__PLACEHOLDER__",
		CreateContext: resourceFailSafeDisableFailsafeCreate,
		UpdateContext: resourceFailSafeDisableFailsafeUpdate,
		ReadContext:   resourceFailSafeDisableFailsafeRead,
		DeleteContext: resourceFailSafeDisableFailsafeDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "all", Description: "'all': Disable All; 'io-buffer': Disable I/O Buffer; 'session-memory': Disable Session Memory; 'system-memory': Disable System Memory;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFailSafeDisableFailsafeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeDisableFailsafeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafeDisableFailsafe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFailSafeDisableFailsafeRead(ctx, d, meta)
	}
	return diags
}

func resourceFailSafeDisableFailsafeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeDisableFailsafeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafeDisableFailsafe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFailSafeDisableFailsafeRead(ctx, d, meta)
	}
	return diags
}
func resourceFailSafeDisableFailsafeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeDisableFailsafeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafeDisableFailsafe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFailSafeDisableFailsafeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeDisableFailsafeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafeDisableFailsafe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFailSafeDisableFailsafe(d *schema.ResourceData) edpt.FailSafeDisableFailsafe {
	var ret edpt.FailSafeDisableFailsafe
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
