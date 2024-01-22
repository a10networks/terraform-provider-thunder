package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugOspfAll() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ospf_all`: Enable all debugging\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugOspfAllCreate,
		UpdateContext: resourceDebugOspfAllUpdate,
		ReadContext:   resourceDebugOspfAllRead,
		DeleteContext: resourceDebugOspfAllDelete,

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
func resourceDebugOspfAllCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfAllCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfAll(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfAllRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugOspfAllUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfAllUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfAll(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfAllRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugOspfAllDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfAllDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfAll(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugOspfAllRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfAllRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfAll(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugOspfAll(d *schema.ResourceData) edpt.DebugOspfAll {
	var ret edpt.DebugOspfAll
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
