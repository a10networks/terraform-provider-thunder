package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugOspfNsm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ospf_nsm`: OSPFv3 NSM information\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugOspfNsmCreate,
		UpdateContext: resourceDebugOspfNsmUpdate,
		ReadContext:   resourceDebugOspfNsmRead,
		DeleteContext: resourceDebugOspfNsmDelete,

		Schema: map[string]*schema.Schema{
			"interface": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NSM interface",
			},
			"redistribute": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NSM redistribute",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugOspfNsmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfNsmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfNsm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfNsmRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugOspfNsmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfNsmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfNsm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfNsmRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugOspfNsmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfNsmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfNsm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugOspfNsmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfNsmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfNsm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugOspfNsm(d *schema.ResourceData) edpt.DebugOspfNsm {
	var ret edpt.DebugOspfNsm
	ret.Inst.Interface = d.Get("interface").(int)
	ret.Inst.Redistribute = d.Get("redistribute").(int)
	//omit uuid
	return ret
}
