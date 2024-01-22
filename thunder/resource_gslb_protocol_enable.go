package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbProtocolEnable() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_protocol_enable`: Enable/Disable GSLB Message Protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbProtocolEnableCreate,
		UpdateContext: resourceGslbProtocolEnableUpdate,
		ReadContext:   resourceGslbProtocolEnableRead,
		DeleteContext: resourceGslbProtocolEnableDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString, Required: true, Description: "'controller': Enable/Disable GSLB protocol as GSLB controller; 'device': Enable/Disable GSLB protocol as site device;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbProtocolEnableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolEnableCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolEnable(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbProtocolEnableRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbProtocolEnableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolEnableUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolEnable(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbProtocolEnableRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbProtocolEnableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolEnableDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolEnable(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbProtocolEnableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolEnableRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolEnable(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbProtocolEnable(d *schema.ResourceData) edpt.GslbProtocolEnable {
	var ret edpt.GslbProtocolEnable
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	return ret
}
