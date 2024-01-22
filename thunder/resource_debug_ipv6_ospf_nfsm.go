package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIpv6OspfNfsm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ipv6_ospf_nfsm`: OSPFv3 Neighbor State Machine\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIpv6OspfNfsmCreate,
		UpdateContext: resourceDebugIpv6OspfNfsmUpdate,
		ReadContext:   resourceDebugIpv6OspfNfsmRead,
		DeleteContext: resourceDebugIpv6OspfNfsmDelete,

		Schema: map[string]*schema.Schema{
			"events": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NFSM Event Information",
			},
			"status": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NFSM Status Information",
			},
			"timers": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NFSM Timer Information",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugIpv6OspfNfsmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfNfsmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfNfsm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfNfsmRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIpv6OspfNfsmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfNfsmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfNfsm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfNfsmRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIpv6OspfNfsmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfNfsmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfNfsm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIpv6OspfNfsmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfNfsmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfNfsm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugIpv6OspfNfsm(d *schema.ResourceData) edpt.DebugIpv6OspfNfsm {
	var ret edpt.DebugIpv6OspfNfsm
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Status = d.Get("status").(int)
	ret.Inst.Timers = d.Get("timers").(int)
	//omit uuid
	return ret
}
