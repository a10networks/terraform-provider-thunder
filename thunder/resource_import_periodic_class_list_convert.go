package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicClassListConvert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_class_list_convert`: Convert Class List File to A10 format\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicClassListConvertCreate,
		UpdateContext: resourceImportPeriodicClassListConvertUpdate,
		ReadContext:   resourceImportPeriodicClassListConvertRead,
		DeleteContext: resourceImportPeriodicClassListConvertDelete,

		Schema: map[string]*schema.Schema{
			"class_list_convert": {
				Type: schema.TypeString, Required: true, Description: "Class List File",
			},
			"class_list_type": {
				Type: schema.TypeString, Optional: true, Description: "'ac': ac; 'ipv4': ipv4; 'ipv6': ipv6; 'string': string; 'string-case-insensitive': string-case-insensitive;",
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
		},
	}
}
func resourceImportPeriodicClassListConvertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicClassListConvertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicClassListConvert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicClassListConvertRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicClassListConvertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicClassListConvertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicClassListConvert(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicClassListConvertRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicClassListConvertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicClassListConvertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicClassListConvert(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicClassListConvertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicClassListConvertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicClassListConvert(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicClassListConvert(d *schema.ResourceData) edpt.ImportPeriodicClassListConvert {
	var ret edpt.ImportPeriodicClassListConvert
	ret.Inst.ClassListConvert = d.Get("class_list_convert").(string)
	ret.Inst.ClassListType = d.Get("class_list_type").(string)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
