package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceTrafficRedirectionEncap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_traffic_redirection_encap`: Encapsulation options to redirect taffic\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString, Required: true, Description: "'vxlan': Use vxlan for encapsulation;",
			},
			"use_v4_vxlan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always use IPv4 VxLAN for redirection",
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
func resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionEncap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionEncap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionEncap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionEncapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionEncap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionEncap(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionEncap {
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionEncap
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UseV4Vxlan = d.Get("use_v4_vxlan").(int)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
