package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicThalesSecworld() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_thales_secworld`: import Thales security world files\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicThalesSecworldCreate,
		UpdateContext: resourceImportPeriodicThalesSecworldUpdate,
		ReadContext:   resourceImportPeriodicThalesSecworldRead,
		DeleteContext: resourceImportPeriodicThalesSecworldDelete,

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
			"thales_secworld": {
				Type: schema.TypeString, Required: true, Description: "import Thales security world files - in .tgz format that has all files needed by AX",
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
func resourceImportPeriodicThalesSecworldCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicThalesSecworldCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicThalesSecworld(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicThalesSecworldRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicThalesSecworldUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicThalesSecworldUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicThalesSecworld(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicThalesSecworldRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicThalesSecworldDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicThalesSecworldDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicThalesSecworld(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicThalesSecworldRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicThalesSecworldRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicThalesSecworld(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicThalesSecworld(d *schema.ResourceData) edpt.ImportPeriodicThalesSecworld {
	var ret edpt.ImportPeriodicThalesSecworld
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.ThalesSecworld = d.Get("thales_secworld").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
