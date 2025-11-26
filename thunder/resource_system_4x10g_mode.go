package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystem4x10gMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_4x10g_mode`: To enable/disable 40G port to split into 4x10g ports\n\n__PLACEHOLDER__",
		CreateContext: resourceSystem4x10gModeCreate,
		UpdateContext: resourceSystem4x10gModeUpdate,
		ReadContext:   resourceSystem4x10gModeRead,
		DeleteContext: resourceSystem4x10gModeDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable 40G port to split into 4x10g ports",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystem4x10gModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystem4x10gModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystem4x10gMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystem4x10gModeRead(ctx, d, meta)
	}
	return diags
}

func resourceSystem4x10gModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystem4x10gModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystem4x10gMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystem4x10gModeRead(ctx, d, meta)
	}
	return diags
}
func resourceSystem4x10gModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystem4x10gModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystem4x10gMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystem4x10gModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystem4x10gModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystem4x10gMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystem4x10gMode(d *schema.ResourceData) edpt.System4x10gMode {
	var ret edpt.System4x10gMode
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
