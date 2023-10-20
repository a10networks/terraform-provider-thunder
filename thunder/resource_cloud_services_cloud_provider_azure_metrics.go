package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesCloudProviderAzureMetrics() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_cloud_provider_azure_metrics`: Azure metrics configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesCloudProviderAzureMetricsCreate,
		UpdateContext: resourceCloudServicesCloudProviderAzureMetricsUpdate,
		ReadContext:   resourceCloudServicesCloudProviderAzureMetricsRead,
		DeleteContext: resourceCloudServicesCloudProviderAzureMetricsDelete,
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Azure Metrics Analytics; 'disable': Disable Azure Metrics Analytics(default);",
			},
			"client_id": {
				Type: schema.TypeString, Optional: true, Description: "Azure Metrics Analytics Client id",
			},
			"cps": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable CPS Metrics; 'disable': Disable CPS Metrics;",
			},
			"cpu": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable CPU Metrics; 'disable': Disable CPU Metrics;",
			},
			"disk": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Disk Metrics; 'disable': Disable Disk Metrics;",
			},
			"interfaces": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Interfaces Metrics; 'disable': Disable Interfaces Metrics;",
			},
			"location": {
				Type: schema.TypeString, Optional: true, Description: "Azure Metrics Location",
			},
			"memory": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Memory Metrics; 'disable': Disable Memory Metrics;",
			},
			"packet_drop": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Packet Drop Metrics; 'disable': Disable Packet Drop Metrics;",
			},
			"packet_rate": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Packet Rate Metrics; 'disable': Disable Packet Rate Metrics;",
			},
			"resource_id": {
				Type: schema.TypeString, Optional: true, Description: "Resource/Instance ID of vThunder VMSS",
			},
			"secret_id": {
				Type: schema.TypeString, Optional: true, Description: "Azure Log Analytics Secret id",
			},
			"server_down_count": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Server Down Count Metrics; 'disable': Disable Server Down Count Metrics;",
			},
			"server_down_percentage": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Server Down Percentage Metrics; 'disable': Disable Server Down Percentage Metrics;",
			},
			"server_error": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Server Error Metrics; 'disable': Disable Server Error Metrics;",
			},
			"sessions": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Sessions Metrics; 'disable': Disable Sessions Metrics;",
			},
			"ssl_cert": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable SSL Cert Metrics; 'disable': Disable SSL Cert Metrics;",
			},
			"tenant_id": {
				Type: schema.TypeString, Optional: true, Description: "Azure Metrics Tenant ID",
			},
			"throughput": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Throughput Metrics; 'disable': Disable Throughput Metrics;",
			},
			"tps": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable TPS Metrics; 'disable': Disable TPS Metrics;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceCloudServicesCloudProviderAzureMetricsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureMetricsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureMetrics(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAzureMetricsRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAzureMetricsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureMetricsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureMetrics(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesCloudProviderAzureMetricsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureMetricsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureMetrics(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAzureMetricsRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAzureMetricsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureMetricsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureMetrics(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCloudServicesCloudProviderAzureMetrics(d *schema.ResourceData) edpt.CloudServicesCloudProviderAzureMetrics {
	var ret edpt.CloudServicesCloudProviderAzureMetrics
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ClientId = d.Get("client_id").(string)
	ret.Inst.Cps = d.Get("cps").(string)
	ret.Inst.Cpu = d.Get("cpu").(string)
	ret.Inst.Disk = d.Get("disk").(string)
	//omit encrypted
	ret.Inst.Interfaces = d.Get("interfaces").(string)
	ret.Inst.Location = d.Get("location").(string)
	ret.Inst.Memory = d.Get("memory").(string)
	ret.Inst.PacketDrop = d.Get("packet_drop").(string)
	ret.Inst.PacketRate = d.Get("packet_rate").(string)
	ret.Inst.ResourceId = d.Get("resource_id").(string)
	ret.Inst.SecretId = d.Get("secret_id").(string)
	ret.Inst.ServerDownCount = d.Get("server_down_count").(string)
	ret.Inst.ServerDownPercentage = d.Get("server_down_percentage").(string)
	ret.Inst.ServerError = d.Get("server_error").(string)
	ret.Inst.Sessions = d.Get("sessions").(string)
	ret.Inst.SslCert = d.Get("ssl_cert").(string)
	ret.Inst.TenantId = d.Get("tenant_id").(string)
	ret.Inst.Throughput = d.Get("throughput").(string)
	ret.Inst.Tps = d.Get("tps").(string)
	//omit uuid
	return ret
}
