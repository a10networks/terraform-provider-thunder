package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_scaleout_infrastructure_service_node`: Enable service node group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeCreate,
		UpdateContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeRead,
		DeleteContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeDelete,

		Schema: map[string]*schema.Schema{
			"local_device_disabled": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable local device disabled trap",
			},
			"service_master": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable service-master trap",
			},
			"traffic_map_update": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable traffic map update trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureServiceNode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureServiceNode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureServiceNode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureServiceNodeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureServiceNode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureServiceNode(d *schema.ResourceData) edpt.SnmpServerEnableTrapsScaleoutInfrastructureServiceNode {
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureServiceNode
	ret.Inst.LocalDeviceDisabled = d.Get("local_device_disabled").(int)
	ret.Inst.ServiceMaster = d.Get("service_master").(int)
	ret.Inst.TrafficMapUpdate = d.Get("traffic_map_update").(int)
	//omit uuid
	return ret
}
