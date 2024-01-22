package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsUsePartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_use_partition`: Configure Partition for Logging\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsUsePartitionCreate,
		UpdateContext: resourceAcosEventsUsePartitionUpdate,
		ReadContext:   resourceAcosEventsUsePartitionRead,
		DeleteContext: resourceAcosEventsUsePartitionDelete,

		Schema: map[string]*schema.Schema{
			"partition_name": {
				Type: schema.TypeString, Optional: true, Description: "Select partition name for logging",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsUsePartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsUsePartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsUsePartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsUsePartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsUsePartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsUsePartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsUsePartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsUsePartitionRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsUsePartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsUsePartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsUsePartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsUsePartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsUsePartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsUsePartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsUsePartition(d *schema.ResourceData) edpt.AcosEventsUsePartition {
	var ret edpt.AcosEventsUsePartition
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	//omit uuid
	return ret
}
