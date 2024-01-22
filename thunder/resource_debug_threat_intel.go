package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugThreatIntel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_threat_intel`: Debug threat intel module\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugThreatIntelCreate,
		UpdateContext: resourceDebugThreatIntelUpdate,
		ReadContext:   resourceDebugThreatIntelRead,
		DeleteContext: resourceDebugThreatIntelDelete,

		Schema: map[string]*schema.Schema{
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugThreatIntelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugThreatIntelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugThreatIntel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugThreatIntelRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugThreatIntelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugThreatIntelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugThreatIntel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugThreatIntelRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugThreatIntelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugThreatIntelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugThreatIntel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugThreatIntelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugThreatIntelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugThreatIntel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugThreatIntel(d *schema.ResourceData) edpt.DebugThreatIntel {
	var ret edpt.DebugThreatIntel
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
