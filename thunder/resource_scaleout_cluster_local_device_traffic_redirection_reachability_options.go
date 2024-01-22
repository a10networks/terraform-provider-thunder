package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_traffic_redirection_reachability_options`: Options to determine reachability for redirection\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsDelete,

		Schema: map[string]*schema.Schema{
			"skip_default_route": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not choose default route for redirection",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"cluster_id": {
				Type: schema.TypeString, Required: true, Description: "ClusterId",
			},
		},
	}
}
func resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions {
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions
	ret.Inst.SkipDefaultRoute = d.Get("skip_default_route").(int)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
