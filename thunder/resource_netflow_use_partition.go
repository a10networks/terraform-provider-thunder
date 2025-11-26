package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowUsePartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_use_partition`: Use NetFlow configuration on a dedicated partition\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowUsePartitionCreate,
		UpdateContext: resourceNetflowUsePartitionUpdate,
		ReadContext:   resourceNetflowUsePartitionRead,
		DeleteContext: resourceNetflowUsePartitionDelete,

		Schema: map[string]*schema.Schema{
			"part_name": {
				Type: schema.TypeString, Optional: true, Description: "Dedicated partition name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetflowUsePartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowUsePartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowUsePartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowUsePartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowUsePartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowUsePartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowUsePartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowUsePartitionRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowUsePartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowUsePartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowUsePartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowUsePartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowUsePartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowUsePartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetflowUsePartition(d *schema.ResourceData) edpt.NetflowUsePartition {
	var ret edpt.NetflowUsePartition
	ret.Inst.PartName = d.Get("part_name").(string)
	//omit uuid
	return ret
}
