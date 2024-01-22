package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTechreportMaxLogfileSize() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_techreport_max_logfile_size`: Maximum logfile size for periodic techsupport\n\n__PLACEHOLDER__",
		CreateContext: resourceTechreportMaxLogfileSizeCreate,
		UpdateContext: resourceTechreportMaxLogfileSizeUpdate,
		ReadContext:   resourceTechreportMaxLogfileSizeRead,
		DeleteContext: resourceTechreportMaxLogfileSizeDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Log file size for periodic techsupport (default is 1)",
			},
		},
	}
}
func resourceTechreportMaxLogfileSizeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportMaxLogfileSizeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportMaxLogfileSize(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportMaxLogfileSizeRead(ctx, d, meta)
	}
	return diags
}

func resourceTechreportMaxLogfileSizeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportMaxLogfileSizeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportMaxLogfileSize(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportMaxLogfileSizeRead(ctx, d, meta)
	}
	return diags
}
func resourceTechreportMaxLogfileSizeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportMaxLogfileSizeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportMaxLogfileSize(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTechreportMaxLogfileSizeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportMaxLogfileSizeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportMaxLogfileSize(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTechreportMaxLogfileSize(d *schema.ResourceData) edpt.TechreportMaxLogfileSize {
	var ret edpt.TechreportMaxLogfileSize
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	return ret
}
