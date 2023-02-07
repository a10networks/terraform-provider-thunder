package thunder

//Thunder resource SlbCommonConnRateLimitSrcIP

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonConnRateLimitSrcIP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbCommonConnRateLimitSrcIPCreate,
		UpdateContext: resourceSlbCommonConnRateLimitSrcIPUpdate,
		ReadContext:   resourceSlbCommonConnRateLimitSrcIPRead,
		DeleteContext: resourceSlbCommonConnRateLimitSrcIPDelete,
		Schema: map[string]*schema.Schema{
			"lock_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"exceed_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"limit_period": {
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

func resourceSlbCommonConnRateLimitSrcIPCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbCommonConnRateLimitSrcIP (Inside resourceSlbCommonConnRateLimitSrcIPCreate) ")
		name := d.Get("protocol").(string)
		data := dataToSlbCommonConnRateLimitSrcIP(d)
		logger.Println("[INFO] received formatted data from method data to SlbCommonConnRateLimitSrcIP --")
		d.SetId(name)
		err := go_thunder.PostSlbCommonConnRateLimitSrcIP(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbCommonConnRateLimitSrcIPRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbCommonConnRateLimitSrcIPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbCommonConnRateLimitSrcIP (Inside resourceSlbCommonConnRateLimitSrcIPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbCommonConnRateLimitSrcIP(client.Token, name, client.Host)
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

func resourceSlbCommonConnRateLimitSrcIPUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbCommonConnRateLimitSrcIP   (Inside resourceSlbCommonConnRateLimitSrcIPUpdate) ")
		name := d.Get("protocol").(string)
		data := dataToSlbCommonConnRateLimitSrcIP(d)
		logger.Println("[INFO] received formatted data from method data to SlbCommonConnRateLimitSrcIP ")
		d.SetId(name)
		err := go_thunder.PutSlbCommonConnRateLimitSrcIP(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbCommonConnRateLimitSrcIPRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbCommonConnRateLimitSrcIPDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbCommonConnRateLimitSrcIPDelete) " + name)
		err := go_thunder.DeleteSlbCommonConnRateLimitSrcIP(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbCommonConnRateLimitSrcIP(d *schema.ResourceData) go_thunder.CommonConnRateLimitSrcIP {
	var vc go_thunder.CommonConnRateLimitSrcIP
	var c go_thunder.CommonConnRateLimitSrcIPInstance

	c.Protocol = d.Get("protocol").(string)
	c.Log = d.Get("log").(int)
	c.LockOut = d.Get("lock_out").(int)
	c.LimitPeriod = d.Get("limit_period").(string)
	c.Limit = d.Get("limit").(int)
	c.ExceedAction = d.Get("exceed_action").(int)
	c.Shared = d.Get("shared").(int)

	vc.UUID = c
	return vc
}
