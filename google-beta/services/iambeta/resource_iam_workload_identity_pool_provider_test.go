// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package iambeta_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIAMBetaWorkloadIdentityPoolProvider_aws(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_aws_full(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_aws_enabled(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_aws_basic(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIAMBetaWorkloadIdentityPoolProvider_oidc(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_oidc_full(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_oidc_update(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_oidc_basic(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolProvider_aws_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "AWS identity pool provider for automated test"
  disabled                           = true
  attribute_condition                = "attribute.aws_role==\"arn:aws:sts::999999999999:assumed-role/stack-eu-central-1-lambdaRole\""
  attribute_mapping                  = {
    "google.subject"        = "assertion.arn"
    "attribute.aws_account" = "assertion.account"
    "attribute.environment" = "assertion.arn.contains(\":instance-profile/Production\") ? \"prod\" : \"test\""
  }
  aws {
    account_id = "999999999999"
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_aws_enabled(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "AWS identity pool provider for automated test"
  disabled                           = false
  attribute_condition                = "attribute.aws_role==\"arn:aws:sts::999999999999:assumed-role/stack-eu-central-1-lambdaRole\""
  attribute_mapping                  = {
    "google.subject"        = "assertion.arn"
    "attribute.aws_account" = "assertion.account"
    "attribute.environment" = "assertion.arn.contains(\":instance-profile/Production\") ? \"prod\" : \"test\""
  }
  aws {
    account_id = "999999999999"
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_oidc_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "OIDC identity pool provider for automated test"
  disabled                           = true
  attribute_condition                = "\"e968c2ef-047c-498d-8d79-16ca1b61e77e\" in assertion.groups"
  attribute_mapping                  = {
    "google.subject"                  = "\"azure::\" + assertion.tid + \"::\" + assertion.sub"
    "attribute.tid"                   = "assertion.tid"
    "attribute.managed_identity_name" = <<EOT
      {
        "8bb39bdb-1cc5-4447-b7db-a19e920eb111":"workload1",
        "55d36609-9bcf-48e0-a366-a3cf19027d2a":"workload2"
      }[assertion.oid]
EOT
  }
  oidc {
    allowed_audiences = ["https://example.com/gcp-oidc-federation", "example.com/gcp-oidc-federation"]
    issuer_uri        = "https://sts.windows.net/azure-tenant-id-full"
    jwks_json         = "{\"keys\":[{\"kty\":\"RSA\",\"alg\":\"RS256\",\"kid\":\"sif0AR-F6MuvksAyAOv-Pds08Bcf2eUMlxE30NofddA\",\"use\":\"sig\",\"e\":\"AQAB\",\"n\":\"ylH1Chl1tpfti3lh51E1g5dPogzXDaQseqjsefGLknaNl5W6Wd4frBhHyE2t41Q5zgz_Ll0-NvWm0FlaG6brhrN9QZu6sJP1bM8WPfJVPgXOanxi7d7TXCkeNubGeiLTf5R3UXtS9Lm_guemU7MxDjDTelxnlgGCihOVTcL526suNJUdfXtpwUsvdU6_ZnAp9IpsuYjCtwPm9hPumlcZGMbxstdh07O4y4O90cVQClJOKSGQjAUCKJWXIQ0cqffGS_HuS_725CPzQ85SzYZzaNpgfhAER7kx_9P16ARM3BJz0PI5fe2hECE61J4GYU_BY43sxDfs7HyJpEXKLU9eWw\"}]}"
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_oidc_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "OIDC identity pool provider for automated test"
  disabled                           = true
  attribute_condition                = "\"e968c2ef-047c-498d-8d79-16ca1b61e77e\" in assertion.groups"
  attribute_mapping                  = {
    "google.subject"                  = "\"azure::\" + assertion.tid + \"::\" + assertion.sub"
    "attribute.tid"                   = "assertion.tid"
    "attribute.managed_identity_name" = <<EOT
      {
        "8bb39bdb-1cc5-4447-b7db-a19e920eb111":"workload1",
        "55d36609-9bcf-48e0-a366-a3cf19027d2a":"workload2"
      }[assertion.oid]
EOT
  }
  oidc {
    allowed_audiences = ["https://example.com/gcp-oidc-federation-update", "example.com/gcp-oidc-federation-update"]
    issuer_uri        = "https://sts.windows.net/azure-tenant-id-update"
    jwks_json         = "{\"keys\":[{\"kty\":\"RSA\",\"alg\":\"RS256\",\"kid\":\"F6MuvksAyAOv-Pds08Bcf2eUMlxE30NofddA\",\"use\":\"sig\",\"e\":\"AQAB\",\"n\":\"ylH1Chl1tpfti3lh51E1g5dPogzXDaQseqjsefGLknaNl5W6Wd4frBhHyE2t41Q5zgz_Ll0-NvWm0FlaG6brhrN9QZu6sJP1bM8WPfJVPgXOanxi7d7TXCkeNubGeiLTf5R3UXtS9Lm_guemU7MxDjDTelxnlgGCihOVTcL526suNJUdfXtpwUsvdU6_ZnAp9IpsuYjCtwPm9hPumlcZGMbxstdh07O4y4O90cVQClJOKSGQjAUCKJWXIQ0cqffGS_HuS_725CPzQ85SzYZzaNpgfhAER7kx_9P16ARM3BJz0PI5fe2hECE61J4GYU_BY43sxDfs7HyJpEXKLU9eWw\"}]}"

  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_aws_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  aws {
    account_id = "999999999999"
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_oidc_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  attribute_mapping                  = {
	"google.subject"                  = "assertion.sub"
  }
  oidc {
    allowed_audiences = ["https://example.com/gcp-oidc-federation", "example.com/gcp-oidc-federation"]
    issuer_uri        = "https://sts.windows.net/azure-tenant-id-full"
  }
}
`, context)
}

func TestAccIAMBetaWorkloadIdentityPoolProvider_x509(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_x509_full(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_provider_id"},
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_x509_update(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_provider_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolProvider_x509_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "example" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "tf-test-example-prvdr%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "X.509 identity pool provider for automated test"
  disabled                           = true
  attribute_mapping                  = {
    "google.subject"        = "assertion.subject.dn.cn"
  }
  x509 {
    trust_store {
        trust_anchors {
            pem_certificate = "-----BEGIN CERTIFICATE-----\nMIIDtjCCAp6gAwIBAgIIMu9pvsFWIUcwDQYJKoZIhvcNAQELBQAwXDEXMBUGA1UE\nChMOR29vZ2xlIFRFU1RJTkcxHDAaBgNVBAsTE0dvb2dsZSBURVNUSU5HIHVuaXQx\nIzAhBgNVBAMMGkdvb2dsZSAqKlRlc3RpbmcqKiBSb290IENBMCAXDTIwMDEwMTAw\nMDAwMFoYDzIxMzAwMTAxMDAwMDAwWjBcMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElO\nRzEcMBoGA1UECxMTR29vZ2xlIFRFU1RJTkcgdW5pdDEjMCEGA1UEAwwaR29vZ2xl\nICoqVGVzdGluZyoqIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK\nAoIBAQDHqdf17AA4iwl+HGtTD/qJJv9XZTjzjHiaRYWuGQ3iddf1LhYklLqHvpqO\nCBlhkg6NIZmoDPKicEi8pfGSzp6btcElyrr1ekECk5jEBcdl6tX/gTSYv7v1h9Dk\nSJDHBAnoSsVW0/PNwI+YGEE2kizGMeg1moXlTHEB+yeGCik/+4eVRas/+wrlrTE5\nlMFMq8WhdnBx6udcc/BEauvlTybHaN3rJUVpVrJWeVoPGGtXR6MJrdbVScn9GYOL\nP6sXMPTr+pTY6ebrjs7K/wTVQqLJ1zArCvAEH8FvuAhiI19yM1uBRAds/fJbSUvF\nsEiIlDC4y7DAK7yWHUAjQIBV1xhDAgMBAAGjejB4MA4GA1UdDwEB/wQEAwICBDAd\nBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDwYDVR0TAQH/BAUwAwEB/zAZ\nBgNVHQ4EEgQQxEHfoOjPZleWgcpHJnyE3TAbBgNVHSMEFDASgBDEQd+g6M9mV5aB\nykcmfITdMA0GCSqGSIb3DQEBCwUAA4IBAQBcECwxgaK4ZgYO7guayK5QRpTb3Y6z\npkmSOkj9h1+HF5Ch/o5FiweJi8k28h9Mfz//gKU2cXXWfXuY81CFEBstjw7jYt3+\nd4owS5sYKu5WswGj4jsQXRrt0tO0+0UngP560eggFCyLChG75lp54r92hTPk4fY4\nvarURZH7X0jg9W7MO6E6/HmKdMIxanuWdkbpPe6kj5I7SNvVhqsncoga8iZJ6QK/\nrsC2LTax4dxUUSkP5vmwwiYbgXYbk9JcXI+OyVcMvVxN/akI9/JgYHOol6NdChTB\nwU0yjNb9B5TrgtP1/fs+LugINrN1R1hgVHDVlE4mwHej+5XL6m9xAxkc\n-----END CERTIFICATE-----"
        }
        intermediate_cas {
            pem_certificate = "-----BEGIN CERTIFICATE-----\nMIIDvjCCAqagAwIBAgIIXHIdYNfGCLMwDQYJKoZIhvcNAQELBQAwXDEXMBUGA1UE\nChMOR29vZ2xlIFRFU1RJTkcxHDAaBgNVBAsTE0dvb2dsZSBURVNUSU5HIHVuaXQx\nIzAhBgNVBAMMGkdvb2dsZSAqKlRlc3RpbmcqKiBSb290IENBMCAXDTIwMDEwMTAw\nMDAwMFoYDzIxMzAwMTAxMDAwMDAwWjBkMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElO\nRzEcMBoGA1UECxMTR29vZ2xlIFRFU1RJTkcgdW5pdDErMCkGA1UEAwwiR29vZ2xl\nICoqVGVzdGluZyoqIEludGVybWVkaWF0ZSBDQTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBALVSyoOVAMfOYf10VxNwmRs/afgbkYJ5gPXzn3aetEtFrcqD\niG8LdgniLhHcp3c/O1U0EKANMLWyMlHP0KC4wjMrYXS8doQ7B6kGXj070hRSHN7a\ncF0ImSRw3i9idiiSgOIJzlbYXeh8NLqDAYESyWFht7RRdbCJx1v2U9T8F5QVeT96\nDw3exSKQAIdL2J8ol9kgJNINimd7GxOQ3f1+vDOBiAtAj4zWCjjqdNqh5ivyrTu2\n8J/umM35wtHS6iX6GvnwqTsRy8zS/KeN+Dq6PEk/j04mrOG3w82SFp+IvTNH6S/D\nxRyiE5yoAfgfunKotV34JVr/INH2RsEgbWtwK5ECAwEAAaN6MHgwDgYDVR0PAQH/\nBAQDAgIEMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAPBgNVHRMBAf8E\nBTADAQH/MBkGA1UdDgQSBBB0QukRw23faHpDoa0xl+L+MBsGA1UdIwQUMBKAEMRB\n36Doz2ZXloHKRyZ8hN0wDQYJKoZIhvcNAQELBQADggEBAIPBiuirnZbv4JWDiAIX\nvCE5pbwej0aKFWEDV2Z8lY0RFPt1CXrDJJL91MXHZ1yviUSJINJErJn7wyGV2bm/\nN7DUpRE0g9IMgEank64UUl+OyQTXd0LIsjlqWA6Sj/hDZUdw6mi9a98ENUr6CiEC\ntOxpGF9kj4G4WcnyvPP/phs1b8cAfQ+tPurrDRBAdeoQIj756QL7fvMijKNdG5Ke\nCURu9L4BZmCeuy/3v2C2XjiFZHx4cZDOHJizrx04GqzV5PSXw5OYZiXfn5WPGMsy\nh7ufkQJKJcpv2t3M0gyc1omD0xWkxCl4dTdVef9HrboJnZkUrma509fpVL6F8I2J\nkt8=\n-----END CERTIFICATE-----"
        }
    }
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_x509_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "example" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "tf-test-example-prvdr%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "X.509 identity pool provider for automated test"
  disabled                           = true
  attribute_mapping                  = {
    "google.subject"        = "assertion.subject.dn.cn"
  }
  x509 {
    trust_store {
        trust_anchors {
            pem_certificate = "-----BEGIN CERTIFICATE-----\nMIIDtjCCAp6gAwIBAgIIMu9pvsFWIUcwDQYJKoZIhvcNAQELBQAwXDEXMBUGA1UE\nChMOR29vZ2xlIFRFU1RJTkcxHDAaBgNVBAsTE0dvb2dsZSBURVNUSU5HIHVuaXQx\nIzAhBgNVBAMMGkdvb2dsZSAqKlRlc3RpbmcqKiBSb290IENBMCAXDTIwMDEwMTAw\nMDAwMFoYDzIxMzAwMTAxMDAwMDAwWjBcMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElO\nRzEcMBoGA1UECxMTR29vZ2xlIFRFU1RJTkcgdW5pdDEjMCEGA1UEAwwaR29vZ2xl\nICoqVGVzdGluZyoqIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK\nAoIBAQDHqdf17AA4iwl+HGtTD/qJJv9XZTjzjHiaRYWuGQ3iddf1LhYklLqHvpqO\nCBlhkg6NIZmoDPKicEi8pfGSzp6btcElyrr1ekECk5jEBcdl6tX/gTSYv7v1h9Dk\nSJDHBAnoSsVW0/PNwI+YGEE2kizGMeg1moXlTHEB+yeGCik/+4eVRas/+wrlrTE5\nlMFMq8WhdnBx6udcc/BEauvlTybHaN3rJUVpVrJWeVoPGGtXR6MJrdbVScn9GYOL\nP6sXMPTr+pTY6ebrjs7K/wTVQqLJ1zArCvAEH8FvuAhiI19yM1uBRAds/fJbSUvF\nsEiIlDC4y7DAK7yWHUAjQIBV1xhDAgMBAAGjejB4MA4GA1UdDwEB/wQEAwICBDAd\nBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDwYDVR0TAQH/BAUwAwEB/zAZ\nBgNVHQ4EEgQQxEHfoOjPZleWgcpHJnyE3TAbBgNVHSMEFDASgBDEQd+g6M9mV5aB\nykcmfITdMA0GCSqGSIb3DQEBCwUAA4IBAQBcECwxgaK4ZgYO7guayK5QRpTb3Y6z\npkmSOkj9h1+HF5Ch/o5FiweJi8k28h9Mfz//gKU2cXXWfXuY81CFEBstjw7jYt3+\nd4owS5sYKu5WswGj4jsQXRrt0tO0+0UngP560eggFCyLChG75lp54r92hTPk4fY4\nvarURZH7X0jg9W7MO6E6/HmKdMIxanuWdkbpPe6kj5I7SNvVhqsncoga8iZJ6QK/\nrsC2LTax4dxUUSkP5vmwwiYbgXYbk9JcXI+OyVcMvVxN/akI9/JgYHOol6NdChTB\nwU0yjNb9B5TrgtP1/fs+LugINrN1R1hgVHDVlE4mwHej+5XL6m9xAxkc\n-----END CERTIFICATE-----"
        }
        trust_anchors {
            pem_certificate = "-----BEGIN CERTIFICATE-----\nMIIDvjCCAqagAwIBAgIIXHIdYNfGCLMwDQYJKoZIhvcNAQELBQAwXDEXMBUGA1UE\nChMOR29vZ2xlIFRFU1RJTkcxHDAaBgNVBAsTE0dvb2dsZSBURVNUSU5HIHVuaXQx\nIzAhBgNVBAMMGkdvb2dsZSAqKlRlc3RpbmcqKiBSb290IENBMCAXDTIwMDEwMTAw\nMDAwMFoYDzIxMzAwMTAxMDAwMDAwWjBkMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElO\nRzEcMBoGA1UECxMTR29vZ2xlIFRFU1RJTkcgdW5pdDErMCkGA1UEAwwiR29vZ2xl\nICoqVGVzdGluZyoqIEludGVybWVkaWF0ZSBDQTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBALVSyoOVAMfOYf10VxNwmRs/afgbkYJ5gPXzn3aetEtFrcqD\niG8LdgniLhHcp3c/O1U0EKANMLWyMlHP0KC4wjMrYXS8doQ7B6kGXj070hRSHN7a\ncF0ImSRw3i9idiiSgOIJzlbYXeh8NLqDAYESyWFht7RRdbCJx1v2U9T8F5QVeT96\nDw3exSKQAIdL2J8ol9kgJNINimd7GxOQ3f1+vDOBiAtAj4zWCjjqdNqh5ivyrTu2\n8J/umM35wtHS6iX6GvnwqTsRy8zS/KeN+Dq6PEk/j04mrOG3w82SFp+IvTNH6S/D\nxRyiE5yoAfgfunKotV34JVr/INH2RsEgbWtwK5ECAwEAAaN6MHgwDgYDVR0PAQH/\nBAQDAgIEMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAPBgNVHRMBAf8E\nBTADAQH/MBkGA1UdDgQSBBB0QukRw23faHpDoa0xl+L+MBsGA1UdIwQUMBKAEMRB\n36Doz2ZXloHKRyZ8hN0wDQYJKoZIhvcNAQELBQADggEBAIPBiuirnZbv4JWDiAIX\nvCE5pbwej0aKFWEDV2Z8lY0RFPt1CXrDJJL91MXHZ1yviUSJINJErJn7wyGV2bm/\nN7DUpRE0g9IMgEank64UUl+OyQTXd0LIsjlqWA6Sj/hDZUdw6mi9a98ENUr6CiEC\ntOxpGF9kj4G4WcnyvPP/phs1b8cAfQ+tPurrDRBAdeoQIj756QL7fvMijKNdG5Ke\nCURu9L4BZmCeuy/3v2C2XjiFZHx4cZDOHJizrx04GqzV5PSXw5OYZiXfn5WPGMsy\nh7ufkQJKJcpv2t3M0gyc1omD0xWkxCl4dTdVef9HrboJnZkUrma509fpVL6F8I2J\nkt8=\n-----END CERTIFICATE-----"
        }
    }
  }
}
`, context)
}
