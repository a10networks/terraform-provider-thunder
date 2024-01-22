package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMonTemplateLinkBlockAsDown() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_mon_template_link_block_as_down`: LACP Link-block will be treated as Link-down template\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMonTemplateLinkBlockAsDownCreate,
		UpdateContext: resourceSystemMonTemplateLinkBlockAsDownUpdate,
		ReadContext:   resourceSystemMonTemplateLinkBlockAsDownRead,
		DeleteContext: resourceSystemMonTemplateLinkBlockAsDownDelete,

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
func resourceSystemMonTemplateLinkBlockAsDownCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateLinkBlockAsDownCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateLinkBlockAsDown(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMonTemplateLinkBlockAsDownRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMonTemplateLinkBlockAsDownUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateLinkBlockAsDownUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateLinkBlockAsDown(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMonTemplateLinkBlockAsDownRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMonTemplateLinkBlockAsDownDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateLinkBlockAsDownDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateLinkBlockAsDown(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMonTemplateLinkBlockAsDownRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateLinkBlockAsDownRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateLinkBlockAsDown(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemMonTemplateLinkBlockAsDown(d *schema.ResourceData) edpt.SystemMonTemplateLinkBlockAsDown {
	var ret edpt.SystemMonTemplateLinkBlockAsDown
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
