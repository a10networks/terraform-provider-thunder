package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTechreportInterval() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_techreport_interval`: Configure polling techreport interval\n\n__PLACEHOLDER__",
		CreateContext: resourceTechreportIntervalCreate,
		UpdateContext: resourceTechreportIntervalUpdate,
		ReadContext:   resourceTechreportIntervalRead,
		DeleteContext: resourceTechreportIntervalDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Default: 15, Description: "Showtech interval in minutes (default is 15)",
			},
		},
	}
}
func resourceTechreportIntervalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportIntervalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportInterval(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportIntervalRead(ctx, d, meta)
	}
	return diags
}

func resourceTechreportIntervalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportIntervalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportInterval(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportIntervalRead(ctx, d, meta)
	}
	return diags
}
func resourceTechreportIntervalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportIntervalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportInterval(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTechreportIntervalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportIntervalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportInterval(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTechreportInterval(d *schema.ResourceData) edpt.TechreportInterval {
	var ret edpt.TechreportInterval
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	return ret
}
