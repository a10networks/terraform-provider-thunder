package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugCgn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_cgn`: CGN packet debugging\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugCgnCreate,
		UpdateContext: resourceDebugCgnUpdate,
		ReadContext:   resourceDebugCgnRead,
		DeleteContext: resourceDebugCgnDelete,

		Schema: map[string]*schema.Schema{
			"drops": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for packet drops in CGN",
			},
			"error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for errors in CGN",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugCgnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCgnCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCgn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugCgnRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugCgnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCgnUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCgn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugCgnRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugCgnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCgnDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCgn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugCgnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCgnRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCgn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugCgn(d *schema.ResourceData) edpt.DebugCgn {
	var ret edpt.DebugCgn
	ret.Inst.Drops = d.Get("drops").(int)
	ret.Inst.Error = d.Get("error").(int)
	//omit uuid
	return ret
}
