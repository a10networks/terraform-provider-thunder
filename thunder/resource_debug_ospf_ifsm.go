package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugOspfIfsm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ospf_ifsm`: OSPFv3 Interface State Machine\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugOspfIfsmCreate,
		UpdateContext: resourceDebugOspfIfsmUpdate,
		ReadContext:   resourceDebugOspfIfsmRead,
		DeleteContext: resourceDebugOspfIfsmDelete,

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
func resourceDebugOspfIfsmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfIfsmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfIfsm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfIfsmRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugOspfIfsmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfIfsmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfIfsm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfIfsmRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugOspfIfsmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfIfsmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfIfsm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugOspfIfsmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfIfsmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfIfsm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugOspfIfsm(d *schema.ResourceData) edpt.DebugOspfIfsm {
	var ret edpt.DebugOspfIfsm
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Status = d.Get("status").(int)
	ret.Inst.Timers = d.Get("timers").(int)
	//omit uuid
	return ret
}
