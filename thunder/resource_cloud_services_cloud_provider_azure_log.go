package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesCloudProviderAzureLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_cloud_provider_azure_log`: Azure log configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesCloudProviderAzureLogCreate,
		UpdateContext: resourceCloudServicesCloudProviderAzureLogUpdate,
		ReadContext:   resourceCloudServicesCloudProviderAzureLogRead,
		DeleteContext: resourceCloudServicesCloudProviderAzureLogDelete,
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Azure Log Analytics; 'disable': Disable Azure Log Analytics(default);",
			},
			"customer_id": {
				Type: schema.TypeString, Optional: true, Description: "Azure Log Analytics Customer ID",
			},
			"resource_id": {
				Type: schema.TypeString, Optional: true, Description: "Resource/Instance ID of vThunder",
			},
			"shared_key": {
				Type: schema.TypeString, Optional: true, Description: "Azure Log Analytics Shared Key",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceCloudServicesCloudProviderAzureLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAzureLogRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAzureLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesCloudProviderAzureLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAzureLogRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAzureLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAzureLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAzureLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCloudServicesCloudProviderAzureLog(d *schema.ResourceData) edpt.CloudServicesCloudProviderAzureLog {
	var ret edpt.CloudServicesCloudProviderAzureLog
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.CustomerId = d.Get("customer_id").(string)
	//omit encrypted
	ret.Inst.ResourceId = d.Get("resource_id").(string)
	ret.Inst.SharedKey = d.Get("shared_key").(string)
	//omit uuid
	return ret
}
