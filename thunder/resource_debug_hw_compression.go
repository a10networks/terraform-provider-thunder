package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugHwCompression() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_hw_compression`: Debug Hardware Compression\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugHwCompressionCreate,
		UpdateContext: resourceDebugHwCompressionUpdate,
		ReadContext:   resourceDebugHwCompressionRead,
		DeleteContext: resourceDebugHwCompressionDelete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-4)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugHwCompressionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHwCompressionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHwCompression(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHwCompressionRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugHwCompressionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHwCompressionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHwCompression(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHwCompressionRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugHwCompressionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHwCompressionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHwCompression(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugHwCompressionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHwCompressionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHwCompression(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugHwCompression(d *schema.ResourceData) edpt.DebugHwCompression {
	var ret edpt.DebugHwCompression
	ret.Inst.Level = d.Get("level").(int)
	//omit uuid
	return ret
}
