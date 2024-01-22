package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicThalesKmdata() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_thales_kmdata`: import Thales Kmdata files\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicThalesKmdataCreate,
		UpdateContext: resourceImportPeriodicThalesKmdataUpdate,
		ReadContext:   resourceImportPeriodicThalesKmdataRead,
		DeleteContext: resourceImportPeriodicThalesKmdataDelete,

		Schema: map[string]*schema.Schema{
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the period in second",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"thales_kmdata": {
				Type: schema.TypeString, Required: true, Description: "import Thales Kmdata files - in .tgz format that has all files needed by AX",
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
func resourceImportPeriodicThalesKmdataCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicThalesKmdataCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicThalesKmdata(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicThalesKmdataRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicThalesKmdataUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicThalesKmdataUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicThalesKmdata(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicThalesKmdataRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicThalesKmdataDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicThalesKmdataDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicThalesKmdata(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicThalesKmdataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicThalesKmdataRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicThalesKmdata(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicThalesKmdata(d *schema.ResourceData) edpt.ImportPeriodicThalesKmdata {
	var ret edpt.ImportPeriodicThalesKmdata
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.ThalesKmdata = d.Get("thales_kmdata").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
