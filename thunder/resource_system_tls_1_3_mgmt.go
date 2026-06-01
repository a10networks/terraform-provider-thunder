package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTls13Mgmt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_tls_1_3_mgmt`: To Enable/disable TLS 1.3 support on ACOS management plane\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemTls13MgmtCreate,
		UpdateContext: resourceSystemTls13MgmtUpdate,
		ReadContext:   resourceSystemTls13MgmtRead,
		DeleteContext: resourceSystemTls13MgmtDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable TLS 1.3 support on ACOS management plane",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemTls13MgmtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTls13MgmtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTls13Mgmt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTls13MgmtRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemTls13MgmtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTls13MgmtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTls13Mgmt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTls13MgmtRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemTls13MgmtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTls13MgmtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTls13Mgmt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemTls13MgmtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTls13MgmtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTls13Mgmt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemTls13Mgmt(d *schema.ResourceData) edpt.SystemTls13Mgmt {
	var ret edpt.SystemTls13Mgmt
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
