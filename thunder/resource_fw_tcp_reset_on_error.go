package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpResetOnError() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_tcp_reset_on_error`: Send TCP resets for invalid sessions (Default: Disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTcpResetOnErrorCreate,
		UpdateContext: resourceFwTcpResetOnErrorUpdate,
		ReadContext:   resourceFwTcpResetOnErrorRead,
		DeleteContext: resourceFwTcpResetOnErrorDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable send TCP reset on error",
			},
			"outbound": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable send TCP reset on error;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwTcpResetOnErrorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpResetOnErrorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpResetOnError(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpResetOnErrorRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTcpResetOnErrorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpResetOnErrorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpResetOnError(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpResetOnErrorRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTcpResetOnErrorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpResetOnErrorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpResetOnError(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTcpResetOnErrorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpResetOnErrorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpResetOnError(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwTcpResetOnError(d *schema.ResourceData) edpt.FwTcpResetOnError {
	var ret edpt.FwTcpResetOnError
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.Outbound = d.Get("outbound").(string)
	//omit uuid
	return ret
}
