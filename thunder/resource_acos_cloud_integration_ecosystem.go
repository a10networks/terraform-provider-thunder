package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosCloudIntegrationEcosystem() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_cloud_integration_ecosystem`: Ecosystem setting for heterogeneous cloud environment integration\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosCloudIntegrationEcosystemCreate,
		UpdateContext: resourceAcosCloudIntegrationEcosystemUpdate,
		ReadContext:   resourceAcosCloudIntegrationEcosystemRead,
		DeleteContext: resourceAcosCloudIntegrationEcosystemDelete,

		Schema: map[string]*schema.Schema{
			"consul": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_label": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_label_name": {
										Type: schema.TypeString, Optional: true, Description: "Name service group to be monitored",
									},
								},
							},
						},
						"ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "Configure the bootstrap server's IPv4 address (the host IPv4 address)",
						},
						"ipv6_address": {
							Type: schema.TypeString, Optional: true, Description: "Configure the bootstrap server's IPv6 address (the host IPv6 address)",
						},
						"host_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure the host name for bootstrap server(e.g www.a10networks.com)",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Configure the http port to use (port 8500)",
						},
						"health_check_interval": {
							Type: schema.TypeString, Optional: true, Description: "'5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '20': 20 seconds;",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Configuration; 'disable': Disable Configuration;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dummy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "dummy to make intermediate obj to single",
			},
			"k8s": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Configuration; 'disable': Disable Configuration;",
						},
						"health_check_interval": {
							Type: schema.TypeString, Optional: true, Description: "'5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '20': 20 seconds;",
						},
						"cluster_config_file": {
							Type: schema.TypeString, Optional: true, Description: "Enter cluster config file name",
						},
						"service_label": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_label_name": {
										Type: schema.TypeString, Optional: true, Description: "Name service group to be monitored",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"oracle": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_label": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_label_name": {
										Type: schema.TypeString, Optional: true, Description: "Name service group to be monitored",
									},
								},
							},
						},
						"ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "Configure the bootstrap server's IPv4 address (the host IPv4 address)",
						},
						"ipv6_address": {
							Type: schema.TypeString, Optional: true, Description: "Configure the bootstrap server's IPv6 address (the host IPv6 address)",
						},
						"host_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure the host name for bootstrap server(e.g www.a10networks.com)",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Configure the http port to use (port 8500)",
						},
						"health_check_interval": {
							Type: schema.TypeString, Optional: true, Description: "'5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '20': 20 seconds;",
						},
						"compartment_id": {
							Type: schema.TypeString, Optional: true, Description: "OCI compartment  id",
						},
						"tenancy_id": {
							Type: schema.TypeString, Optional: true, Description: "OCI tenancy  id",
						},
						"user_id": {
							Type: schema.TypeString, Optional: true, Description: "OCI user id",
						},
						"fingerprint": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"private_key_path": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Configuration; 'disable': Disable Configuration;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceAcosCloudIntegrationEcosystemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystem(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationEcosystemRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosCloudIntegrationEcosystemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystem(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationEcosystemRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosCloudIntegrationEcosystemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystem(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosCloudIntegrationEcosystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystem(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAcosCloudIntegrationEcosystemConsul41(d []interface{}) edpt.AcosCloudIntegrationEcosystemConsul41 {

	count1 := len(d)
	var ret edpt.AcosCloudIntegrationEcosystemConsul41
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceLabel = getSliceAcosCloudIntegrationEcosystemConsulServiceLabel42(in["service_label"].([]interface{}))
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv6Address = in["ipv6_address"].(string)
		ret.HostName = in["host_name"].(string)
		ret.Port = in["port"].(int)
		ret.HealthCheckInterval = in["health_check_interval"].(string)
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func getSliceAcosCloudIntegrationEcosystemConsulServiceLabel42(d []interface{}) []edpt.AcosCloudIntegrationEcosystemConsulServiceLabel42 {

	count1 := len(d)
	ret := make([]edpt.AcosCloudIntegrationEcosystemConsulServiceLabel42, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosCloudIntegrationEcosystemConsulServiceLabel42
		oi.ServiceLabelName = in["service_label_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAcosCloudIntegrationEcosystemK8s43(d []interface{}) edpt.AcosCloudIntegrationEcosystemK8s43 {

	count1 := len(d)
	var ret edpt.AcosCloudIntegrationEcosystemK8s43
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.HealthCheckInterval = in["health_check_interval"].(string)
		ret.ClusterConfigFile = in["cluster_config_file"].(string)
		ret.ServiceLabel = getSliceAcosCloudIntegrationEcosystemK8sServiceLabel44(in["service_label"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceAcosCloudIntegrationEcosystemK8sServiceLabel44(d []interface{}) []edpt.AcosCloudIntegrationEcosystemK8sServiceLabel44 {

	count1 := len(d)
	ret := make([]edpt.AcosCloudIntegrationEcosystemK8sServiceLabel44, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosCloudIntegrationEcosystemK8sServiceLabel44
		oi.ServiceLabelName = in["service_label_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAcosCloudIntegrationEcosystemOracle45(d []interface{}) edpt.AcosCloudIntegrationEcosystemOracle45 {

	count1 := len(d)
	var ret edpt.AcosCloudIntegrationEcosystemOracle45
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceLabel = getSliceAcosCloudIntegrationEcosystemOracleServiceLabel46(in["service_label"].([]interface{}))
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv6Address = in["ipv6_address"].(string)
		ret.HostName = in["host_name"].(string)
		ret.Port = in["port"].(int)
		ret.HealthCheckInterval = in["health_check_interval"].(string)
		ret.CompartmentId = in["compartment_id"].(string)
		ret.TenancyId = in["tenancy_id"].(string)
		ret.UserId = in["user_id"].(string)
		ret.Fingerprint = in["fingerprint"].(string)
		ret.PrivateKeyPath = in["private_key_path"].(string)
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func getSliceAcosCloudIntegrationEcosystemOracleServiceLabel46(d []interface{}) []edpt.AcosCloudIntegrationEcosystemOracleServiceLabel46 {

	count1 := len(d)
	ret := make([]edpt.AcosCloudIntegrationEcosystemOracleServiceLabel46, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosCloudIntegrationEcosystemOracleServiceLabel46
		oi.ServiceLabelName = in["service_label_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosCloudIntegrationEcosystem(d *schema.ResourceData) edpt.AcosCloudIntegrationEcosystem {
	var ret edpt.AcosCloudIntegrationEcosystem
	ret.Inst.Consul = getObjectAcosCloudIntegrationEcosystemConsul41(d.Get("consul").([]interface{}))
	ret.Inst.Dummy = d.Get("dummy").(int)
	ret.Inst.K8s = getObjectAcosCloudIntegrationEcosystemK8s43(d.Get("k8s").([]interface{}))
	ret.Inst.Oracle = getObjectAcosCloudIntegrationEcosystemOracle45(d.Get("oracle").([]interface{}))
	//omit uuid
	return ret
}
