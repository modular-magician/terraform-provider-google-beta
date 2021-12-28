package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeBackendService_basic(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	extraCheckName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_basic(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_basicModified(
					serviceName, checkName, extraCheckName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withBackend(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	igName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	itName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withBackend(
					serviceName, igName, itName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withBackend(
					serviceName, igName, itName, checkName, 20),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withBackendAndMaxUtilization(t *testing.T) {
	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	igName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	itName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withBackend(
					serviceName, igName, itName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withBackendAndMaxUtilization(
					serviceName, igName, itName, checkName, 10),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccComputeBackendService_withBackendAndMaxUtilization(
					serviceName, igName, itName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withBackendAndIAP(t *testing.T) {
	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	igName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	itName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withBackendAndIAP(
					serviceName, igName, itName, checkName, 10),
			},
			{
				ResourceName:            "google_compute_backend_service.lipsum",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"iap.0.oauth2_client_secret"},
			},
			{
				Config: testAccComputeBackendService_withBackend(
					serviceName, igName, itName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_updatePreservesOptionalParameters(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withSessionAffinity(
					serviceName, checkName, "initial-description", "GENERATED_COOKIE"),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withSessionAffinity(
					serviceName, checkName, "updated-description", "GENERATED_COOKIE"),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withConnectionDraining(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withConnectionDraining(serviceName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withConnectionDrainingAndUpdate(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withConnectionDraining(serviceName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_basic(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withHttpsHealthCheck(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withHttpsHealthCheck(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withCdnPolicy(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withCdnPolicy(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withCdnPolicy2(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withSecurityPolicy(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	polName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withSecurityPolicy(serviceName, checkName, polName, "google_compute_security_policy.policy.self_link"),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withSecurityPolicy(serviceName, checkName, polName, "\"\""),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withCDNEnabled(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withCDNEnabled(
					serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withSessionAffinity(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withSessionAffinity(
					serviceName, checkName, "description", "CLIENT_IP"),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withSessionAffinity(
					serviceName, checkName, "description", "GENERATED_COOKIE"),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withAffinityCookieTtlSec(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withAffinityCookieTtlSec(
					serviceName, checkName, "description", "GENERATED_COOKIE", 300),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withMaxConnections(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	igName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	itName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withMaxConnections(
					serviceName, igName, itName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withMaxConnections(
					serviceName, igName, itName, checkName, 20),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withMaxConnectionsPerInstance(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	igName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	itName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withMaxConnectionsPerInstance(
					serviceName, igName, itName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withMaxConnectionsPerInstance(
					serviceName, igName, itName, checkName, 20),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withMaxRatePerEndpoint(t *testing.T) {
	t.Parallel()

	randSuffix := randString(t, 10)
	service := fmt.Sprintf("tf-test-%s", randSuffix)
	instance := fmt.Sprintf("tf-test-%s", randSuffix)
	neg := fmt.Sprintf("tf-test-%s", randSuffix)
	network := fmt.Sprintf("tf-test-%s", randSuffix)
	check := fmt.Sprintf("tf-test-%s", randSuffix)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withMaxRatePerEndpoint(
					service, instance, neg, network, check, 0.2),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withMaxRatePerEndpoint(
					service, instance, neg, network, check, 0.4),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withMaxConnectionsPerEndpoint(t *testing.T) {
	t.Parallel()

	randSuffix := randString(t, 10)
	service := fmt.Sprintf("tf-test-%s", randSuffix)
	instance := fmt.Sprintf("tf-test-%s", randSuffix)
	neg := fmt.Sprintf("tf-test-%s", randSuffix)
	network := fmt.Sprintf("tf-test-%s", randSuffix)
	check := fmt.Sprintf("tf-test-%s", randSuffix)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withMaxConnectionsPerEndpoint(
					service, instance, neg, network, check, 5),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withMaxConnectionsPerEndpoint(
					service, instance, neg, network, check, 10),
			},
			{
				ResourceName:      "google_compute_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withCustomHeaders(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withCustomHeaders(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_basic(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_internalLoadBalancing(t *testing.T) {
	// Instance template uses UniqueId in some cases
	skipIfVcr(t)
	t.Parallel()

	fr := fmt.Sprintf("forwardrule-test-%s", randString(t, 10))
	proxy := fmt.Sprintf("forwardrule-test-%s", randString(t, 10))
	backend := fmt.Sprintf("forwardrule-test-%s", randString(t, 10))
	hc := fmt.Sprintf("forwardrule-test-%s", randString(t, 10))
	urlmap := fmt.Sprintf("forwardrule-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_internalLoadBalancing(fr, proxy, backend, hc, urlmap),
			},
			{
				ResourceName:      "google_compute_backend_service.backend_service",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_withLogConfig(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", randString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_withLogConfig(serviceName, checkName, 0.7, true),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withLogConfig(serviceName, checkName, 0.4, true),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withLogConfig(serviceName, checkName, 0.4, false),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withLogConfig2(serviceName, checkName, false),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_withLogConfig(serviceName, checkName, 0.7, false),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_trafficDirectorUpdateBasic(t *testing.T) {
	t.Parallel()

	backendName := fmt.Sprintf("foo-%s", randString(t, 10))
	checkName := fmt.Sprintf("bar-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_trafficDirectorBasic(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_trafficDirectorUpdateBasic(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_trafficDirectorUpdateFull(t *testing.T) {
	t.Parallel()

	backendName := fmt.Sprintf("foo-%s", randString(t, 10))
	checkName := fmt.Sprintf("bar-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_trafficDirectorFull(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendService_trafficDirectorUpdateFull(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendService_regionNegBackend(t *testing.T) {
	t.Parallel()

	suffix := randString(t, 10)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_regionNegBackend(suffix),
			},
			{
				ResourceName:      "google_compute_backend_service.backend",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_trafficDirectorBasic(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  locality_lb_policy    = "RING_HASH"
  circuit_breakers {
    max_connections = 10
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 11
        nanos   = 1234
      }
      name = "mycookie"
    }
  }
  outlier_detection {
    consecutive_errors = 2
  }
}

resource "google_compute_health_check" "health_check" {
  name = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeBackendService_trafficDirectorUpdateBasic(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  locality_lb_policy    = "RANDOM"
  circuit_breakers {
    max_connections = 10
  }
  outlier_detection {
    consecutive_errors = 2
  }
}

resource "google_compute_health_check" "health_check" {
  name = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeBackendService_trafficDirectorFull(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  locality_lb_policy    = "MAGLEV"
  circuit_breakers {
    max_connections = 10
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 11
        nanos   = 1234
      }
      name = "mycookie"
    }
  }
  outlier_detection {
    consecutive_errors = 2
  }
}

resource "google_compute_health_check" "health_check" {
  name = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeBackendService_trafficDirectorUpdateFull(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  locality_lb_policy    = "MAGLEV"
  circuit_breakers {
    connect_timeout {
      seconds = 3
      nanos   = 4
    }
    max_connections             = 11
    max_requests_per_connection = 12
    max_pending_requests        = 13
    max_requests                = 14
    max_retries                 = 15
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 12
      }
      name = "mycookie2"
      path = "mycookie2/path"
    }
    minimum_ring_size = 16
  }
  outlier_detection {
    base_ejection_time {
      seconds = 0
      nanos   = 5
    }
    consecutive_errors                    = 1
    consecutive_gateway_failure           = 3
    enforcing_consecutive_errors          = 4
    enforcing_consecutive_gateway_failure = 5
    enforcing_success_rate                = 6
    interval {
      seconds = 7
    }
    max_ejection_percent        = 99
    success_rate_minimum_hosts  = 98
    success_rate_request_volume = 97
    success_rate_stdev_factor   = 1800
  }
}

resource "google_compute_health_check" "health_check" {
  name = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeBackendService_basic(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, checkName)
}

func testAccComputeBackendService_withCDNEnabled(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]
  enable_cdn    = true
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, checkName)
}

func testAccComputeBackendService_basicModified(serviceName, checkOne, checkTwo string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.one.self_link]
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_http_health_check" "one" {
  name               = "%s"
  request_path       = "/one"
  check_interval_sec = 30
  timeout_sec        = 30
}
`, serviceName, checkOne, checkTwo)
}

func testAccComputeBackendService_withBackend(
	serviceName, igName, itName, checkName string, timeout int64) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = %v

  backend {
    group = google_compute_instance_group_manager.foobar.instance_group
  }

  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "%s"
  version {
    instance_template = google_compute_instance_template.foobar.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-foobar"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "foobar" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.my_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_http_health_check" "default" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, timeout, igName, itName, checkName)
}

func testAccComputeBackendService_withBackendAndMaxUtilization(
	serviceName, igName, itName, checkName string, timeout int64) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = %v

  backend {
    group           = google_compute_instance_group_manager.foobar.instance_group
    max_utilization = 1.0
  }

  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "%s"
  version {
    instance_template = google_compute_instance_template.foobar.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-foobar"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "foobar" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.my_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_http_health_check" "default" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, timeout, igName, itName, checkName)
}

func testAccComputeBackendService_withBackendAndIAP(
	serviceName, igName, itName, checkName string, timeout int64) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = %v

  backend {
    group = google_compute_instance_group_manager.foobar.instance_group
  }

  iap {
    oauth2_client_id     = "test"
    oauth2_client_secret = "test"
  }

  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "%s"
  version {
    instance_template = google_compute_instance_template.foobar.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-foobar"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "foobar" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.my_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_http_health_check" "default" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, timeout, igName, itName, checkName)
}

func testAccComputeBackendService_withSessionAffinity(serviceName, checkName, description, affinityName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name             = "%s"
  description      = "%s"
  health_checks    = [google_compute_http_health_check.zero.self_link]
  session_affinity = "%s"
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, description, affinityName, checkName)
}

func testAccComputeBackendService_withAffinityCookieTtlSec(serviceName, checkName, description, affinityName string, affinityCookieTtlSec int64) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name                    = "%s"
  description             = "%s"
  health_checks           = [google_compute_http_health_check.zero.self_link]
  session_affinity        = "%s"
  affinity_cookie_ttl_sec = %v
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, description, affinityName, affinityCookieTtlSec, checkName)
}

func testAccComputeBackendService_withConnectionDraining(serviceName, checkName string, drainingTimeout int64) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name                            = "%s"
  health_checks                   = [google_compute_http_health_check.zero.self_link]
  connection_draining_timeout_sec = %v
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, drainingTimeout, checkName)
}

func testAccComputeBackendService_withHttpsHealthCheck(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_https_health_check.zero.self_link]
  protocol      = "HTTPS"
}

resource "google_compute_https_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, checkName)
}

func testAccComputeBackendService_withCdnPolicy(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]

  cdn_policy {
    negative_caching = false
    serve_while_stale = 0
    cache_key_policy {
      include_protocol       = true
      include_host           = true
      include_query_string   = true
      query_string_whitelist = ["foo", "bar"]
    }
  }
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, checkName)
}

func testAccComputeBackendService_withCdnPolicy2(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]

  cdn_policy {
    cache_key_policy {
      include_protocol       = true
      include_host           = true
      include_query_string   = true
      query_string_whitelist = ["foo", "bar"]
    }
  }
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, checkName)
}

func testAccComputeBackendService_withSecurityPolicy(serviceName, checkName, polName, polLink string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name            = "%s"
  health_checks   = [google_compute_http_health_check.zero.self_link]
  security_policy = %s
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_security_policy" "policy" {
  name        = "%s"
  description = "basic security policy"
}
`, serviceName, polLink, checkName, polName)
}

func testAccComputeBackendService_withMaxConnections(
	serviceName, igName, itName, checkName string, maxConnections int64) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  port_name   = "http"
  protocol    = "TCP"

  backend {
    group           = google_compute_instance_group_manager.foobar.instance_group
    max_connections = %v
  }

  health_checks = [google_compute_health_check.default.self_link]
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "%s"
  version {
    instance_template = google_compute_instance_template.foobar.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-foobar"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "foobar" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.my_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_health_check" "default" {
  name = "%s"
  tcp_health_check {
    port = "110"
  }
}
`, serviceName, maxConnections, igName, itName, checkName)
}

func testAccComputeBackendService_withMaxConnectionsPerInstance(
	serviceName, igName, itName, checkName string, maxConnectionsPerInstance int64) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  port_name   = "http"
  protocol    = "TCP"

  backend {
    group                        = google_compute_instance_group_manager.foobar.instance_group
    max_connections_per_instance = %v
  }

  health_checks = [google_compute_health_check.default.self_link]
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "%s"
  version {
    instance_template = google_compute_instance_template.foobar.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-foobar"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "foobar" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.my_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_health_check" "default" {
  name = "%s"
  tcp_health_check {
    port = "110"
  }
}
`, serviceName, maxConnectionsPerInstance, igName, itName, checkName)
}

func testAccComputeBackendService_withMaxConnectionsPerEndpoint(
	service, instance, neg, network, check string, maxConnections int64) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  port_name   = "http"
  protocol    = "TCP"

  backend {
    group                        = google_compute_network_endpoint_group.lb-neg.self_link
    balancing_mode               = "CONNECTION"
    max_connections_per_endpoint = %v
  }

  health_checks = [google_compute_health_check.default.self_link]
}

data "google_compute_image" "my_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_instance" "endpoint-instance" {
  name         = "%s"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.my_image.self_link
    }
  }

  network_interface {
    subnetwork = google_compute_subnetwork.default.self_link
    access_config {
      network_tier = "PREMIUM"
    }
  }
}

resource "google_compute_network_endpoint_group" "lb-neg" {
  name         = "%s"
  network      = google_compute_network.default.self_link
  subnetwork   = google_compute_subnetwork.default.self_link
  default_port = "90"
  zone         = "us-central1-a"
}

resource "google_compute_network_endpoint" "lb-endpoint" {
  network_endpoint_group = google_compute_network_endpoint_group.lb-neg.name

  instance   = google_compute_instance.endpoint-instance.name
  port       = google_compute_network_endpoint_group.lb-neg.default_port
  ip_address = google_compute_instance.endpoint-instance.network_interface[0].network_ip
}

resource "google_compute_network" "default" {
  name                    = "%s"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "%s"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.self_link
}

resource "google_compute_health_check" "default" {
  name = "%s"
  tcp_health_check {
    port = "110"
  }
}
`, service, maxConnections, instance, neg, network, network, check)
}

func testAccComputeBackendService_withMaxRatePerEndpoint(
	service, instance, neg, network, check string, maxRate float64) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  port_name   = "https"
  protocol    = "HTTPS"

  backend {
    group                 = google_compute_network_endpoint_group.lb-neg.self_link
    balancing_mode        = "RATE"
    max_rate_per_endpoint = %v
  }

  health_checks = [google_compute_health_check.default.self_link]
}

data "google_compute_image" "my_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_instance" "endpoint-instance" {
  name         = "%s"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.my_image.self_link
    }
  }

  network_interface {
    subnetwork = google_compute_subnetwork.default.self_link
    access_config {
      network_tier = "PREMIUM"
    }
  }
}

resource "google_compute_network_endpoint_group" "lb-neg" {
  name         = "%s"
  network      = google_compute_network.default.self_link
  subnetwork   = google_compute_subnetwork.default.self_link
  default_port = "90"
  zone         = "us-central1-a"
}

resource "google_compute_network_endpoint" "lb-endpoint" {
  network_endpoint_group = google_compute_network_endpoint_group.lb-neg.name

  instance   = google_compute_instance.endpoint-instance.name
  port       = google_compute_network_endpoint_group.lb-neg.default_port
  ip_address = google_compute_instance.endpoint-instance.network_interface[0].network_ip
}

resource "google_compute_network" "default" {
  name                    = "%s"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "%s"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.self_link
}

resource "google_compute_health_check" "default" {
  name                = "%s"
  check_interval_sec  = 3
  healthy_threshold   = 3
  timeout_sec         = 2
  unhealthy_threshold = 3
  https_health_check {
    port = "443"
  }
}
`, service, maxRate, instance, neg, network, network, check)
}

func testAccComputeBackendService_withCustomHeaders(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]

  custom_request_headers = ["Client-Region: {client_region}", "Client-Rtt: {client_rtt_msec}"]
  custom_response_headers = ["X-Cache-Hit: {cdn_cache_status}", "X-Cache-Id: {cdn_cache_id}"]
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, checkName)
}

func testAccComputeBackendService_internalLoadBalancing(fr, proxy, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
resource "google_compute_global_forwarding_rule" "forwarding_rule" {
  name                  = "%s"
  target                = google_compute_target_http_proxy.default.self_link
  port_range            = "80"
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  ip_address            = "0.0.0.0"
}

resource "google_compute_target_http_proxy" "default" {
  name        = "%s"
  description = "a description"
  url_map     = google_compute_url_map.default.self_link
  proxy_bind  = true
}

resource "google_compute_backend_service" "backend_service" {
  name                  = "%s"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"

  backend {
    group                 = google_compute_instance_group_manager.foobar.instance_group
    balancing_mode        = "RATE"
    capacity_scaler       = 0.4
    max_rate_per_instance = 50
  }

  health_checks = [google_compute_health_check.default.self_link]
}

resource "google_compute_health_check" "default" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_url_map" "default" {
  name            = "%s"
  description     = "a description"
  default_service = google_compute_backend_service.backend_service.self_link

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.backend_service.self_link

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.backend_service.self_link
    }
  }
}

data "google_compute_image" "debian_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "tf-test-igm-internal"
  version {
    instance_template = google_compute_instance_template.foobar.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-foobar"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "foobar" {
  name_prefix  = "tf-test-"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.debian_image.self_link
    auto_delete  = true
    boot         = true
  }
}
`, fr, proxy, backend, hc, urlmap)
}

func testAccComputeBackendService_withLogConfig(serviceName, checkName string, sampleRate float64, enabled bool) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]

  log_config {
    enable      = %t
    sample_rate = %v
  }
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, enabled, sampleRate, checkName)
}

func testAccComputeBackendService_withLogConfig2(serviceName, checkName string, enabled bool) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]

  log_config {
	enable      = %t
  }
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, serviceName, enabled, checkName)
}

func testAccComputeBackendService_regionNegBackend(suffix string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "backend" {
  name                            = "tf-test-backend%s"
  enable_cdn                      = true
  connection_draining_timeout_sec = 10

  backend {
    group = google_compute_region_network_endpoint_group.cloudrun_neg.id
  }
}

resource "google_compute_region_network_endpoint_group" "cloudrun_neg" {
  name                  = "tf-test-neg%s"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  cloud_run {
    service = google_cloud_run_service.cloudrun_neg.name
  }
}

resource "google_cloud_run_service" "cloudrun_neg" {
  name     = "tf-test-cr%s"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

`, suffix, suffix, suffix)
}
