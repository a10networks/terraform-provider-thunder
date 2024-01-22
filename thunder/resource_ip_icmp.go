package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpIcmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_icmp`: Global ICMP commands\n\n__PLACEHOLDER__",
		CreateContext: resourceIpIcmpCreate,
		UpdateContext: resourceIpIcmpUpdate,
		ReadContext:   resourceIpIcmpRead,
		DeleteContext: resourceIpIcmpDelete,

		Schema: map[string]*schema.Schema{
			"redirect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable outbound ICMP redirect messages",
			},
			"unreachable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable outbound ICMP unreachable messages",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpIcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpIcmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpIcmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpIcmpRead(ctx, d, meta)
	}
	return diags
}

func resourceIpIcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpIcmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpIcmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpIcmpRead(ctx, d, meta)
	}
	return diags
}
func resourceIpIcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpIcmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpIcmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpIcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpIcmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpIcmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpIcmp(d *schema.ResourceData) edpt.IpIcmp {
	var ret edpt.IpIcmp
	ret.Inst.Redirect = d.Get("redirect").(int)
	ret.Inst.Unreachable = d.Get("unreachable").(int)
	//omit uuid
	return ret
}
