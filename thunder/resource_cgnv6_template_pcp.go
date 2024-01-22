package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplatePcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_pcp`: Port Control Protocol Template\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplatePcpCreate,
		UpdateContext: resourceCgnv6TemplatePcpUpdate,
		ReadContext:   resourceCgnv6TemplatePcpRead,
		DeleteContext: resourceCgnv6TemplatePcpDelete,

		Schema: map[string]*schema.Schema{
			"allow_third_party_from_lan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow third party request coming from LAN (default is disabled)",
			},
			"allow_third_party_from_wan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow third party request coming from WAN (default is disabled)",
			},
			"announce": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "PCP ANNOUNCE Opcode (default is enabled)",
			},
			"check_client_nonce": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "To validate NONCE value in PCP request (default: disabled)",
			},
			"disable_map_filter": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "To disable processing of FILTER options in MAP request",
			},
			"map": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "PCP MAP Opcode (default is enabled)",
			},
			"maximum": {
				Type: schema.TypeInt, Optional: true, Default: 1440, Description: "To set maximum lifetime of PCP mappings (default 1440 minutes)",
			},
			"minimum": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "To set minimum lifetime of PCP mappings (default 2 minutes)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "PCP Template name",
			},
			"pcp_server_port": {
				Type: schema.TypeInt, Optional: true, Default: 5351, Description: "PCP server listening port (default 5351) (PCP UDP destination port)",
			},
			"peer": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "PCP PEER Opcode (default is enabled)",
			},
			"source_ip": {
				Type: schema.TypeString, Optional: true, Description: "Specify source IP address for IPv4 ANNOUNCE message",
			},
			"source_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Specify source IPv6 address for IPv6 ANNOUNCE message",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6TemplatePcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplatePcpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplatePcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplatePcpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplatePcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplatePcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplatePcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplatePcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6TemplatePcp(d *schema.ResourceData) edpt.Cgnv6TemplatePcp {
	var ret edpt.Cgnv6TemplatePcp
	ret.Inst.AllowThirdPartyFromLan = d.Get("allow_third_party_from_lan").(int)
	ret.Inst.AllowThirdPartyFromWan = d.Get("allow_third_party_from_wan").(int)
	ret.Inst.Announce = d.Get("announce").(int)
	ret.Inst.CheckClientNonce = d.Get("check_client_nonce").(int)
	ret.Inst.DisableMapFilter = d.Get("disable_map_filter").(int)
	ret.Inst.Map = d.Get("map").(int)
	ret.Inst.Maximum = d.Get("maximum").(int)
	ret.Inst.Minimum = d.Get("minimum").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PcpServerPort = d.Get("pcp_server_port").(int)
	ret.Inst.Peer = d.Get("peer").(int)
	ret.Inst.SourceIp = d.Get("source_ip").(string)
	ret.Inst.SourceIpv6 = d.Get("source_ipv6").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
