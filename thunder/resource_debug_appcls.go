package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugAppcls() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_appcls`: Debug application classification\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugAppclsCreate,
		UpdateContext: resourceDebugAppclsUpdate,
		ReadContext:   resourceDebugAppclsRead,
		DeleteContext: resourceDebugAppclsDelete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeString, Optional: true, Description: "'1': Display critical errors only; '2': Display all unusual results; '3': Display all appcls engine information;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugAppclsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugAppclsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugAppcls(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugAppclsRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugAppclsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugAppclsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugAppcls(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugAppclsRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugAppclsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugAppclsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugAppcls(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugAppclsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugAppclsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugAppcls(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugAppcls(d *schema.ResourceData) edpt.DebugAppcls {
	var ret edpt.DebugAppcls
	ret.Inst.Level = d.Get("level").(string)
	//omit uuid
	return ret
}
