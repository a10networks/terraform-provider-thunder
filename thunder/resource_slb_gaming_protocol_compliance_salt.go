package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbGamingProtocolComplianceSalt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_gaming_protocol_compliance_salt`: Gaming Protocol Compliance Salt\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbGamingProtocolComplianceSaltCreate,
		UpdateContext: resourceSlbGamingProtocolComplianceSaltUpdate,
		ReadContext:   resourceSlbGamingProtocolComplianceSaltRead,
		DeleteContext: resourceSlbGamingProtocolComplianceSaltDelete,

		Schema: map[string]*schema.Schema{
			"current_salt": {
				Type: schema.TypeInt, Optional: true, Description: "Current salt value",
			},
			"previous_salt": {
				Type: schema.TypeInt, Optional: true, Description: "Previous salt value",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbGamingProtocolComplianceSaltCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGamingProtocolComplianceSaltCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGamingProtocolComplianceSalt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbGamingProtocolComplianceSaltRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbGamingProtocolComplianceSaltUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGamingProtocolComplianceSaltUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGamingProtocolComplianceSalt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbGamingProtocolComplianceSaltRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbGamingProtocolComplianceSaltDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGamingProtocolComplianceSaltDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGamingProtocolComplianceSalt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbGamingProtocolComplianceSaltRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGamingProtocolComplianceSaltRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGamingProtocolComplianceSalt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbGamingProtocolComplianceSalt(d *schema.ResourceData) edpt.SlbGamingProtocolComplianceSalt {
	var ret edpt.SlbGamingProtocolComplianceSalt
	ret.Inst.CurrentSalt = d.Get("current_salt").(int)
	ret.Inst.PreviousSalt = d.Get("previous_salt").(int)
	//omit uuid
	return ret
}
