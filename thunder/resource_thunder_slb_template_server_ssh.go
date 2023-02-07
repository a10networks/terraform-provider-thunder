package thunder

//Thunder resource SlbTemplateServerSSH

import (
	"context"
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateServerSSH() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateServerSSHCreate,
		UpdateContext: resourceSlbTemplateServerSSHUpdate,
		ReadContext:   resourceSlbTemplateServerSSHRead,
		DeleteContext: resourceSlbTemplateServerSSHDelete,
		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateServerSSHCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateServerSSH (Inside resourceSlbTemplateServerSSHCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateServerSSH(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServerSSH --")
		d.SetId(name)
		err := go_thunder.PostSlbTemplateServerSSH(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateServerSSHRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateServerSSHRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTemplateServerSSH (Inside resourceSlbTemplateServerSSHRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateServerSSH(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbTemplateServerSSHUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateServerSSH   (Inside resourceSlbTemplateServerSSHUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateServerSSH(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServerSSH ")
		d.SetId(name)
		err := go_thunder.PutSlbTemplateServerSSH(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateServerSSHRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateServerSSHDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateServerSSHDelete) " + name)
		err := go_thunder.DeleteSlbTemplateServerSSH(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateServerSSH(d *schema.ResourceData) go_thunder.ServerSSH {
	var vc go_thunder.ServerSSH
	var c go_thunder.ServerSSHInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable2, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable2
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	c.Name = d.Get("name").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.UUID = c
	return vc
}
