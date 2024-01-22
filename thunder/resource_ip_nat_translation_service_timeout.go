package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatTranslationServiceTimeout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_translation_service_timeout`: Specify any custom service timeout\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatTranslationServiceTimeoutCreate,
		UpdateContext: resourceIpNatTranslationServiceTimeoutUpdate,
		ReadContext:   resourceIpNatTranslationServiceTimeoutRead,
		DeleteContext: resourceIpNatTranslationServiceTimeoutDelete,

		Schema: map[string]*schema.Schema{
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"service_type": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP Protocol; 'udp': UDP Protocol;",
			},
			"timeout_type": {
				Type: schema.TypeString, Optional: true, Description: "'age': Expiration time; 'fast': Use Fast aging;",
			},
			"timeout_val": {
				Type: schema.TypeInt, Optional: true, Description: "Timeout in seconds (Interval of 60 seconds)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpNatTranslationServiceTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTranslationServiceTimeoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTranslationServiceTimeout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatTranslationServiceTimeoutRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatTranslationServiceTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTranslationServiceTimeoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTranslationServiceTimeout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatTranslationServiceTimeoutRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatTranslationServiceTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTranslationServiceTimeoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTranslationServiceTimeout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatTranslationServiceTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTranslationServiceTimeoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTranslationServiceTimeout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatTranslationServiceTimeout(d *schema.ResourceData) edpt.IpNatTranslationServiceTimeout {
	var ret edpt.IpNatTranslationServiceTimeout
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.ServiceType = d.Get("service_type").(string)
	ret.Inst.TimeoutType = d.Get("timeout_type").(string)
	ret.Inst.TimeoutVal = d.Get("timeout_val").(int)
	//omit uuid
	return ret
}
