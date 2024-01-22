package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTemplateLoggingSourceAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_template_logging_source_address`: Specify source address of logging packet\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTemplateLoggingSourceAddressCreate,
		UpdateContext: resourceFwTemplateLoggingSourceAddressUpdate,
		ReadContext:   resourceFwTemplateLoggingSourceAddressRead,
		DeleteContext: resourceFwTemplateLoggingSourceAddressDelete,

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
func resourceFwTemplateLoggingSourceAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingSourceAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingSourceAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingSourceAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTemplateLoggingSourceAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingSourceAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingSourceAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingSourceAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTemplateLoggingSourceAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingSourceAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingSourceAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTemplateLoggingSourceAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingSourceAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingSourceAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwTemplateLoggingSourceAddress(d *schema.ResourceData) edpt.FwTemplateLoggingSourceAddress {
	var ret edpt.FwTemplateLoggingSourceAddress
	ret.Inst.Ip = d.Get("ip").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
