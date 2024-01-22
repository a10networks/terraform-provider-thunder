package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_alg_dns`: Change Firewall DNS ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceFwAlgDnsCreate,
		UpdateContext: resourceFwAlgDnsUpdate,
		ReadContext:   resourceFwAlgDnsRead,
		DeleteContext: resourceFwAlgDnsDelete,

		Schema: map[string]*schema.Schema{
			"default_port_disable": {
				Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable DNS ALG default port 53;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwAlgDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceFwAlgDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceFwAlgDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwAlgDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwAlgDns(d *schema.ResourceData) edpt.FwAlgDns {
	var ret edpt.FwAlgDns
	ret.Inst.DefaultPortDisable = d.Get("default_port_disable").(string)
	//omit uuid
	return ret
}
