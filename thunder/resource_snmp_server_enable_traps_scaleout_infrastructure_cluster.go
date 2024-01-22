package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsScaleoutInfrastructureCluster() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_scaleout_infrastructure_cluster`: Enable cluster group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterCreate,
		UpdateContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterRead,
		DeleteContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterDelete,

		Schema: map[string]*schema.Schema{
			"election": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable election status trap",
			},
			"master_calling_re_election": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable re-election trap",
			},
			"node_status": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable active node status trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureCluster(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureCluster(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureCluster(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureClusterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureCluster(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructureCluster(d *schema.ResourceData) edpt.SnmpServerEnableTrapsScaleoutInfrastructureCluster {
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureCluster
	ret.Inst.Election = d.Get("election").(int)
	ret.Inst.MasterCallingReElection = d.Get("master_calling_re_election").(int)
	ret.Inst.NodeStatus = d.Get("node_status").(int)
	//omit uuid
	return ret
}
