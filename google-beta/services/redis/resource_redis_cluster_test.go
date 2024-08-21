// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package redis_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccRedisCluster_createClusterWithNodeType(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with replica count 1
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 3, preventDestroy: true, nodeType: "REDIS_STANDARD_SMALL", zoneDistributionMode: "MULTI_ZONE"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, preventDestroy: false, nodeType: "REDIS_STANDARD_SMALL", zoneDistributionMode: "MULTI_ZONE"}),
			},
		},
	})
}

// Validate zone distribution for the cluster.
func TestAccRedisCluster_createClusterWithZoneDistribution(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with replica count 1
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, preventDestroy: false, zoneDistributionMode: "SINGLE_ZONE", zone: "us-central1-b"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, preventDestroy: false, zoneDistributionMode: "SINGLE_ZONE", zone: "us-central1-b"}),
			},
		},
	})
}

// Validate that replica count is updated for the cluster
func TestAccRedisCluster_updateReplicaCount(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with replica count 1
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 3, preventDestroy: true, zoneDistributionMode: "MULTI_ZONE"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// update replica count to 2
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 2, shardCount: 3, preventDestroy: true, zoneDistributionMode: "MULTI_ZONE"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 3, preventDestroy: false, zoneDistributionMode: "MULTI_ZONE"}),
			},
			{
				// update replica count to 0
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, preventDestroy: true, zoneDistributionMode: "MULTI_ZONE"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, preventDestroy: false, zoneDistributionMode: "MULTI_ZONE"}),
			},
		},
	})
}

// Validate that shard count is updated for the cluster
func TestAccRedisCluster_updateShardCount(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with shard count 3
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 3, preventDestroy: true, zoneDistributionMode: "MULTI_ZONE"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// update shard count to 5
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 5, preventDestroy: true, zoneDistributionMode: "MULTI_ZONE"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 5, preventDestroy: false, zoneDistributionMode: "MULTI_ZONE"}),
			},
		},
	})
}

// Validate that redisConfigs is updated for the cluster
func TestAccRedisCluster_updateRedisConfigs(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster
				Config: createOrUpdateRedisCluster(&ClusterParams{
					name:                 name,
					shardCount:           3,
					zoneDistributionMode: "MULTI_ZONE",
					redisConfigs: map[string]string{
						"maxmemory-policy": "volatile-ttl",
					}}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// add a new redis config key-value pair and update existing redis config
				Config: createOrUpdateRedisCluster(&ClusterParams{
					name:                 name,
					shardCount:           3,
					zoneDistributionMode: "MULTI_ZONE",
					redisConfigs: map[string]string{
						"maxmemory-policy":  "allkeys-lru",
						"maxmemory-clients": "90%",
					}}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// remove all redis configs
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, shardCount: 3, zoneDistributionMode: "MULTI_ZONE"}),
			},
		},
	})
}

type ClusterParams struct {
	name                 string
	replicaCount         int
	shardCount           int
	preventDestroy       bool
	nodeType             string
	redisConfigs         map[string]string
	zoneDistributionMode string
	zone                 string
}

func createOrUpdateRedisCluster(params *ClusterParams) string {
	lifecycleBlock := ""
	if params.preventDestroy {
		lifecycleBlock = `
		lifecycle {
			prevent_destroy = true
		}`
	}
	var strBuilder strings.Builder
	for key, value := range params.redisConfigs {
		strBuilder.WriteString(fmt.Sprintf("%s =  \"%s\"\n", key, value))
	}

	zoneDistributionConfigBlock := ``
	if params.zoneDistributionMode != "" {
		zoneDistributionConfigBlock = fmt.Sprintf(`
		zone_distribution_config {
			mode = "%s"
			zone = "%s"
		}
		`, params.zoneDistributionMode, params.zone)
	}

	return fmt.Sprintf(`
resource "google_redis_cluster" "test" {
	provider = google-beta
	name           = "%s"
	replica_count = %d
	shard_count = %d
	node_type = "%s"
	region         = "us-central1"
	psc_configs {
			network = google_compute_network.producer_net.id
	}
	redis_configs = {
		%s
	}
  %s
	depends_on = [
			google_network_connectivity_service_connection_policy.default
		]
	%s
}

resource "google_network_connectivity_service_connection_policy" "default" {
	provider = google-beta
	name = "%s"
	location = "us-central1"
	service_class = "gcp-memorystore-redis"
	description   = "my basic service connection policy"
	network = google_compute_network.producer_net.id
	psc_config {
	subnetworks = [google_compute_subnetwork.producer_subnet.id]
	}
}

resource "google_compute_subnetwork" "producer_subnet" {
	provider      = google-beta
	name          = "%s"
	ip_cidr_range = "10.0.0.248/29"
	region        = "us-central1"
	network       = google_compute_network.producer_net.id
}

resource "google_compute_network" "producer_net" {
	provider                = google-beta
	name                    = "%s"
	auto_create_subnetworks = false
}
`, params.name, params.replicaCount, params.shardCount, params.nodeType, strBuilder.String(), zoneDistributionConfigBlock, lifecycleBlock, params.name, params.name, params.name)
}
