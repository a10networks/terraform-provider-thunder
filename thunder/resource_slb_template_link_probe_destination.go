package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateLinkProbeDestination() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_link_probe_destination`: Specify Target type\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateLinkProbeDestinationCreate,
		UpdateContext: resourceSlbTemplateLinkProbeDestinationUpdate,
		ReadContext:   resourceSlbTemplateLinkProbeDestinationRead,
		DeleteContext: resourceSlbTemplateLinkProbeDestinationDelete,

		Schema: map[string]*schema.Schema{
			"hostname": {
				Type: schema.TypeString, Optional: true, Description: "Target Hostname",
			},
			"resolve_as": {
				Type: schema.TypeString, Optional: true, Description: "'resolve-to-ipv4': Use A Query only to resolve the configured hostname; 'resolve-to-ipv6': Use AAAA Query only to resolve the configured hostname;",
			},
			"static_ipv4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Target IPv4 Address",
			},
			"static_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Target IPv6 Address",
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
func resourceSlbTemplateLinkProbeDestinationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkProbeDestinationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkProbeDestination(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkProbeDestinationRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateLinkProbeDestinationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkProbeDestinationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkProbeDestination(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkProbeDestinationRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateLinkProbeDestinationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkProbeDestinationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkProbeDestination(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateLinkProbeDestinationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkProbeDestinationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkProbeDestination(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateLinkProbeDestination(d *schema.ResourceData) edpt.SlbTemplateLinkProbeDestination {
	var ret edpt.SlbTemplateLinkProbeDestination
	ret.Inst.Hostname = d.Get("hostname").(string)
	ret.Inst.ResolveAs = d.Get("resolve_as").(string)
	ret.Inst.StaticIpv4Addr = d.Get("static_ipv4_addr").(string)
	ret.Inst.StaticIpv6Addr = d.Get("static_ipv6_addr").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
