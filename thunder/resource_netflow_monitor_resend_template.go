package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorResendTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor_resend_template`: Configure when to resend netflow template\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorResendTemplateCreate,
		UpdateContext: resourceNetflowMonitorResendTemplateUpdate,
		ReadContext:   resourceNetflowMonitorResendTemplateRead,
		DeleteContext: resourceNetflowMonitorResendTemplateDelete,

		Schema: map[string]*schema.Schema{
			"records": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "To resend template once for each number of records (Number of records: default is 1000, 0 means disable template resend based on record-count)",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 1800, Description: "To set time interval to resend template (number of seconds: default is 1800, 0 means disable template resend based on timeout)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceNetflowMonitorResendTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorResendTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorResendTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorResendTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorResendTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorResendTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorResendTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorResendTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorResendTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorResendTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorResendTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorResendTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorResendTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorResendTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetflowMonitorResendTemplate(d *schema.ResourceData) edpt.NetflowMonitorResendTemplate {
	var ret edpt.NetflowMonitorResendTemplate
	ret.Inst.Records = d.Get("records").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
