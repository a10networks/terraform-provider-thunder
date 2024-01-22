package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicSslCrl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_ssl_crl`: SSL Crl File\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicSslCrlCreate,
		UpdateContext: resourceImportPeriodicSslCrlUpdate,
		ReadContext:   resourceImportPeriodicSslCrlRead,
		DeleteContext: resourceImportPeriodicSslCrlDelete,

		Schema: map[string]*schema.Schema{
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the period in second",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"ssl_crl": {
				Type: schema.TypeString, Required: true, Description: "SSL Crl File",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceImportPeriodicSslCrlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCrlCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCrl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicSslCrlRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicSslCrlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCrlUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCrl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicSslCrlRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicSslCrlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCrlDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCrl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicSslCrlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCrlRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCrl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicSslCrl(d *schema.ResourceData) edpt.ImportPeriodicSslCrl {
	var ret edpt.ImportPeriodicSslCrl
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.SslCrl = d.Get("ssl_crl").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
