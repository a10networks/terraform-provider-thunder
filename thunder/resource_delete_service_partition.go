package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteServicePartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_service_partition`: Hard delete a service partition\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteServicePartitionCreate,
		UpdateContext: resourceDeleteServicePartitionUpdate,
		ReadContext:   resourceDeleteServicePartitionRead,
		DeleteContext: resourceDeleteServicePartitionDelete,

		Schema: map[string]*schema.Schema{
			"id1": {
				Type: schema.TypeInt, Optional: true, Description: "Specify unique service partition id",
			},
			"partition_name": {
				Type: schema.TypeString, Optional: true, Description: "Object service partition name",
			},
		},
	}
}
func resourceDeleteServicePartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteServicePartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteServicePartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteServicePartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteServicePartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteServicePartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteServicePartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteServicePartitionRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteServicePartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteServicePartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteServicePartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteServicePartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteServicePartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteServicePartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteServicePartition(d *schema.ResourceData) edpt.DeleteServicePartition {
	var ret edpt.DeleteServicePartition
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	return ret
}
