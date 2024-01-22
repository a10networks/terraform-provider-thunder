package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChassisApplicationType() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_chassis_application_type`: Configure application to be used for chassis (ADC/CGNv6)\n\n__PLACEHOLDER__",
		CreateContext: resourceChassisApplicationTypeCreate,
		UpdateContext: resourceChassisApplicationTypeUpdate,
		ReadContext:   resourceChassisApplicationTypeRead,
		DeleteContext: resourceChassisApplicationTypeDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString, Optional: true, Default: "cgnv6", Description: "'adc': Application type ADC; 'cgnv6': Application type CGNv6;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceChassisApplicationTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisApplicationTypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisApplicationType(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceChassisApplicationTypeRead(ctx, d, meta)
	}
	return diags
}

func resourceChassisApplicationTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisApplicationTypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisApplicationType(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceChassisApplicationTypeRead(ctx, d, meta)
	}
	return diags
}
func resourceChassisApplicationTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisApplicationTypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisApplicationType(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceChassisApplicationTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisApplicationTypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisApplicationType(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointChassisApplicationType(d *schema.ResourceData) edpt.ChassisApplicationType {
	var ret edpt.ChassisApplicationType
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	return ret
}
