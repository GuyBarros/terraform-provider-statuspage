package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
type AutoGenerated struct {
	Page struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	Status struct {
		Description string `json:"description"`
		Indicator   string `json:"indicator"`
	} `json:"status"`
}
*/

func dataSourceStatus() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider scaffolding.",

		ReadContext: dataSourceStatusRead,

		Schema: map[string]*schema.Schema{
			"page_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"page_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"page_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"page_updated_at": {
				Type:         schema.TypeString,
				Computed:     true,
				ValidateFunc: validation.IsRFC3339Time,
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

func dataSourceStatusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}
