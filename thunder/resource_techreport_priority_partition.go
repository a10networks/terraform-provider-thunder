package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTechreportPriorityPartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_techreport_priority_partition`: Configure partition to always poll for techreport\n\n__PLACEHOLDER__",
		CreateContext: resourceTechreportPriorityPartitionCreate,
		UpdateContext: resourceTechreportPriorityPartitionUpdate,
		ReadContext:   resourceTechreportPriorityPartitionRead,
		DeleteContext: resourceTechreportPriorityPartitionDelete,

		Schema: map[string]*schema.Schema{
			"part_name": {
				Type: schema.TypeString, Required: true, Description: "Name of partition always want to show in showtech (shared is always shown by default)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTechreportPriorityPartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportPriorityPartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportPriorityPartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportPriorityPartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceTechreportPriorityPartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportPriorityPartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportPriorityPartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportPriorityPartitionRead(ctx, d, meta)
	}
	return diags
}
func resourceTechreportPriorityPartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportPriorityPartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportPriorityPartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTechreportPriorityPartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportPriorityPartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreportPriorityPartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTechreportPriorityPartition(d *schema.ResourceData) edpt.TechreportPriorityPartition {
	var ret edpt.TechreportPriorityPartition
	ret.Inst.PartName = d.Get("part_name").(string)
	//omit uuid
	return ret
}
