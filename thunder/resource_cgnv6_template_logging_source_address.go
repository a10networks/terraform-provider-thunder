package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateLoggingSourceAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_logging_source_address`: Specify source address of logging packet\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateLoggingSourceAddressCreate,
		UpdateContext: resourceCgnv6TemplateLoggingSourceAddressUpdate,
		ReadContext:   resourceCgnv6TemplateLoggingSourceAddressRead,
		DeleteContext: resourceCgnv6TemplateLoggingSourceAddressDelete,

		Schema: map[string]*schema.Schema{
			"ip": {
				Type: schema.TypeString, Optional: true, Description: "Specify source IP address",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Specify source IPv6 address",
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
func resourceCgnv6TemplateLoggingSourceAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingSourceAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingSourceAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingSourceAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateLoggingSourceAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingSourceAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingSourceAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingSourceAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateLoggingSourceAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingSourceAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingSourceAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateLoggingSourceAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingSourceAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingSourceAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6TemplateLoggingSourceAddress(d *schema.ResourceData) edpt.Cgnv6TemplateLoggingSourceAddress {
	var ret edpt.Cgnv6TemplateLoggingSourceAddress
	ret.Inst.Ip = d.Get("ip").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
