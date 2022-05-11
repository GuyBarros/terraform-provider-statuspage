package provider

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type StatusResponse struct {
	Page struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		URL       string `json:"url"`
		UpdatedAt string `json:"updated_at"`
	} `json:"page"`
	Status struct {
		Description string `json:"description"`
		Indicator   string `json:"indicator"`
	} `json:"status"`
}

func dataSourceStatus() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider scaffolding.",

		ReadContext: dataSourceStatusRead,

		Schema: map[string]*schema.Schema{
			"page_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"page_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"page_updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"indicator": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceStatusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {

	// Talk to the API
	// Thie status.hashicorp.com bit should be parameterised as part of the provider config
	resp, err := http.Get("https://status.hashicorp.com/api/v2/status.json")
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	var response StatusResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(response.Page.ID)

	if err := d.Set("page_name", response.Page.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("page_url", response.Page.URL); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("page_updated_at", response.Page.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", response.Status.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("indicator", response.Status.Indicator); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
