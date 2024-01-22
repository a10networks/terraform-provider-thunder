package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugFw() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_fw`: Debug Firewall\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugFwCreate,
		UpdateContext: resourceDebugFwUpdate,
		ReadContext:   resourceDebugFwRead,
		DeleteContext: resourceDebugFwDelete,

		Schema: map[string]*schema.Schema{
			"ddos": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "FW DDoS detection/mitigation logs",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugFwCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugFwCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugFw(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugFwRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugFwUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugFwUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugFw(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugFwRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugFwDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugFwDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugFw(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugFwRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugFwRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugFw(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugFw(d *schema.ResourceData) edpt.DebugFw {
	var ret edpt.DebugFw
	ret.Inst.Ddos = d.Get("ddos").(int)
	//omit uuid
	return ret
}
