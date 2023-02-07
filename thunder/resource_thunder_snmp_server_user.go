package thunder

//Thunder resource SnmpServerUser

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerUserCreate,
		UpdateContext: resourceSnmpServerUserUpdate,
		ReadContext:   resourceSnmpServerUserRead,
		DeleteContext: resourceSnmpServerUserDelete,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_val": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"encpasswd": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"passwd": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priv_pw_encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"v3": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pw_encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priv": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerUser (Inside resourceSnmpServerUserCreate) ")
		name1 := d.Get("username").(string)
		data := dataToSnmpServerUser(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerUser --")
		d.SetId(name1)
		err := go_thunder.PostSnmpServerUser(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerUserRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerUser (Inside resourceSnmpServerUserRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerUser(client.Token, name1, client.Host)
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

func resourceSnmpServerUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SnmpServerUser   (Inside resourceSnmpServerUserUpdate) ")
		data := dataToSnmpServerUser(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerUser ")
		err := go_thunder.PutSnmpServerUser(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerUserRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerUserDelete) " + name1)
		err := go_thunder.DeleteSnmpServerUser(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToSnmpServerUser(d *schema.ResourceData) go_thunder.SnmpServerUser {
	var vc go_thunder.SnmpServerUser
	var c go_thunder.SnmpServerUserInstance
	c.Username = d.Get("username").(string)
	c.AuthVal = d.Get("auth_val").(string)
	c.Group = d.Get("group").(string)
	c.Encpasswd = d.Get("encpasswd").(string)
	c.Passwd = d.Get("passwd").(string)
	c.PrivPwEncrypted = d.Get("priv_pw_encrypted").(string)
	c.V3 = d.Get("v3").(string)
	c.PwEncrypted = d.Get("pw_encrypted").(string)
	c.Priv = d.Get("priv").(string)

	vc.Username = c
	return vc
}
