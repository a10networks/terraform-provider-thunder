package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIpv6OspfIfsm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ipv6_ospf_ifsm`: OSPFv3 Interface State Machine\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIpv6OspfIfsmCreate,
		UpdateContext: resourceDebugIpv6OspfIfsmUpdate,
		ReadContext:   resourceDebugIpv6OspfIfsmRead,
		DeleteContext: resourceDebugIpv6OspfIfsmDelete,

		Schema: map[string]*schema.Schema{
			"events": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IFSM Event Information",
			},
			"status": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IFSM Status Information",
			},
			"timers": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IFSM Timer Information",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugIpv6OspfIfsmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfIfsmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfIfsm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfIfsmRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIpv6OspfIfsmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfIfsmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfIfsm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfIfsmRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIpv6OspfIfsmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfIfsmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfIfsm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIpv6OspfIfsmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfIfsmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfIfsm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugIpv6OspfIfsm(d *schema.ResourceData) edpt.DebugIpv6OspfIfsm {
	var ret edpt.DebugIpv6OspfIfsm
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Status = d.Get("status").(int)
	ret.Inst.Timers = d.Get("timers").(int)
	//omit uuid
	return ret
}
