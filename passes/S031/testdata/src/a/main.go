package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure MaxItems"
		Computed: true,
		MaxItems: 1,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		MaxItems: 1,
		Optional: true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure MaxItems"
			Computed: true,
			MaxItems: 1,
		},
	}
}
