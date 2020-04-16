// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccTPUNode_tpuNodeBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTPUNodeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTPUNode_tpuNodeBasicExample(context),
			},
			{
				ResourceName:            "google_tpu_node.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone"},
			},
		},
	})
}

func testAccTPUNode_tpuNodeBasicExample(context map[string]interface{}) string {
	return Nprintf(`

data "google_tpu_tensorflow_versions" "available" {
}

resource "google_tpu_node" "tpu" {
  name = "tf-test-test-tpu%{random_suffix}"
  zone = "us-central1-b"

  accelerator_type   = "v3-8"
  tensorflow_version = data.google_tpu_tensorflow_versions.available.versions[0]
  cidr_block         = "10.2.0.0/29"
}
`, context)
}

func TestAccTPUNode_tpuNodeFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTPUNodeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTPUNode_tpuNodeFullExample(context),
			},
			{
				ResourceName:            "google_tpu_node.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone"},
			},
		},
	})
}

func testAccTPUNode_tpuNodeFullExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_tpu_tensorflow_versions" "available" {
}


resource "google_tpu_node" "tpu" {
  name = "tf-test-test-tpu%{random_suffix}"
  zone = "us-central1-b"

  accelerator_type = "v3-8"

  cidr_block         = "10.3.0.0/29"
  tensorflow_version = data.google_tpu_tensorflow_versions.available.versions[0]

  description = "Terraform Google Provider test TPU"
  network = "default"

  labels = {
    foo = "bar"
  }

  scheduling_config {
    preemptible = true
  }
}
`, context)
}

func testAccCheckTPUNodeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_tpu_node" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{TPUBasePath}}projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("TPUNode still exists at %s", url)
			}
		}

		return nil
	}
}
