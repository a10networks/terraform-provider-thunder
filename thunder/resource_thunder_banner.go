package thunder

//Thunder resource Banner

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBanner() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBannerCreate,
		UpdateContext: resourceBannerUpdate,
		ReadContext:   resourceBannerRead,
		DeleteContext: resourceBannerDelete,
		Schema: map[string]*schema.Schema{
			"exec_banner_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exec": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exec_banner": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"login_banner_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"login": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"login_banner": {
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
		},
	}
}

func resourceBannerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating Banner (Inside resourceBannerCreate) ")

		data := dataToBanner(d)
		logger.Println("[INFO] received formatted data from method data to Banner --")
		d.SetId("1")
		err := go_thunder.PostBanner(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceBannerRead(ctx, d, meta)

	}
	return diags
}

func resourceBannerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading Banner (Inside resourceBannerRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetBanner(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return diags
}

func resourceBannerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Updating Banner (Inside resourceBannerCreate) ")

		data := dataToBanner(d)
		logger.Println("[INFO] received formatted data from method data to Banner --")
		d.SetId("1")
		err := go_thunder.PutBanner(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceBannerRead(ctx, d, meta)

	}
	return diags
}

func resourceBannerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Deleting Banner (Inside resourceBannerDelete)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		err := go_thunder.DeleteBanner(client.Token, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete banner")
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToBanner(d *schema.ResourceData) go_thunder.Banner {
	var vc go_thunder.Banner
	var c go_thunder.BannerInstance

	var obj1 go_thunder.BannerInstanceExecBannerCfg
	prefix1 := "exec_banner_cfg.0."
	obj1.BannerInstanceExecBannerCfgExec = d.Get(prefix1 + "exec").(int)
	obj1.BannerInstanceExecBannerCfgExecBanner = d.Get(prefix1 + "exec_banner").(string)

	c.BannerInstanceExecBannerCfgExec = obj1

	var obj2 go_thunder.BannerInstanceLoginBannerCfg
	prefix2 := "login_banner_cfg.0."
	obj2.BannerInstanceLoginBannerCfgLogin = d.Get(prefix2 + "login").(int)
	obj2.BannerInstanceLoginBannerCfgLoginBanner = d.Get(prefix2 + "login_banner").(string)

	c.BannerInstanceLoginBannerCfgLogin = obj2

	vc.BannerInstanceExecBannerCfg = c
	return vc
}
