package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicWsdl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_wsdl`: Web Services Definition Language File\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicWsdlCreate,
		UpdateContext: resourceImportPeriodicWsdlUpdate,
		ReadContext:   resourceImportPeriodicWsdlRead,
		DeleteContext: resourceImportPeriodicWsdlDelete,

		Schema: map[string]*schema.Schema{
			"password": {
				Type: schema.TypeString, Optional: true, Description: "Config remote server  password",
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the period in second",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"wsdl": {
				Type: schema.TypeString, Required: true, Description: "Web Services Definition Language File",
			},
		},
	}
}
func resourceImportPeriodicWsdlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicWsdlCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicWsdl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicWsdlRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicWsdlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicWsdlUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicWsdl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicWsdlRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicWsdlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicWsdlDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicWsdl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicWsdlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicWsdlRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicWsdl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicWsdl(d *schema.ResourceData) edpt.ImportPeriodicWsdl {
	var ret edpt.ImportPeriodicWsdl
	//omit encrypted
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	ret.Inst.Wsdl = d.Get("wsdl").(string)
	return ret
}
