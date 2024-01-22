package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplicationType() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_application_type`: Configure application to be used in partition (ADC/CGNV6)\n\n__PLACEHOLDER__",
		CreateContext: resourceApplicationTypeCreate,
		UpdateContext: resourceApplicationTypeUpdate,
		ReadContext:   resourceApplicationTypeRead,
		DeleteContext: resourceApplicationTypeDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'adc': Application type ADC; 'cgnv6': Application type CGNv6;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceApplicationTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceApplicationTypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointApplicationType(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceApplicationTypeRead(ctx, d, meta)
	}
	return diags
}

func resourceApplicationTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceApplicationTypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointApplicationType(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceApplicationTypeRead(ctx, d, meta)
	}
	return diags
}
func resourceApplicationTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceApplicationTypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointApplicationType(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceApplicationTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceApplicationTypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointApplicationType(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointApplicationType(d *schema.ResourceData) edpt.ApplicationType {
	var ret edpt.ApplicationType
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	return ret
}
