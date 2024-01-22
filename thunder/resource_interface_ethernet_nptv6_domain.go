package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetNptv6Domain() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_nptv6_domain`: Apply NPTv6 translation domain on interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetNptv6DomainCreate,
		UpdateContext: resourceInterfaceEthernetNptv6DomainUpdate,
		ReadContext:   resourceInterfaceEthernetNptv6DomainRead,
		DeleteContext: resourceInterfaceEthernetNptv6DomainDelete,

		Schema: map[string]*schema.Schema{
			"bind_type": {
				Type: schema.TypeString, Required: true, Description: "'inside': This interface is connected to NPTv6 inside networks; 'outside': This interface is connected to NPTv6 outside networks;",
			},
			"domain_name": {
				Type: schema.TypeString, Required: true, Description: "NPTv6 domain name",
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
func resourceInterfaceEthernetNptv6DomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetNptv6DomainCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetNptv6Domain(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetNptv6DomainRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetNptv6DomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetNptv6DomainUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetNptv6Domain(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetNptv6DomainRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetNptv6DomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetNptv6DomainDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetNptv6Domain(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetNptv6DomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetNptv6DomainRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetNptv6Domain(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceEthernetNptv6Domain(d *schema.ResourceData) edpt.InterfaceEthernetNptv6Domain {
	var ret edpt.InterfaceEthernetNptv6Domain
	ret.Inst.BindType = d.Get("bind_type").(string)
	ret.Inst.DomainName = d.Get("domain_name").(string)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
