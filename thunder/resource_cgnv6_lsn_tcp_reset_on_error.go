package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnTcpResetOnError() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_tcp_reset_on_error`: Send TCP resets for invalid sessions (Default: enabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnTcpResetOnErrorCreate,
		UpdateContext: resourceCgnv6LsnTcpResetOnErrorUpdate,
		ReadContext:   resourceCgnv6LsnTcpResetOnErrorRead,
		DeleteContext: resourceCgnv6LsnTcpResetOnErrorDelete,

		Schema: map[string]*schema.Schema{
			"outbound": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable send TCP reset on error;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnTcpResetOnErrorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnTcpResetOnErrorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnTcpResetOnError(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnTcpResetOnErrorRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnTcpResetOnErrorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnTcpResetOnErrorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnTcpResetOnError(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnTcpResetOnErrorRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnTcpResetOnErrorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnTcpResetOnErrorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnTcpResetOnError(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnTcpResetOnErrorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnTcpResetOnErrorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnTcpResetOnError(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LsnTcpResetOnError(d *schema.ResourceData) edpt.Cgnv6LsnTcpResetOnError {
	var ret edpt.Cgnv6LsnTcpResetOnError
	ret.Inst.Outbound = d.Get("outbound").(string)
	//omit uuid
	return ret
}
