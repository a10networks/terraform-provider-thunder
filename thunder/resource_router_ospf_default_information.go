package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspfDefaultInformation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ospf_default_information`: Control distribution of default information\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterOspfDefaultInformationCreate,
		UpdateContext: resourceRouterOspfDefaultInformationUpdate,
		ReadContext:   resourceRouterOspfDefaultInformationRead,
		DeleteContext: resourceRouterOspfDefaultInformationDelete,

		Schema: map[string]*schema.Schema{
			"always": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always advertise default route",
			},
			"metric": {
				Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
			},
			"metric_type": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "OSPF metric type for default routes",
			},
			"originate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Distribute a default route",
			},
			"route_map": {
				Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"process_id": {
				Type: schema.TypeString, Required: true, Description: "ProcessId",
			},
		},
	}
}
func resourceRouterOspfDefaultInformationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfDefaultInformationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfDefaultInformation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterOspfDefaultInformationRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterOspfDefaultInformationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfDefaultInformationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfDefaultInformation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterOspfDefaultInformationRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterOspfDefaultInformationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfDefaultInformationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfDefaultInformation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterOspfDefaultInformationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfDefaultInformationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfDefaultInformation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterOspfDefaultInformation(d *schema.ResourceData) edpt.RouterOspfDefaultInformation {
	var ret edpt.RouterOspfDefaultInformation
	ret.Inst.Always = d.Get("always").(int)
	ret.Inst.Metric = d.Get("metric").(int)
	ret.Inst.MetricType = d.Get("metric_type").(int)
	ret.Inst.Originate = d.Get("originate").(int)
	ret.Inst.RouteMap = d.Get("route_map").(string)
	//omit uuid
	ret.Inst.ProcessId = d.Get("process_id").(string)
	return ret
}
