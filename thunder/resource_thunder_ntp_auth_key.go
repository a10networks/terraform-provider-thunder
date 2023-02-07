package thunder

//Thunder resource NtpAuthKey

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNtpAuthKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNtpAuthKeyCreate,
		UpdateContext: resourceNtpAuthKeyUpdate,
		ReadContext:   resourceNtpAuthKeyRead,
		DeleteContext: resourceNtpAuthKeyDelete,
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alg_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"asc_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"hex_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"hex_encrypted": {
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

func resourceNtpAuthKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating NtpAuthKey (Inside resourceNtpAuthKeyCreate) ")
		name1 := d.Get("key").(int)
		data := dataToNtpAuthKey(d)
		logger.Println("[INFO] received formatted data from method data to NtpAuthKey --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostNtpAuthKey(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceNtpAuthKeyRead(ctx, d, meta)

	}
	return diags
}

func resourceNtpAuthKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading NtpAuthKey (Inside resourceNtpAuthKeyRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetNtpAuthKey(client.Token, name1, client.Host)
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

func resourceNtpAuthKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying NtpAuthKey   (Inside resourceNtpAuthKeyUpdate) ")
		data := dataToNtpAuthKey(d)
		logger.Println("[INFO] received formatted data from method data to NtpAuthKey ")
		err := go_thunder.PutNtpAuthKey(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceNtpAuthKeyRead(ctx, d, meta)

	}
	return diags
}

func resourceNtpAuthKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceNtpAuthKeyDelete) " + name1)
		err := go_thunder.DeleteNtpAuthKey(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToNtpAuthKey(d *schema.ResourceData) go_thunder.NtpAuthKey {
	var vc go_thunder.NtpAuthKey
	var c go_thunder.NtpAuthKeyInstance
	c.Key = d.Get("key").(int)
	c.AlgType = d.Get("alg_type").(string)
	c.KeyType = d.Get("key_type").(string)
	c.AscKey = d.Get("asc_key").(string)
	c.Encrypted = d.Get("encrypted").(string)
	c.HexKey = d.Get("hex_key").(string)
	c.HexEncrypted = d.Get("hex_encrypted").(string)

	vc.Key = c
	return vc
}
