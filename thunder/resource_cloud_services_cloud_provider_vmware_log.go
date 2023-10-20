package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesCloudProviderVmwareLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_cloud_provider_vmware_log`: VMware log configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesCloudProviderVmwareLogCreate,
		UpdateContext: resourceCloudServicesCloudProviderVmwareLogUpdate,
		ReadContext:   resourceCloudServicesCloudProviderVmwareLogRead,
		DeleteContext: resourceCloudServicesCloudProviderVmwareLogDelete,
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable VMware Log Analytics; 'disable': Disable VMware Log Analytics(default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrli_host": {
				Type: schema.TypeString, Optional: true, Description: "VMware Log Analytics vrli-host",
			},
		},
	}
}

func resourceCloudServicesCloudProviderVmwareLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderVmwareLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderVmwareLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderVmwareLogRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderVmwareLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderVmwareLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderVmwareLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesCloudProviderVmwareLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderVmwareLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderVmwareLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderVmwareLogRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderVmwareLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderVmwareLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderVmwareLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCloudServicesCloudProviderVmwareLog(d *schema.ResourceData) edpt.CloudServicesCloudProviderVmwareLog {
	var ret edpt.CloudServicesCloudProviderVmwareLog
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	ret.Inst.VrliHost = d.Get("vrli_host").(string)
	return ret
}
