package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TranslationServiceTimeout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_translation_service_timeout`: Specify any custom service timeout\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TranslationServiceTimeoutCreate,
		UpdateContext: resourceCgnv6TranslationServiceTimeoutUpdate,
		ReadContext:   resourceCgnv6TranslationServiceTimeoutRead,
		DeleteContext: resourceCgnv6TranslationServiceTimeoutDelete,

		Schema: map[string]*schema.Schema{
			"fast": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Fast Aging",
			},
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"port_end": {
				Type: schema.TypeInt, Optional: true, Description: "End Port Number",
			},
			"service_type": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP Protocol; 'udp': UDP Protocol;",
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
func resourceCgnv6TranslationServiceTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TranslationServiceTimeoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TranslationServiceTimeout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TranslationServiceTimeoutRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TranslationServiceTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TranslationServiceTimeoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TranslationServiceTimeout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TranslationServiceTimeoutRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TranslationServiceTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TranslationServiceTimeoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TranslationServiceTimeout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TranslationServiceTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TranslationServiceTimeoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TranslationServiceTimeout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6TranslationServiceTimeout(d *schema.ResourceData) edpt.Cgnv6TranslationServiceTimeout {
	var ret edpt.Cgnv6TranslationServiceTimeout
	ret.Inst.Fast = d.Get("fast").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PortEnd = d.Get("port_end").(int)
	ret.Inst.ServiceType = d.Get("service_type").(string)
	ret.Inst.TimeoutVal = d.Get("timeout_val").(int)
	//omit uuid
	return ret
}
