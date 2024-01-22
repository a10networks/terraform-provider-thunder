package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventPartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_event_partition`: module partition\n\n__PLACEHOLDER__",
		CreateContext: resourceEventPartitionCreate,
		UpdateContext: resourceEventPartitionUpdate,
		ReadContext:   resourceEventPartitionRead,
		DeleteContext: resourceEventPartitionDelete,

		Schema: map[string]*schema.Schema{
			"email": {
				Type: schema.TypeString, Optional: true, Description: "'on': enable this action; 'off': disable this action;",
			},
			"logging": {
				Type: schema.TypeString, Optional: true, Description: "'on': enable this action; 'off': disable this action;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vnp_events": {
				Type: schema.TypeString, Required: true, Description: "'part-create': Create new partition; 'part-del': Delete a partition;",
			},
		},
	}
}
func resourceEventPartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventPartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventPartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEventPartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceEventPartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventPartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventPartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEventPartitionRead(ctx, d, meta)
	}
	return diags
}
func resourceEventPartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventPartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventPartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEventPartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventPartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventPartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointEventPartition(d *schema.ResourceData) edpt.EventPartition {
	var ret edpt.EventPartition
	ret.Inst.Email = d.Get("email").(string)
	ret.Inst.Logging = d.Get("logging").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VnpEvents = d.Get("vnp_events").(string)
	return ret
}
