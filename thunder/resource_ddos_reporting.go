package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosReporting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_reporting`: Reporting service configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosReportingCreate,
		UpdateContext: resourceDdosReportingUpdate,
		ReadContext:   resourceDdosReportingRead,
		DeleteContext: resourceDdosReportingDelete,

		Schema: map[string]*schema.Schema{
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "reject-on-limit-reached", Description: "'disable-on-limit-reached': Disable reporting on DST/Port entry when the max reporting count is reached; 'reject-on-limit-reached': Reject the configuration when the max reporting count is reached;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosReportingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosReportingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosReporting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosReportingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosReportingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosReportingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosReporting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosReportingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosReportingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosReportingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosReporting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosReportingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosReportingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosReporting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosReporting(d *schema.ResourceData) edpt.DdosReporting {
	var ret edpt.DdosReporting
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
