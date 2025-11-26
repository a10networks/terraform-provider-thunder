package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosCloudIntegrationEcosystemK8s() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_cloud_integration_ecosystem_k8s`: Configure the kubernetes ecosystem for integration\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosCloudIntegrationEcosystemK8sCreate,
		UpdateContext: resourceAcosCloudIntegrationEcosystemK8sUpdate,
		ReadContext:   resourceAcosCloudIntegrationEcosystemK8sRead,
		DeleteContext: resourceAcosCloudIntegrationEcosystemK8sDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Configuration; 'disable': Disable Configuration;",
			},
			"cluster_config_file": {
				Type: schema.TypeString, Optional: true, Description: "Enter cluster config file name",
			},
			"health_check_interval": {
				Type: schema.TypeString, Optional: true, Description: "'5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '20': 20 seconds;",
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
func resourceAcosCloudIntegrationEcosystemK8sCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemK8sCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemK8s(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationEcosystemK8sRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosCloudIntegrationEcosystemK8sUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemK8sUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemK8s(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosCloudIntegrationEcosystemK8sRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosCloudIntegrationEcosystemK8sDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemK8sDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemK8s(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosCloudIntegrationEcosystemK8sRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosCloudIntegrationEcosystemK8sRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosCloudIntegrationEcosystemK8s(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosCloudIntegrationEcosystemK8sServiceLabel(d []interface{}) []edpt.AcosCloudIntegrationEcosystemK8sServiceLabel {

	count1 := len(d)
	ret := make([]edpt.AcosCloudIntegrationEcosystemK8sServiceLabel, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosCloudIntegrationEcosystemK8sServiceLabel
		oi.ServiceLabelName = in["service_label_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosCloudIntegrationEcosystemK8s(d *schema.ResourceData) edpt.AcosCloudIntegrationEcosystemK8s {
	var ret edpt.AcosCloudIntegrationEcosystemK8s
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ClusterConfigFile = d.Get("cluster_config_file").(string)
	ret.Inst.HealthCheckInterval = d.Get("health_check_interval").(string)
	ret.Inst.ServiceLabel = getSliceAcosCloudIntegrationEcosystemK8sServiceLabel(d.Get("service_label").([]interface{}))
	//omit uuid
	return ret
}
