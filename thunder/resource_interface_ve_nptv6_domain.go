package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVeNptv6Domain() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ve_nptv6_domain`: Apply NPTv6 translation domain on interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceVeNptv6DomainCreate,
		UpdateContext: resourceInterfaceVeNptv6DomainUpdate,
		ReadContext:   resourceInterfaceVeNptv6DomainRead,
		DeleteContext: resourceInterfaceVeNptv6DomainDelete,

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
func resourceInterfaceVeNptv6DomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeNptv6DomainCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeNptv6Domain(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeNptv6DomainRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceVeNptv6DomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeNptv6DomainUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeNptv6Domain(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeNptv6DomainRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceVeNptv6DomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeNptv6DomainDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeNptv6Domain(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceVeNptv6DomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeNptv6DomainRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeNptv6Domain(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceVeNptv6Domain(d *schema.ResourceData) edpt.InterfaceVeNptv6Domain {
	var ret edpt.InterfaceVeNptv6Domain
	ret.Inst.BindType = d.Get("bind_type").(string)
	ret.Inst.DomainName = d.Get("domain_name").(string)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
