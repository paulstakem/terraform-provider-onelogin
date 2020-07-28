package onelogin

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAppRule_crud(t *testing.T) {
	base := GetFixture("onelogin_saml_app_example.tf", t)
	update := GetFixture("onelogin_saml_app_updated_example.tf", t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { TestAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: base,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("onelogin_app_rules.test_a", "name", "first rule"),
					resource.TestCheckResourceAttr("onelogin_app_rules.test_a", "position", "1"),
					resource.TestCheckResourceAttr("onelogin_app_rules.test_b", "name", "second rule"),
					resource.TestCheckResourceAttr("onelogin_app_rules.test_b", "position", "2"),
				),
			},
			{
				Config: update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("onelogin_app_rules.test_a", "name", "updated first rule"),
					resource.TestCheckResourceAttr("onelogin_app_rules.test_a", "position", "2"),
					resource.TestCheckResourceAttr("onelogin_app_rules.test_b", "name", "updated second rule"),
					resource.TestCheckResourceAttr("onelogin_app_rules.test_b", "position", "1"),
				),
			},
		},
	})
}
