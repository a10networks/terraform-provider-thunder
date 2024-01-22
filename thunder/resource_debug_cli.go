package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugCli() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_cli`: CLI module parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugCliCreate,
		UpdateContext: resourceDebugCliUpdate,
		ReadContext:   resourceDebugCliRead,
		DeleteContext: resourceDebugCliDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "CLI all debug switch",
			},
			"io": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "CLI debug input-output switch",
			},
			"parser": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "CLI debug parser switch",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugCliCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCliCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCli(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugCliRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugCliUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCliUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCli(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugCliRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugCliDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCliDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCli(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugCliRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCliRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCli(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugCli(d *schema.ResourceData) edpt.DebugCli {
	var ret edpt.DebugCli
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Io = d.Get("io").(int)
	ret.Inst.Parser = d.Get("parser").(int)
	//omit uuid
	return ret
}
