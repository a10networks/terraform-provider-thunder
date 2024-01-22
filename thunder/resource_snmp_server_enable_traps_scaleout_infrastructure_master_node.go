package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_scaleout_infrastructure_master_node`: Enable master node group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeCreate,
		UpdateContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeRead,
		DeleteContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeDelete,

		Schema: map[string]*schema.Schema{
			"traffic_map_distribution": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Traffic-map distribution trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vserver_traffic_map_update": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VServer Traffic-map trap",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureMasterNode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureMasterNode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureMasterNode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureMasterNodeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureMasterNode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureMasterNode(d *schema.ResourceData) edpt.SnmpServerEnableTrapsScaleoutInfrastructureMasterNode {
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureMasterNode
	ret.Inst.TrafficMapDistribution = d.Get("traffic_map_distribution").(int)
	//omit uuid
	ret.Inst.VserverTrafficMapUpdate = d.Get("vserver_traffic_map_update").(int)
	return ret
}
