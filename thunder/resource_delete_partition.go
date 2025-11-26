package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeletePartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_partition`: hard delete a partition\n\n__PLACEHOLDER__",
		CreateContext: resourceDeletePartitionCreate,
		UpdateContext: resourceDeletePartitionUpdate,
		ReadContext:   resourceDeletePartitionRead,
		DeleteContext: resourceDeletePartitionDelete,

		Schema: map[string]*schema.Schema{
			"id1": {
				Type: schema.TypeInt, Optional: true, Description: "Specify unique Partition id",
			},
			"partition_name": {
				Type: schema.TypeString, Optional: true, Description: "Object partition name",
			},
		},
	}
}
func resourceDeletePartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeletePartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeletePartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeletePartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceDeletePartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeletePartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeletePartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeletePartitionRead(ctx, d, meta)
	}
	return diags
}
func resourceDeletePartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeletePartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeletePartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeletePartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeletePartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeletePartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeletePartition(d *schema.ResourceData) edpt.DeletePartition {
	var ret edpt.DeletePartition
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	return ret
}
