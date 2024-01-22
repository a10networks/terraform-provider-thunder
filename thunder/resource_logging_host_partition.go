package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingHostPartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_host_partition`: Specify preferred partition configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingHostPartitionCreate,
		UpdateContext: resourceLoggingHostPartitionUpdate,
		ReadContext:   resourceLoggingHostPartitionRead,
		DeleteContext: resourceLoggingHostPartitionDelete,

		Schema: map[string]*schema.Schema{
			"partition_name": {
				Type: schema.TypeString, Optional: true, Description: "Select partition name for logging",
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select shared partition",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingHostPartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostPartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostPartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingHostPartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingHostPartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostPartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostPartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingHostPartitionRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingHostPartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostPartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostPartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingHostPartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostPartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostPartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingHostPartition(d *schema.ResourceData) edpt.LoggingHostPartition {
	var ret edpt.LoggingHostPartition
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	ret.Inst.Shared = d.Get("shared").(int)
	//omit uuid
	return ret
}
