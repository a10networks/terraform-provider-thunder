package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceErase() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_erase`: Erase Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceEraseCreate,
		UpdateContext: resourceEraseUpdate,
		ReadContext:   resourceEraseRead,
		DeleteContext: resourceEraseDelete,

		Schema: map[string]*schema.Schema{
			"all_partitions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Wipe out all service config for all partitions",
			},
			"grubconfig": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reset boot grub settings",
			},
			"preserve_accounts": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "preserve admin accounts",
			},
			"preserve_management": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "preserve managememt ip and default gateway",
			},
			"reload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "reload after erase",
			},
			"service_config": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Wipe out all service config",
			},
		},
	}
}
func resourceEraseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEraseCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointErase(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEraseRead(ctx, d, meta)
	}
	return diags
}

func resourceEraseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEraseUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointErase(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEraseRead(ctx, d, meta)
	}
	return diags
}
func resourceEraseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEraseDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointErase(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEraseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEraseRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointErase(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointErase(d *schema.ResourceData) edpt.Erase {
	var ret edpt.Erase
	ret.Inst.AllPartitions = d.Get("all_partitions").(int)
	ret.Inst.Grubconfig = d.Get("grubconfig").(int)
	ret.Inst.PreserveAccounts = d.Get("preserve_accounts").(int)
	ret.Inst.PreserveManagement = d.Get("preserve_management").(int)
	ret.Inst.Reload = d.Get("reload").(int)
	ret.Inst.ServiceConfig = d.Get("service_config").(int)
	return ret
}
