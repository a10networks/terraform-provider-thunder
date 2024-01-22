package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_session_sync_reachability_options`: Options to determine reachability for session sync\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsDelete,

		Schema: map[string]*schema.Schema{
			"skip_default_route": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not choose default route for redirection(Not applicable in 'layer-2' mode)",
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
func resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSyncReachabilityOptions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSyncReachabilityOptions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSyncReachabilityOptions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncReachabilityOptionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSyncReachabilityOptions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutClusterLocalDeviceSessionSyncReachabilityOptions(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions {
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions
	ret.Inst.SkipDefaultRoute = d.Get("skip_default_route").(int)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
