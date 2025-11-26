package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsCategoryLookup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_category_lookup`: Configure web-category list for DNS domains matching (Up to 8 lists)\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsCategoryLookupCreate,
		UpdateContext: resourceSlbTemplateDnsCategoryLookupUpdate,
		ReadContext:   resourceSlbTemplateDnsCategoryLookupRead,
		DeleteContext: resourceSlbTemplateDnsCategoryLookupDelete,

		Schema: map[string]*schema.Schema{
			"category_name": {
				Type: schema.TypeString, Required: true, Description: "category-list name",
			},
			"drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Deny matching DNS domains",
			},
			"permit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Permit matching DNS domains",
			},
			"respond": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond to matching DNS domains",
			},
			"respond_cname_str": {
				Type: schema.TypeString, Optional: true, Description: "CNAME to respond (Canonical name)",
			},
			"respond_ip_addr": {
				Type: schema.TypeString, Optional: true, Description: "Type A record to respond (IPv4 address)",
			},
			"respond_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "TYPE AAAA record to respond (IPv6 address)",
			},
			"respond_nxdomain": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond with NXDOMAIN",
			},
			"response_ttl": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Set response TTL in seconds (TTL value in seconds)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"dns_name": {
				Type: schema.TypeString, Required: true, Description: "Dns_name",
			},
		},
	}
}
func resourceSlbTemplateDnsCategoryLookupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsCategoryLookupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsCategoryLookup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsCategoryLookupRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsCategoryLookupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsCategoryLookupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsCategoryLookup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsCategoryLookupRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsCategoryLookupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsCategoryLookupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsCategoryLookup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsCategoryLookupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsCategoryLookupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsCategoryLookup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsCategoryLookup(d *schema.ResourceData) edpt.SlbTemplateDnsCategoryLookup {
	var ret edpt.SlbTemplateDnsCategoryLookup
	ret.Inst.CategoryName = d.Get("category_name").(string)
	ret.Inst.Drop = d.Get("drop").(int)
	ret.Inst.Permit = d.Get("permit").(int)
	ret.Inst.Respond = d.Get("respond").(int)
	ret.Inst.RespondCnameStr = d.Get("respond_cname_str").(string)
	ret.Inst.RespondIpAddr = d.Get("respond_ip_addr").(string)
	ret.Inst.RespondIpv6Addr = d.Get("respond_ipv6_addr").(string)
	ret.Inst.RespondNxdomain = d.Get("respond_nxdomain").(int)
	ret.Inst.ResponseTtl = d.Get("response_ttl").(int)
	//omit uuid
	ret.Inst.Dns_name = d.Get("dns_name").(string)
	return ret
}
