package thunder

//Thunder resource TemplateDBLB

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateDBLB() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateDBLBCreate,
		UpdateContext: resourceTemplateDBLBUpdate,
		ReadContext:   resourceTemplateDBLBRead,
		DeleteContext: resourceTemplateDBLBDelete,
		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"calc_sha1": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sha1_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"server_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"class_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTemplateDBLBCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateDBLB (Inside resourceTemplateDBLBCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateDBLB(d)
		logger.Println("[INFO] received formatted data from method data to TemplateDBLB --")
		d.SetId(name)
		err := go_thunder.PostTemplateDBLB(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateDBLBRead(ctx, d, meta)

	}
	return diags
}

func resourceTemplateDBLBRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading TemplateDBLB (Inside resourceTemplateDBLBRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetTemplateDBLB(client.Token, name, client.Host)
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

func resourceTemplateDBLBUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateDBLB   (Inside resourceTemplateDBLBUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateDBLB(d)
		logger.Println("[INFO] received formatted data from method data to TemplateDBLB ")
		d.SetId(name)
		err := go_thunder.PutTemplateDBLB(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateDBLBRead(ctx, d, meta)

	}
	return diags
}

func resourceTemplateDBLBDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateDBLBDelete) " + name)
		err := go_thunder.DeleteTemplateDBLB(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateDBLB(d *schema.ResourceData) go_thunder.DBLB {
	var vc go_thunder.DBLB
	var c go_thunder.DblbInstance
	var CalcSha1Obj go_thunder.CalcSha1
	c.Name = d.Get("name").(string)
	c.ServerVersion = d.Get("server_version").(string)
	CalcSha1Obj.Sha1Value = d.Get("calc_sha1.0.sha1_value").(string)
	c.Sha1Value = CalcSha1Obj
	c.UserTag = d.Get("user_tag").(string)
	c.ClassList = d.Get("class_list").(string)

	vc.UUID = c
	return vc
}
