package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsScaleoutInfrastructure() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_scaleout_infrastructure`: Enable Intrastructure group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureCreate,
		UpdateContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsScaleoutInfrastructureRead,
		DeleteContext: resourceSnmpServerEnableTrapsScaleoutInfrastructureDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all infra traps",
			},
			"cluster": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"master_node": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"traffic_map_distribution": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Traffic-map distribution trap",
						},
						"vserver_traffic_map_update": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VServer Traffic-map trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"service_node": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"test_send_all_traps": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send all infra traps",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsScaleoutInfrastructureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructure(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsScaleoutInfrastructureRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsScaleoutInfrastructureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructure(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsScaleoutInfrastructureRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsScaleoutInfrastructureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructure(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsScaleoutInfrastructureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsScaleoutInfrastructureRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructure(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructureCluster1479(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructureCluster1479 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureCluster1479
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Election = in["election"].(int)
		ret.MasterCallingReElection = in["master_calling_re_election"].(int)
		ret.NodeStatus = in["node_status"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructureMasterNode1480(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructureMasterNode1480 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureMasterNode1480
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TrafficMapDistribution = in["traffic_map_distribution"].(int)
		ret.VserverTrafficMapUpdate = in["vserver_traffic_map_update"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructureServiceNode1481(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructureServiceNode1481 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureServiceNode1481
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalDeviceDisabled = in["local_device_disabled"].(int)
		ret.ServiceMaster = in["service_master"].(int)
		ret.TrafficMapUpdate = in["traffic_map_update"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSnmpServerEnableTrapsScaleoutInfrastructure(d *schema.ResourceData) edpt.SnmpServerEnableTrapsScaleoutInfrastructure {
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructure
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Cluster = getObjectSnmpServerEnableTrapsScaleoutInfrastructureCluster1479(d.Get("cluster").([]interface{}))
	ret.Inst.MasterNode = getObjectSnmpServerEnableTrapsScaleoutInfrastructureMasterNode1480(d.Get("master_node").([]interface{}))
	ret.Inst.ServiceNode = getObjectSnmpServerEnableTrapsScaleoutInfrastructureServiceNode1481(d.Get("service_node").([]interface{}))
	ret.Inst.TestSendAllTraps = d.Get("test_send_all_traps").(int)
	//omit uuid
	return ret
}
