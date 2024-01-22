package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpmiIpsrc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ipmi_ipsrc`: Source of IP Address\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpmiIpsrcCreate,
		UpdateContext: resourceSystemIpmiIpsrcUpdate,
		ReadContext:   resourceSystemIpmiIpsrcRead,
		DeleteContext: resourceSystemIpmiIpsrcDelete,

		Schema: map[string]*schema.Schema{
			"dhcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP addr obtained by BMC running DHCP",
			},
			"static": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually configured static IP address",
			},
		},
	}
}
func resourceSystemIpmiIpsrcCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiIpsrcCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiIpsrc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiIpsrcRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpmiIpsrcUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiIpsrcUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiIpsrc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiIpsrcRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpmiIpsrcDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiIpsrcDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiIpsrc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpmiIpsrcRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiIpsrcRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiIpsrc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemIpmiIpsrc(d *schema.ResourceData) edpt.SystemIpmiIpsrc {
	var ret edpt.SystemIpmiIpsrc
	ret.Inst.Dhcp = d.Get("dhcp").(int)
	ret.Inst.Static = d.Get("static").(int)
	return ret
}
