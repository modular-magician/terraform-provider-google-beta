// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

package datalossprevention_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDeidentifyTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateBasicExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_deidentify_template.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "template_id"},
			},
		},
	})
}

func testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_deidentify_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Displayname"

	deidentify_config {
		info_type_transformations {
			transformations {
				info_types {
					name = "FIRST_NAME"
				}

				primitive_transformation {
					replace_with_info_type_config = true
				}
			}

			transformations {
				info_types {
					name = "PHONE_NUMBER"
				}
				info_types {
					name = "AGE"
				}

				primitive_transformation {
					replace_config {
						new_value {
							integer_value = 9
						}
					}
				}
			}

			transformations {
				info_types {
					name = "EMAIL_ADDRESS"
				}
				info_types {
					name = "LAST_NAME"
				}

				primitive_transformation {
					character_mask_config {
						masking_character = "X"
						number_to_mask = 4
						reverse_order = true
						characters_to_ignore {
							common_characters_to_ignore = "PUNCTUATION"
						}
					}
				}
			}

			transformations {
				info_types {
					name = "DATE_OF_BIRTH"
				}

				primitive_transformation {
					replace_config {
						new_value {
							date_value {
								year  = 2020
								month = 1
								day   = 1
							}
						}
					}
				}
			}

      transformations {
        info_types {
          name = "CREDIT_CARD_NUMBER"
        }

        primitive_transformation {
          crypto_deterministic_config {
            context {
              name = "sometweak"
            }
            crypto_key {
              transient {
                name = "beep"
              }
            }
            surrogate_info_type {
              name = "abc"
            }
          }
        }
      }
		}
	}
}
`, context)
}

func TestAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateSkipCharactersExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDeidentifyTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateSkipCharactersExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_deidentify_template.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "template_id"},
			},
		},
	})
}

func testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateSkipCharactersExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_deidentify_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Displayname"

	deidentify_config {
		info_type_transformations {
			transformations {
				info_types {
					name = "FIRST_NAME"
				}

				primitive_transformation {
					replace_with_info_type_config = true
				}
			}

			transformations {
				info_types {
					name = "PHONE_NUMBER"
				}
				info_types {
					name = "AGE"
				}

				primitive_transformation {
					replace_config {
						new_value {
							integer_value = 9
						}
					}
				}
			}

			transformations {
				info_types {
					name = "EMAIL_ADDRESS"
				}
				info_types {
					name = "LAST_NAME"
				}

				primitive_transformation {
					character_mask_config {
						masking_character = "X"
						number_to_mask = 4
						reverse_order = true
						characters_to_ignore {
							characters_to_skip = "@"
						}
					}
				}
			}

			transformations {
				info_types {
					name = "DATE_OF_BIRTH"
				}

				primitive_transformation {
					replace_config {
						new_value {
							date_value {
								year  = 2020
								month = 1
								day   = 1
							}
						}
					}
				}
			}

      transformations {
        info_types {
          name = "CREDIT_CARD_NUMBER"
        }

        primitive_transformation {
          crypto_deterministic_config {
            context {
              name = "sometweak"
            }
            crypto_key {
              transient {
                name = "beep"
              }
            }
            surrogate_info_type {
              name = "abc"
            }
          }
        }
      }
		}
	}
}
`, context)
}

func TestAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateImageTransformationsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDeidentifyTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateImageTransformationsExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_deidentify_template.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "template_id"},
			},
		},
	})
}

func testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateImageTransformationsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_deidentify_template" "basic" {
  parent = "projects/%{project}"
  description = "Description"
  display_name = "Displayname"
  
  deidentify_config {
    image_transformations {
      transforms {
        redaction_color {
          red = 0.5
          blue = 1
          green = 0.2
        }
        selected_info_types {
          info_types {
            name = "COLOR_INFO"
            version = "latest"
          }
        }
      }

      transforms {
        all_info_types {}
      }

      transforms {
        all_text {}
      }
    }
  }
}
`, context)
}

func TestAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateWithTemplateIdExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDeidentifyTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateWithTemplateIdExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_deidentify_template.with_template_id",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "template_id"},
			},
		},
	})
}

func testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateWithTemplateIdExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_deidentify_template" "with_template_id" {
  parent = "projects/%{project}"
  template_id = "tf-test-my-template%{random_suffix}"

  deidentify_config {
    info_type_transformations {
      transformations {
        info_types {
          name = "PHONE_NUMBER"
        }
        info_types {
          name = "AGE"
        }

        primitive_transformation {
          replace_config {
            new_value {
              integer_value = 9
            }
          }
        }
      }
    }
  }
}
`, context)
}

func testAccCheckDataLossPreventionDeidentifyTemplateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_loss_prevention_deidentify_template" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataLossPreventionBasePath}}{{parent}}/deidentifyTemplates/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("DataLossPreventionDeidentifyTemplate still exists at %s", url)
			}
		}

		return nil
	}
}
