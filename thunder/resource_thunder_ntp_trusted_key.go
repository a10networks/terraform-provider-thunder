package thunder

//Thunder resource NtpTrustedKey

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNtpTrustedKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNtpTrustedKeyCreate,
		UpdateContext: resourceNtpTrustedKeyUpdate,
		ReadContext:   resourceNtpTrustedKeyRead,
		DeleteContext: resourceNtpTrustedKeyDelete,
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeInt,
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

func resourceNtpTrustedKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating NtpTrustedKey (Inside resourceNtpTrustedKeyCreate) ")
		name1 := d.Get("key").(int)
		data := dataToNtpTrustedKey(d)
		logger.Println("[INFO] received formatted data from method data to NtpTrustedKey --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostNtpTrustedKey(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceNtpTrustedKeyRead(ctx, d, meta)

	}
	return diags
}

func resourceNtpTrustedKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading NtpTrustedKey (Inside resourceNtpTrustedKeyRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetNtpTrustedKey(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceNtpTrustedKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying NtpTrustedKey   (Inside resourceNtpTrustedKeyUpdate) ")
		data := dataToNtpTrustedKey(d)
		logger.Println("[INFO] received formatted data from method data to NtpTrustedKey ")
		err := go_thunder.PutNtpTrustedKey(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceNtpTrustedKeyRead(ctx, d, meta)

	}
	return diags
}

func resourceNtpTrustedKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceNtpTrustedKeyDelete) " + name1)
		err := go_thunder.DeleteNtpTrustedKey(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToNtpTrustedKey(d *schema.ResourceData) go_thunder.NtpTrustedKey {
	var vc go_thunder.NtpTrustedKey
	var c go_thunder.NtpTrustedKeyInstance
	c.Key = d.Get("key").(int)

	vc.Key = c
	return vc
}
