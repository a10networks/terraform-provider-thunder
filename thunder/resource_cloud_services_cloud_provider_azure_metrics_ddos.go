package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesCloudProviderAzureMetricsDdos() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_cloud_provider_azure_metrics_ddos`: Azure metrics configuration for TPS\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesCloudProviderAzureMetricsDdosCreate,
		UpdateContext: resourceCloudServicesCloudProviderAzureMetricsDdosUpdate,
		ReadContext:   resourceCloudServicesCloudProviderAzureMetricsDdosRead,
		DeleteContext: resourceCloudServicesCloudProviderAzureMetricsDdosDelete,

		Schema: map[string]*schema.Schema{
			"entry_zone": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable All DDoS Entries and Zones Stats; 'disable': Disable All DDoS Entries and Zones Stats;",
			},
			"port": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable All Port Metrics; 'disable': Disable All Port Metrics;",
			},
			"tcp": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable All TCP Metrics; 'disable': Disable All TCP Metrics;",
			},
			"udp": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable All UDP Metrics; 'disable': Disable All UDP Metrics;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCloudServicesCloudProviderAzureMetricsDdosCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureMetricsDdosCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureMetricsDdos(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAzureMetricsDdosRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAzureMetricsDdosUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureMetricsDdosUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureMetricsDdos(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAzureMetricsDdosRead(ctx, d, meta)
	}
	return diags
}
func resourceCloudServicesCloudProviderAzureMetricsDdosDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureMetricsDdosDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureMetricsDdos(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesCloudProviderAzureMetricsDdosRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureMetricsDdosRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureMetricsDdos(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCloudServicesCloudProviderAzureMetricsDdos(d *schema.ResourceData) edpt.CloudServicesCloudProviderAzureMetricsDdos {
	var ret edpt.CloudServicesCloudProviderAzureMetricsDdos
	ret.Inst.EntryZone = d.Get("entry_zone").(string)
	ret.Inst.Port = d.Get("port").(string)
	ret.Inst.Tcp = d.Get("tcp").(string)
	ret.Inst.Udp = d.Get("udp").(string)
	//omit uuid
	return ret
}
