package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosCloudIntegration() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_cloud_integration`: Setting for heterogeneous cloud environment integration\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosCloudIntegrationCreate,
		UpdateContext: resourceAcosCloudIntegrationUpdate,
		ReadContext:   resourceAcosCloudIntegrationRead,
		DeleteContext: resourceAcosCloudIntegrationDelete,

		Schema: map[string]*schema.Schema{
			"dummy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "dummy to make intermediate obj to single",
			},
			"ecosystem": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dummy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "dummy to make intermediate obj to single",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
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
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosCloudIntegrationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegration(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosCloudIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegration(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosCloudIntegrationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegration(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosCloudIntegrationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegration(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAcosCloudIntegrationEcosystem48(d []interface{}) edpt.AcosCloudIntegrationEcosystem48 {

	count1 := len(d)
	var ret edpt.AcosCloudIntegrationEcosystem48
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dummy = in["dummy"].(int)
		//omit uuid
		ret.Consul = getObjectAcosCloudIntegrationEcosystemConsul49(in["consul"].([]interface{}))
		ret.Oracle = getObjectAcosCloudIntegrationEcosystemOracle51(in["oracle"].([]interface{}))
		ret.K8s = getObjectAcosCloudIntegrationEcosystemK8s53(in["k8s"].([]interface{}))
	}
	return ret
}

func getObjectAcosCloudIntegrationEcosystemConsul49(d []interface{}) edpt.AcosCloudIntegrationEcosystemConsul49 {

	count1 := len(d)
	var ret edpt.AcosCloudIntegrationEcosystemConsul49
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceLabel = getSliceAcosCloudIntegrationEcosystemConsulServiceLabel50(in["service_label"].([]interface{}))
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

func getSliceAcosCloudIntegrationEcosystemConsulServiceLabel50(d []interface{}) []edpt.AcosCloudIntegrationEcosystemConsulServiceLabel50 {

	count1 := len(d)
	ret := make([]edpt.AcosCloudIntegrationEcosystemConsulServiceLabel50, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosCloudIntegrationEcosystemConsulServiceLabel50
		oi.ServiceLabelName = in["service_label_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAcosCloudIntegrationEcosystemOracle51(d []interface{}) edpt.AcosCloudIntegrationEcosystemOracle51 {

	count1 := len(d)
	var ret edpt.AcosCloudIntegrationEcosystemOracle51
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceLabel = getSliceAcosCloudIntegrationEcosystemOracleServiceLabel52(in["service_label"].([]interface{}))
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

func getSliceAcosCloudIntegrationEcosystemOracleServiceLabel52(d []interface{}) []edpt.AcosCloudIntegrationEcosystemOracleServiceLabel52 {

	count1 := len(d)
	ret := make([]edpt.AcosCloudIntegrationEcosystemOracleServiceLabel52, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosCloudIntegrationEcosystemOracleServiceLabel52
		oi.ServiceLabelName = in["service_label_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAcosCloudIntegrationEcosystemK8s53(d []interface{}) edpt.AcosCloudIntegrationEcosystemK8s53 {

	count1 := len(d)
	var ret edpt.AcosCloudIntegrationEcosystemK8s53
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.HealthCheckInterval = in["health_check_interval"].(string)
		ret.ClusterConfigFile = in["cluster_config_file"].(string)
		ret.ServiceLabel = getSliceAcosCloudIntegrationEcosystemK8sServiceLabel54(in["service_label"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceAcosCloudIntegrationEcosystemK8sServiceLabel54(d []interface{}) []edpt.AcosCloudIntegrationEcosystemK8sServiceLabel54 {

	count1 := len(d)
	ret := make([]edpt.AcosCloudIntegrationEcosystemK8sServiceLabel54, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosCloudIntegrationEcosystemK8sServiceLabel54
		oi.ServiceLabelName = in["service_label_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosCloudIntegration(d *schema.ResourceData) edpt.AcosCloudIntegration {
	var ret edpt.AcosCloudIntegration
	ret.Inst.Dummy = d.Get("dummy").(int)
	ret.Inst.Ecosystem = getObjectAcosCloudIntegrationEcosystem48(d.Get("ecosystem").([]interface{}))
	//omit uuid
	return ret
}
