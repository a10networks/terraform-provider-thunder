package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpRstCloseImmediate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_tcp_rst_close_immediate`: Configure TCP RST behavior\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTcpRstCloseImmediateCreate,
		UpdateContext: resourceFwTcpRstCloseImmediateUpdate,
		ReadContext:   resourceFwTcpRstCloseImmediateRead,
		DeleteContext: resourceFwTcpRstCloseImmediateDelete,

		Schema: map[string]*schema.Schema{
			"status": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable TCP RST close immediate (default); 'disable': Disable TCP RST close immediate;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwTcpRstCloseImmediateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpRstCloseImmediateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpRstCloseImmediate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpRstCloseImmediateRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTcpRstCloseImmediateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpRstCloseImmediateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpRstCloseImmediate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpRstCloseImmediateRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTcpRstCloseImmediateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpRstCloseImmediateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpRstCloseImmediate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTcpRstCloseImmediateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpRstCloseImmediateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpRstCloseImmediate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwTcpRstCloseImmediate(d *schema.ResourceData) edpt.FwTcpRstCloseImmediate {
	var ret edpt.FwTcpRstCloseImmediate
	ret.Inst.Status = d.Get("status").(string)
	//omit uuid
	return ret
}
