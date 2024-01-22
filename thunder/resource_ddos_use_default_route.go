package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosUseDefaultRoute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_use_default_route`: Use default route for traffic\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosUseDefaultRouteCreate,
		UpdateContext: resourceDdosUseDefaultRouteUpdate,
		ReadContext:   resourceDdosUseDefaultRouteRead,
		DeleteContext: resourceDdosUseDefaultRouteDelete,

		Schema: map[string]*schema.Schema{
			"ethernet_start_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_start": {
							Type: schema.TypeInt, Optional: true, Description: "Traffic receive from the ethernet port will use default route",
						},
						"ethernet_end": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosUseDefaultRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosUseDefaultRouteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosUseDefaultRoute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosUseDefaultRouteRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosUseDefaultRouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosUseDefaultRouteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosUseDefaultRoute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosUseDefaultRouteRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosUseDefaultRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosUseDefaultRouteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosUseDefaultRoute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosUseDefaultRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosUseDefaultRouteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosUseDefaultRoute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosUseDefaultRouteEthernetStartCfg(d []interface{}) []edpt.DdosUseDefaultRouteEthernetStartCfg {

	count1 := len(d)
	ret := make([]edpt.DdosUseDefaultRouteEthernetStartCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosUseDefaultRouteEthernetStartCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosUseDefaultRoute(d *schema.ResourceData) edpt.DdosUseDefaultRoute {
	var ret edpt.DdosUseDefaultRoute
	ret.Inst.EthernetStartCfg = getSliceDdosUseDefaultRouteEthernetStartCfg(d.Get("ethernet_start_cfg").([]interface{}))
	//omit uuid
	return ret
}
