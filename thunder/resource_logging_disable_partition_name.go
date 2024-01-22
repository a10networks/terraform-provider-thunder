package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingDisablePartitionName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_disable_partition_name`: \n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingDisablePartitionNameCreate,
		UpdateContext: resourceLoggingDisablePartitionNameUpdate,
		ReadContext:   resourceLoggingDisablePartitionNameRead,
		DeleteContext: resourceLoggingDisablePartitionNameDelete,

		Schema: map[string]*schema.Schema{
			"disable_partition_name": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable partition name in logs",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingDisablePartitionNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingDisablePartitionNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingDisablePartitionName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingDisablePartitionNameRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingDisablePartitionNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingDisablePartitionNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingDisablePartitionName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingDisablePartitionNameRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingDisablePartitionNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingDisablePartitionNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingDisablePartitionName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingDisablePartitionNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingDisablePartitionNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingDisablePartitionName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingDisablePartitionName(d *schema.ResourceData) edpt.LoggingDisablePartitionName {
	var ret edpt.LoggingDisablePartitionName
	ret.Inst.DisablePartitionName = d.Get("disable_partition_name").(int)
	//omit uuid
	return ret
}
