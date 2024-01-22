package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBiosProg() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_bios_prog`: Programming for BIOS\n\n__PLACEHOLDER__",
		CreateContext: resourceBiosProgCreate,
		UpdateContext: resourceBiosProgUpdate,
		ReadContext:   resourceBiosProgRead,
		DeleteContext: resourceBiosProgDelete,

		Schema: map[string]*schema.Schema{
			"filename": {
				Type: schema.TypeString, Optional: true, Description: "Give the filename of the flash.rom file",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceBiosProgCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBiosProgCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBiosProg(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBiosProgRead(ctx, d, meta)
	}
	return diags
}

func resourceBiosProgUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBiosProgUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBiosProg(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBiosProgRead(ctx, d, meta)
	}
	return diags
}
func resourceBiosProgDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBiosProgDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBiosProg(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBiosProgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBiosProgRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBiosProg(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointBiosProg(d *schema.ResourceData) edpt.BiosProg {
	var ret edpt.BiosProg
	ret.Inst.Filename = d.Get("filename").(string)
	//omit uuid
	return ret
}
