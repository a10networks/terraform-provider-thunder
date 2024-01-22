package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIpv6OspfNsm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ipv6_ospf_nsm`: OSPFv3 NSM information\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIpv6OspfNsmCreate,
		UpdateContext: resourceDebugIpv6OspfNsmUpdate,
		ReadContext:   resourceDebugIpv6OspfNsmRead,
		DeleteContext: resourceDebugIpv6OspfNsmDelete,

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
func resourceDebugIpv6OspfNsmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfNsmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfNsm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfNsmRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIpv6OspfNsmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfNsmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfNsm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfNsmRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIpv6OspfNsmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfNsmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfNsm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIpv6OspfNsmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfNsmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfNsm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugIpv6OspfNsm(d *schema.ResourceData) edpt.DebugIpv6OspfNsm {
	var ret edpt.DebugIpv6OspfNsm
	ret.Inst.Interface = d.Get("interface").(int)
	ret.Inst.Redistribute = d.Get("redistribute").(int)
	//omit uuid
	return ret
}
