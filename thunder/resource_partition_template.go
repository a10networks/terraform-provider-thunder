package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePartitionTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_partition_template`: Attach a template to the partition\n\n__PLACEHOLDER__",
		CreateContext: resourcePartitionTemplateCreate,
		UpdateContext: resourcePartitionTemplateUpdate,
		ReadContext:   resourcePartitionTemplateRead,
		DeleteContext: resourcePartitionTemplateDelete,

		Schema: map[string]*schema.Schema{
			"resource_accounting": {
				Type: schema.TypeString, Optional: true, Description: "Attach a resource accounting template (Name of the template)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"partition_name": {
				Type: schema.TypeString, Required: true, Description: "PartitionName",
			},
		},
	}
}
func resourcePartitionTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourcePartitionTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourcePartitionTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePartitionTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPartitionTemplate(d *schema.ResourceData) edpt.PartitionTemplate {
	var ret edpt.PartitionTemplate
	ret.Inst.ResourceAccounting = d.Get("resource_accounting").(string)
	//omit uuid
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	return ret
}
