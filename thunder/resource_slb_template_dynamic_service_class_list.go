package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDynamicServiceClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dynamic_service_class_list`: Configure Aho-Corasick class-list for specific DNS server\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDynamicServiceClassListCreate,
		UpdateContext: resourceSlbTemplateDynamicServiceClassListUpdate,
		ReadContext:   resourceSlbTemplateDynamicServiceClassListRead,
		DeleteContext: resourceSlbTemplateDynamicServiceClassListDelete,

		Schema: map[string]*schema.Schema{
			"dns_class_list": {
				Type: schema.TypeString, Required: true, Description: "Name of Aho-Corasick class-list",
			},
			"dns_server": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_dns_server": {
							Type: schema.TypeString, Optional: true, Description: "DNS Server IPv4 Address",
						},
						"ipv6_dns_server": {
							Type: schema.TypeString, Optional: true, Description: "DNS Server IPv6 Address",
						},
					},
				},
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Description: "Priority of the class-list(the larger number, the higher priority)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceSlbTemplateDynamicServiceClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDynamicServiceClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDynamicServiceClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDynamicServiceClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDynamicServiceClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDynamicServiceClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDynamicServiceClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDynamicServiceClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDynamicServiceClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDynamicServiceClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDynamicServiceClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDynamicServiceClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDynamicServiceClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDynamicServiceClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDynamicServiceClassListDnsServer(d []interface{}) []edpt.SlbTemplateDynamicServiceClassListDnsServer {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDynamicServiceClassListDnsServer, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDynamicServiceClassListDnsServer
		oi.Ipv4DnsServer = in["ipv4_dns_server"].(string)
		oi.Ipv6DnsServer = in["ipv6_dns_server"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDynamicServiceClassList(d *schema.ResourceData) edpt.SlbTemplateDynamicServiceClassList {
	var ret edpt.SlbTemplateDynamicServiceClassList
	ret.Inst.DnsClassList = d.Get("dns_class_list").(string)
	ret.Inst.DnsServer = getSliceSlbTemplateDynamicServiceClassListDnsServer(d.Get("dns_server").([]interface{}))
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
