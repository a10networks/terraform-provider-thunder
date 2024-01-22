package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsSsl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_ssl`: Enable SSL group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsSslCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSslUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSslRead,
		DeleteContext: resourceSnmpServerEnableTrapsSslDelete,

		Schema: map[string]*schema.Schema{
			"server_certificate_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL server certificate error trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsSslCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSslCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSsl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSslRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsSslUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSslUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSsl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSslRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsSslDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSslDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSsl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsSslRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSslRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSsl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsSsl(d *schema.ResourceData) edpt.SnmpServerEnableTrapsSsl {
	var ret edpt.SnmpServerEnableTrapsSsl
	ret.Inst.ServerCertificateError = d.Get("server_certificate_error").(int)
	//omit uuid
	return ret
}
