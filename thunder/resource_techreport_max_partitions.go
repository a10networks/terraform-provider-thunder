package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTechreportMaxPartitions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_techreport_max_partitions`: Maximum partitions to show in per periodic techsupport\n\n__PLACEHOLDER__",
		CreateContext: resourceTechreportMaxPartitionsCreate,
		UpdateContext: resourceTechreportMaxPartitionsUpdate,
		ReadContext:   resourceTechreportMaxPartitionsRead,
		DeleteContext: resourceTechreportMaxPartitionsDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Maximum partions to show in per periodic techsupport (default is 30)",
			},
		},
	}
}
func resourceTechreportMaxPartitionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportMaxPartitionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportMaxPartitions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportMaxPartitionsRead(ctx, d, meta)
	}
	return diags
}

func resourceTechreportMaxPartitionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportMaxPartitionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportMaxPartitions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportMaxPartitionsRead(ctx, d, meta)
	}
	return diags
}
func resourceTechreportMaxPartitionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportMaxPartitionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportMaxPartitions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTechreportMaxPartitionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportMaxPartitionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportMaxPartitions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTechreportMaxPartitions(d *schema.ResourceData) edpt.TechreportMaxPartitions {
	var ret edpt.TechreportMaxPartitions
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	return ret
}
