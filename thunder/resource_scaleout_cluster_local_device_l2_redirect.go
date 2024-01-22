package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceL2Redirect() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_l2_redirect`: Configure interface for redirection\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceL2RedirectCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceL2RedirectUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceL2RedirectRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceL2RedirectDelete,

		Schema: map[string]*schema.Schema{
			"ethernet_vlan": {
				Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
			},
			"redirect_eth": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Port Value)",
			},
			"redirect_trunk": {
				Type: schema.TypeInt, Optional: true, Description: "L2 Trunk group",
			},
			"trunk_vlan": {
				Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
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
func resourceScaleoutClusterLocalDeviceL2RedirectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceL2RedirectCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceL2Redirect(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceL2RedirectRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceL2RedirectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceL2RedirectUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceL2Redirect(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceL2RedirectRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceL2RedirectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceL2RedirectDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceL2Redirect(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceL2RedirectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceL2RedirectRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceL2Redirect(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutClusterLocalDeviceL2Redirect(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceL2Redirect {
	var ret edpt.ScaleoutClusterLocalDeviceL2Redirect
	ret.Inst.EthernetVlan = d.Get("ethernet_vlan").(int)
	ret.Inst.RedirectEth = d.Get("redirect_eth").(int)
	ret.Inst.RedirectTrunk = d.Get("redirect_trunk").(int)
	ret.Inst.TrunkVlan = d.Get("trunk_vlan").(int)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
