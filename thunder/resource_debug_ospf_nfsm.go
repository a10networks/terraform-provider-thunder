package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugOspfNfsm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ospf_nfsm`: OSPFv3 Neighbor State Machine\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugOspfNfsmCreate,
		UpdateContext: resourceDebugOspfNfsmUpdate,
		ReadContext:   resourceDebugOspfNfsmRead,
		DeleteContext: resourceDebugOspfNfsmDelete,

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
func resourceDebugOspfNfsmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfNfsmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfNfsm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfNfsmRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugOspfNfsmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfNfsmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfNfsm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfNfsmRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugOspfNfsmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfNfsmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfNfsm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugOspfNfsmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfNfsmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfNfsm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugOspfNfsm(d *schema.ResourceData) edpt.DebugOspfNfsm {
	var ret edpt.DebugOspfNfsm
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Status = d.Get("status").(int)
	ret.Inst.Timers = d.Get("timers").(int)
	//omit uuid
	return ret
}
