package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateLinkBlockAsDown() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_link_block_as_down`: LACP Link-block will be treated as Link-down template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateLinkBlockAsDownCreate,
		UpdateContext: resourceSlbTemplateLinkBlockAsDownUpdate,
		ReadContext:   resourceSlbTemplateLinkBlockAsDownRead,
		DeleteContext: resourceSlbTemplateLinkBlockAsDownDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateLinkBlockAsDownCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkBlockAsDownCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkBlockAsDown(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkBlockAsDownRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateLinkBlockAsDownUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkBlockAsDownUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkBlockAsDown(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkBlockAsDownRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateLinkBlockAsDownDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkBlockAsDownDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkBlockAsDown(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateLinkBlockAsDownRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkBlockAsDownRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkBlockAsDown(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateLinkBlockAsDown(d *schema.ResourceData) edpt.SlbTemplateLinkBlockAsDown {
	var ret edpt.SlbTemplateLinkBlockAsDown
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
