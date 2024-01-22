package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosCloudIntegrationEcosystemOracle() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_cloud_integration_ecosystem_oracle`: Configure the oracle ecosystem for integration\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosCloudIntegrationEcosystemOracleCreate,
		UpdateContext: resourceAcosCloudIntegrationEcosystemOracleUpdate,
		ReadContext:   resourceAcosCloudIntegrationEcosystemOracleRead,
		DeleteContext: resourceAcosCloudIntegrationEcosystemOracleDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Configuration; 'disable': Disable Configuration;",
			},
			"compartment_id": {
				Type: schema.TypeString, Optional: true, Description: "OCI compartment  id",
			},
			"fingerprint": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"health_check_interval": {
				Type: schema.TypeString, Optional: true, Description: "'5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '20': 20 seconds;",
			},
			"host_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure the host name for bootstrap server(e.g www.a10networks.com)",
			},
			"ipv4_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure the bootstrap server's IPv4 address (the host IPv4 address)",
			},
			"ipv6_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure the bootstrap server's IPv6 address (the host IPv6 address)",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Configure the http port to use (port 8500)",
			},
			"private_key_path": {
				Type: schema.TypeString, Optional: true, Description: "",
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
			"tenancy_id": {
				Type: schema.TypeString, Optional: true, Description: "OCI tenancy  id",
			},
			"user_id": {
				Type: schema.TypeString, Optional: true, Description: "OCI user id",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosCloudIntegrationEcosystemOracleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemOracleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemOracle(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationEcosystemOracleRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosCloudIntegrationEcosystemOracleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemOracleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemOracle(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationEcosystemOracleRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosCloudIntegrationEcosystemOracleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemOracleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemOracle(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosCloudIntegrationEcosystemOracleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemOracleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemOracle(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosCloudIntegrationEcosystemOracleServiceLabel(d []interface{}) []edpt.AcosCloudIntegrationEcosystemOracleServiceLabel {

	count1 := len(d)
	ret := make([]edpt.AcosCloudIntegrationEcosystemOracleServiceLabel, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosCloudIntegrationEcosystemOracleServiceLabel
		oi.ServiceLabelName = in["service_label_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosCloudIntegrationEcosystemOracle(d *schema.ResourceData) edpt.AcosCloudIntegrationEcosystemOracle {
	var ret edpt.AcosCloudIntegrationEcosystemOracle
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.CompartmentId = d.Get("compartment_id").(string)
	ret.Inst.Fingerprint = d.Get("fingerprint").(string)
	ret.Inst.HealthCheckInterval = d.Get("health_check_interval").(string)
	ret.Inst.HostName = d.Get("host_name").(string)
	ret.Inst.Ipv4Address = d.Get("ipv4_address").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PrivateKeyPath = d.Get("private_key_path").(string)
	ret.Inst.ServiceLabel = getSliceAcosCloudIntegrationEcosystemOracleServiceLabel(d.Get("service_label").([]interface{}))
	ret.Inst.TenancyId = d.Get("tenancy_id").(string)
	ret.Inst.UserId = d.Get("user_id").(string)
	//omit uuid
	return ret
}
