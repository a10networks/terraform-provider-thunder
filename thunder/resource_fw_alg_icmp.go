package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgIcmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_alg_icmp`: Change Firewall ICMP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceFwAlgIcmpCreate,
		UpdateContext: resourceFwAlgIcmpUpdate,
		ReadContext:   resourceFwAlgIcmpRead,
		DeleteContext: resourceFwAlgIcmpDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable ICMP ALG which allows ICMP errors to pass the firewall;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwAlgIcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgIcmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgIcmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgIcmpRead(ctx, d, meta)
	}
	return diags
}

func resourceFwAlgIcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgIcmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgIcmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgIcmpRead(ctx, d, meta)
	}
	return diags
}
func resourceFwAlgIcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgIcmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgIcmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwAlgIcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgIcmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgIcmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwAlgIcmp(d *schema.ResourceData) edpt.FwAlgIcmp {
	var ret edpt.FwAlgIcmp
	ret.Inst.Disable = d.Get("disable").(string)
	//omit uuid
	return ret
}
