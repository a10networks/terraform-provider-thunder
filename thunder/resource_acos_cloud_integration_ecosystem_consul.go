package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosCloudIntegrationEcosystemConsul() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_cloud_integration_ecosystem_consul`: Configure the consul ecosystem for integration\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosCloudIntegrationEcosystemConsulCreate,
		UpdateContext: resourceAcosCloudIntegrationEcosystemConsulUpdate,
		ReadContext:   resourceAcosCloudIntegrationEcosystemConsulRead,
		DeleteContext: resourceAcosCloudIntegrationEcosystemConsulDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Configuration; 'disable': Disable Configuration;",
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
	}
}
func resourceAcosCloudIntegrationEcosystemConsulCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemConsulCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemConsul(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationEcosystemConsulRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosCloudIntegrationEcosystemConsulUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemConsulUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemConsul(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationEcosystemConsulRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosCloudIntegrationEcosystemConsulDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemConsulDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemConsul(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosCloudIntegrationEcosystemConsulRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemConsulRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemConsul(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosCloudIntegrationEcosystemConsulServiceLabel(d []interface{}) []edpt.AcosCloudIntegrationEcosystemConsulServiceLabel {

	count1 := len(d)
	ret := make([]edpt.AcosCloudIntegrationEcosystemConsulServiceLabel, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosCloudIntegrationEcosystemConsulServiceLabel
		oi.ServiceLabelName = in["service_label_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosCloudIntegrationEcosystemConsul(d *schema.ResourceData) edpt.AcosCloudIntegrationEcosystemConsul {
	var ret edpt.AcosCloudIntegrationEcosystemConsul
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.HealthCheckInterval = d.Get("health_check_interval").(string)
	ret.Inst.HostName = d.Get("host_name").(string)
	ret.Inst.Ipv4Address = d.Get("ipv4_address").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.ServiceLabel = getSliceAcosCloudIntegrationEcosystemConsulServiceLabel(d.Get("service_label").([]interface{}))
	//omit uuid
	return ret
}
