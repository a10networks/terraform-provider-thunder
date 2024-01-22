package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesKafkaServerProfile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_kafka_server_profile`: configure kafka server profile\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesKafkaServerProfileCreate,
		UpdateContext: resourceCloudServicesKafkaServerProfileUpdate,
		ReadContext:   resourceCloudServicesKafkaServerProfileRead,
		DeleteContext: resourceCloudServicesKafkaServerProfileDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'register': Register the device to the cloud; 'deregister': Deregister the device from the cloud;",
			},
			"active_ns": {
				Type: schema.TypeString, Optional: true, Description: "Set active_ns in cloud credentials file",
			},
			"alias_ns": {
				Type: schema.TypeString, Optional: true, Description: "Set active_ns in cloud credentials file",
			},
			"bootstrap_servers": {
				Type: schema.TypeString, Optional: true, Description: "Set bootstrap_servers in cloud credentials file",
			},
			"client_id": {
				Type: schema.TypeString, Optional: true, Description: "Set client_id in cloud credentials file",
			},
			"client_secret": {
				Type: schema.TypeString, Optional: true, Description: "Set client_secret in cloud credentials file",
			},
			"primary_ns": {
				Type: schema.TypeString, Optional: true, Description: "Set active_ns in cloud credentials file",
			},
			"resource_group": {
				Type: schema.TypeString, Optional: true, Description: "Set resource_group in cloud credentials file",
			},
			"sasl_mechanisms": {
				Type: schema.TypeString, Optional: true, Description: "Set sasl_mechanisms in cloud credentials file",
			},
			"sasl_password": {
				Type: schema.TypeString, Optional: true, Description: "Set sasl_password in cloud credentials file",
			},
			"schema_group_name": {
				Type: schema.TypeString, Optional: true, Description: "Set schema_group_name in cloud credentials file",
			},
			"secondary_ns": {
				Type: schema.TypeString, Optional: true, Description: "Set active_ns in cloud credentials file",
			},
			"security_protocol": {
				Type: schema.TypeString, Optional: true, Description: "Set security_protocol in cloud credentials file",
			},
			"subscription_id": {
				Type: schema.TypeString, Optional: true, Description: "Set subscription_id in cloud credentials file",
			},
			"tenant_id": {
				Type: schema.TypeString, Optional: true, Description: "Set tenant_id in cloud credentials file",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCloudServicesKafkaServerProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesKafkaServerProfileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesKafkaServerProfile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesKafkaServerProfileRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesKafkaServerProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesKafkaServerProfileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesKafkaServerProfile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesKafkaServerProfileRead(ctx, d, meta)
	}
	return diags
}
func resourceCloudServicesKafkaServerProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesKafkaServerProfileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesKafkaServerProfile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesKafkaServerProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesKafkaServerProfileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesKafkaServerProfile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCloudServicesKafkaServerProfile(d *schema.ResourceData) edpt.CloudServicesKafkaServerProfile {
	var ret edpt.CloudServicesKafkaServerProfile
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Active_ns = d.Get("active_ns").(string)
	ret.Inst.Alias_ns = d.Get("alias_ns").(string)
	ret.Inst.Bootstrap_servers = d.Get("bootstrap_servers").(string)
	ret.Inst.Client_id = d.Get("client_id").(string)
	ret.Inst.Client_secret = d.Get("client_secret").(string)
	ret.Inst.Primary_ns = d.Get("primary_ns").(string)
	ret.Inst.Resource_group = d.Get("resource_group").(string)
	ret.Inst.Sasl_mechanisms = d.Get("sasl_mechanisms").(string)
	ret.Inst.Sasl_password = d.Get("sasl_password").(string)
	ret.Inst.Schema_group_name = d.Get("schema_group_name").(string)
	ret.Inst.Secondary_ns = d.Get("secondary_ns").(string)
	ret.Inst.Security_protocol = d.Get("security_protocol").(string)
	ret.Inst.Subscription_id = d.Get("subscription_id").(string)
	ret.Inst.Tenant_id = d.Get("tenant_id").(string)
	//omit uuid
	return ret
}
