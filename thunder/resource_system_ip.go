package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ip`: IP Setting\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpCreate,
		UpdateContext: resourceSystemIpUpdate,
		ReadContext:   resourceSystemIpRead,
		DeleteContext: resourceSystemIpDelete,

		Schema: map[string]*schema.Schema{
			"class_e_address_range_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable class E (240.0.0.0/4) configuration",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemIp(d *schema.ResourceData) edpt.SystemIp {
	var ret edpt.SystemIp
	ret.Inst.ClassEAddressRangeEnable = d.Get("class_e_address_range_enable").(int)
	//omit uuid
	return ret
}
