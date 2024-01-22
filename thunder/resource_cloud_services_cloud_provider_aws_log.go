package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesCloudProviderAwsLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_cloud_provider_aws_log`: AWS log configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesCloudProviderAwsLogCreate,
		UpdateContext: resourceCloudServicesCloudProviderAwsLogUpdate,
		ReadContext:   resourceCloudServicesCloudProviderAwsLogRead,
		DeleteContext: resourceCloudServicesCloudProviderAwsLogDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable AWS CloudWatch; 'disable': Disable AWS CloudWatch (default);",
			},
			"active_partitions": {
				Type: schema.TypeString, Optional: true, Description: "Specifies the thunder active partition name separated by a comma for multiple values",
			},
			"log_group_name": {
				Type: schema.TypeString, Optional: true, Description: "Specifies the log group name under which all logs are sent to AWS CloudWatch",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCloudServicesCloudProviderAwsLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAwsLogRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAwsLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAwsLogRead(ctx, d, meta)
	}
	return diags
}
func resourceCloudServicesCloudProviderAwsLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesCloudProviderAwsLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCloudServicesCloudProviderAwsLog(d *schema.ResourceData) edpt.CloudServicesCloudProviderAwsLog {
	var ret edpt.CloudServicesCloudProviderAwsLog
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ActivePartitions = d.Get("active_partitions").(string)
	ret.Inst.LogGroupName = d.Get("log_group_name").(string)
	//omit uuid
	return ret
}
