package thunder

//Thunder resource SnmpServerSNMPv3User

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerSNMPv3User() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerSNMPv3UserCreate,
		UpdateContext: resourceSnmpServerSNMPv3UserUpdate,
		ReadContext:   resourceSnmpServerSNMPv3UserRead,
		DeleteContext: resourceSnmpServerSNMPv3UserDelete,
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

func resourceSnmpServerSNMPv3UserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerSNMPv3User (Inside resourceSnmpServerSNMPv3UserCreate) ")
		name1 := d.Get("username").(string)
		data := dataToSnmpServerSNMPv3User(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerSNMPv3User --")
		d.SetId(name1)
		err := go_thunder.PostSnmpServerSNMPv3User(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerSNMPv3UserRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerSNMPv3UserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerSNMPv3User (Inside resourceSnmpServerSNMPv3UserRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerSNMPv3User(client.Token, name1, client.Host)
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

func resourceSnmpServerSNMPv3UserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SnmpServerSNMPv3User   (Inside resourceSnmpServerSNMPv3UserUpdate) ")
		data := dataToSnmpServerSNMPv3User(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerSNMPv3User ")
		err := go_thunder.PutSnmpServerSNMPv3User(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerSNMPv3UserRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerSNMPv3UserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerSNMPv3UserDelete) " + name1)
		err := go_thunder.DeleteSnmpServerSNMPv3User(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToSnmpServerSNMPv3User(d *schema.ResourceData) go_thunder.SnmpServerSNMPv3User {
	var vc go_thunder.SnmpServerSNMPv3User
	var c go_thunder.SnmpServerSNMPv3UserInstance
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
