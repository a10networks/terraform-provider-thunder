package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6OspfDefaultInformation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_ospf_default_information`: Control distribution of default information\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6OspfDefaultInformationCreate,
		UpdateContext: resourceRouterIpv6OspfDefaultInformationUpdate,
		ReadContext:   resourceRouterIpv6OspfDefaultInformationRead,
		DeleteContext: resourceRouterIpv6OspfDefaultInformationDelete,

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
func resourceRouterIpv6OspfDefaultInformationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfDefaultInformationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfDefaultInformation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6OspfDefaultInformationRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6OspfDefaultInformationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfDefaultInformationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfDefaultInformation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6OspfDefaultInformationRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6OspfDefaultInformationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfDefaultInformationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfDefaultInformation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6OspfDefaultInformationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfDefaultInformationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfDefaultInformation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterIpv6OspfDefaultInformation(d *schema.ResourceData) edpt.RouterIpv6OspfDefaultInformation {
	var ret edpt.RouterIpv6OspfDefaultInformation
	ret.Inst.Always = d.Get("always").(int)
	ret.Inst.Metric = d.Get("metric").(int)
	ret.Inst.MetricType = d.Get("metric_type").(int)
	ret.Inst.Originate = d.Get("originate").(int)
	ret.Inst.RouteMap = d.Get("route_map").(string)
	//omit uuid
	ret.Inst.ProcessId = d.Get("process_id").(string)
	return ret
}
