package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSshd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sshd`: SSHD service operation\n\n__PLACEHOLDER__",
		CreateContext: resourceSshdCreate,
		UpdateContext: resourceSshdUpdate,
		ReadContext:   resourceSshdRead,
		DeleteContext: resourceSshdDelete,

		Schema: map[string]*schema.Schema{
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"generate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Generate SSH key",
			},
			"load": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load SSH key",
			},
			"re_add_rsa": {
				Type: schema.TypeString, Optional: true, Description: "remove ip address from known_hosts",
			},
			"regenerate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Wipe and generate SSH key",
			},
			"restart": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Restart SSH service",
			},
			"size": {
				Type: schema.TypeString, Optional: true, Description: "'2048': Key size 2048bit; '4096': Key size 4096bit;",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"wipe": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Wipe SSH key",
			},
		},
	}
}
func resourceSshdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSshdRead(ctx, d, meta)
	}
	return diags
}

func resourceSshdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSshdRead(ctx, d, meta)
	}
	return diags
}
func resourceSshdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSshdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSshd(d *schema.ResourceData) edpt.Sshd {
	var ret edpt.Sshd
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.Generate = d.Get("generate").(int)
	ret.Inst.Load = d.Get("load").(int)
	ret.Inst.ReAddRsa = d.Get("re_add_rsa").(string)
	ret.Inst.Regenerate = d.Get("regenerate").(int)
	ret.Inst.Restart = d.Get("restart").(int)
	ret.Inst.Size = d.Get("size").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.Wipe = d.Get("wipe").(int)
	return ret
}
