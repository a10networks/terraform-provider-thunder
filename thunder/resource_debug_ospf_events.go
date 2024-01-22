package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugOspfEvents() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ospf_events`: OSPFv3 event\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugOspfEventsCreate,
		UpdateContext: resourceDebugOspfEventsUpdate,
		ReadContext:   resourceDebugOspfEventsRead,
		DeleteContext: resourceDebugOspfEventsDelete,

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
func resourceDebugOspfEventsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfEventsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfEvents(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfEventsRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugOspfEventsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfEventsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfEvents(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfEventsRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugOspfEventsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfEventsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfEvents(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugOspfEventsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfEventsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfEvents(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugOspfEvents(d *schema.ResourceData) edpt.DebugOspfEvents {
	var ret edpt.DebugOspfEvents
	ret.Inst.Abr = d.Get("abr").(int)
	ret.Inst.Asbr = d.Get("asbr").(int)
	ret.Inst.Os = d.Get("os").(int)
	ret.Inst.Router = d.Get("router").(int)
	//omit uuid
	ret.Inst.Vlink = d.Get("vlink").(int)
	return ret
}
