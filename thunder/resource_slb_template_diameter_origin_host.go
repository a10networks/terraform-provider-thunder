package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDiameterOriginHost() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_diameter_origin_host`: origin host name avp\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDiameterOriginHostCreate,
		UpdateContext: resourceSlbTemplateDiameterOriginHostUpdate,
		ReadContext:   resourceSlbTemplateDiameterOriginHostRead,
		DeleteContext: resourceSlbTemplateDiameterOriginHostDelete,

		Schema: map[string]*schema.Schema{
			"origin_host_name": {
				Type: schema.TypeString, Optional: true, Description: "origin-host name avp",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateDiameterOriginHostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDiameterOriginHostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDiameterOriginHost(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDiameterOriginHostRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDiameterOriginHostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDiameterOriginHostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDiameterOriginHost(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDiameterOriginHostRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDiameterOriginHostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDiameterOriginHostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDiameterOriginHost(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDiameterOriginHostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDiameterOriginHostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDiameterOriginHost(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDiameterOriginHost(d *schema.ResourceData) edpt.SlbTemplateDiameterOriginHost {
	var ret edpt.SlbTemplateDiameterOriginHost
	ret.Inst.OriginHostName = d.Get("origin_host_name").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
