package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetIpv6RouterIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_ipv6_router_isis`: IS-IS Routing for IP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetIpv6RouterIsisCreate,
		UpdateContext: resourceInterfaceEthernetIpv6RouterIsisUpdate,
		ReadContext:   resourceInterfaceEthernetIpv6RouterIsisRead,
		DeleteContext: resourceInterfaceEthernetIpv6RouterIsisDelete,

		Schema: map[string]*schema.Schema{
			"tag": {
				Type: schema.TypeString, Optional: true, Description: "ISO routing area tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceEthernetIpv6RouterIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6RouterIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6RouterIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpv6RouterIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetIpv6RouterIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6RouterIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6RouterIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpv6RouterIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetIpv6RouterIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6RouterIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6RouterIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetIpv6RouterIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6RouterIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6RouterIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceEthernetIpv6RouterIsis(d *schema.ResourceData) edpt.InterfaceEthernetIpv6RouterIsis {
	var ret edpt.InterfaceEthernetIpv6RouterIsis
	ret.Inst.Tag = d.Get("tag").(string)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
