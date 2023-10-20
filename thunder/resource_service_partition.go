package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServicePartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_service_partition`: Create/unload a service partition\n\n__PLACEHOLDER__",
		CreateContext: resourceServicePartitionCreate,
		UpdateContext: resourceServicePartitionUpdate,
		ReadContext:   resourceServicePartitionRead,
		DeleteContext: resourceServicePartitionDelete,
		Schema: map[string]*schema.Schema{
			"application_type": {
				Type: schema.TypeString, Optional: true, Description: "'adc': Application type ADC;",
			},
			"follow_vrid": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Join a vrrp group (Specify ha VRRP-A vrid)",
			},
			"id1": {
				Type: schema.TypeInt, Optional: true, Description: "Specify unique service partition id",
			},
			"partition_name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Object service partition name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceServicePartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceServicePartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointServicePartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceServicePartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceServicePartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceServicePartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointServicePartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceServicePartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceServicePartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointServicePartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceServicePartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceServicePartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceServicePartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointServicePartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointServicePartition(d *schema.ResourceData) edpt.ServicePartition {
	var ret edpt.ServicePartition
	ret.Inst.ApplicationType = d.Get("application_type").(string)
	ret.Inst.FollowVrid = d.Get("follow_vrid").(int)
	ret.Inst.Id = d.Get("id1").(int)
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
