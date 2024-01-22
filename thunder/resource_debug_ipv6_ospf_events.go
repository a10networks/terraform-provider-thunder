package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIpv6OspfEvents() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ipv6_ospf_events`: OSPFv3 event\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIpv6OspfEventsCreate,
		UpdateContext: resourceDebugIpv6OspfEventsUpdate,
		ReadContext:   resourceDebugIpv6OspfEventsRead,
		DeleteContext: resourceDebugIpv6OspfEventsDelete,

		Schema: map[string]*schema.Schema{
			"abr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPF ABR events",
			},
			"asbr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPF ASBR events",
			},
			"os": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "OS events",
			},
			"router": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Other router events",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlink": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Virtual-Link event",
			},
		},
	}
}
func resourceDebugIpv6OspfEventsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfEventsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfEvents(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfEventsRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIpv6OspfEventsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfEventsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfEvents(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfEventsRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIpv6OspfEventsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfEventsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfEvents(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIpv6OspfEventsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfEventsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfEvents(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugIpv6OspfEvents(d *schema.ResourceData) edpt.DebugIpv6OspfEvents {
	var ret edpt.DebugIpv6OspfEvents
	ret.Inst.Abr = d.Get("abr").(int)
	ret.Inst.Asbr = d.Get("asbr").(int)
	ret.Inst.Os = d.Get("os").(int)
	ret.Inst.Router = d.Get("router").(int)
	//omit uuid
	ret.Inst.Vlink = d.Get("vlink").(int)
	return ret
}
