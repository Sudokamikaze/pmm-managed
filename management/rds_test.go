package management

import (
	"fmt"
	"os"
	"testing"

	"github.com/AlekSi/pointer"
	inventoryClient "github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/agents"
	"github.com/percona/pmm/api/inventorypb/json/client/nodes"
	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/rds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestRDSDiscovery(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		accessKey, secretKey := os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY")
		if accessKey == "" || secretKey == "" {
			// TODO remove skip once secrets are added
			t.Skip("Environment variables AWS_ACCESS_KEY / AWS_SECRET_KEY are not defined, skipping test")
		}

		params := &rds.DiscoverRDSParams{
			Body: rds.DiscoverRDSBody{
				AWSAccessKey: accessKey,
				AWSSecretKey: secretKey,
			},
			Context: pmmapitests.Context,
		}
		discoverOK, err := client.Default.RDS.DiscoverRDS(params)
		require.NoError(t, err)
		require.NotNil(t, discoverOK.Payload)
		assert.NotEmpty(t, discoverOK.Payload.RDSInstances)

		// TODO Better tests: https://jira.percona.com/browse/PMM-4896
	})
}

func TestAddRds(t *testing.T) {
	t.Run("BasicAddRDS", func(t *testing.T) {
		params := &rds.AddRDSParams{
			Body: rds.AddRDSBody{
				Region:                    "region",
				Az:                        "az",
				InstanceID:                "d752f1a9-31c9-4b8c-bb2d-d26bc000001",
				NodeModel:                 "some-model",
				Address:                   "some.example.rds",
				Port:                      3306,
				Engine:                    pointer.ToString("DISCOVER_RDS_MYSQL"),
				NodeName:                  "some-node-name-000001",
				ServiceName:               "test-add-rds-service000001",
				Environment:               "some-env",
				Cluster:                   "cluster-01",
				ReplicationSet:            "rs-01",
				Username:                  "some-username",
				Password:                  "some-password",
				AWSAccessKey:              "my-aws-access-key",
				AWSSecretKey:              "my-aws-secret-key",
				RDSExporter:               true,
				QANMysqlPerfschema:        true,
				CustomLabels:              map[string]string{},
				SkipConnectionCheck:       true,
				TLS:                       false,
				TLSSkipVerify:             false,
				DisableQueryExamples:      false,
				TablestatsGroupTableLimit: 2000,
				DisableBasicMetrics:       true,
				DisableEnhancedMetrics:    true,
			},
			Context: pmmapitests.Context,
		}
		addRDSOK, err := client.Default.RDS.AddRDS(params)
		require.NoError(t, err)
		require.NotNil(t, addRDSOK.Payload)

		body := addRDSOK.Payload
		assert.True(t, body.RDSExporter.BasicMetricsDisabled)
		assert.True(t, body.RDSExporter.EnhancedMetricsDisabled)

		pmmapitests.RemoveAgents(t, body.MysqldExporter.AgentID)
		pmmapitests.RemoveAgents(t, body.QANMysqlPerfschema.AgentID)
		pmmapitests.RemoveServices(t, body.Mysql.ServiceID)

		_, err = inventoryClient.Default.Agents.GetAgent(&agents.GetAgentParams{
			Body: agents.GetAgentBody{
				AgentID: body.RDSExporter.AgentID,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, fmt.Sprintf(`Agent with ID "%s" not found.`, body.RDSExporter.AgentID))

		_, err = inventoryClient.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body: nodes.GetNodeBody{
				NodeID: body.Mysql.NodeID,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, fmt.Sprintf(`Node with ID "%s" not found.`, body.Mysql.NodeID))
	})

	t.Run("AddRDSPostgres", func(t *testing.T) {
		params := &rds.AddRDSParams{
			Body: rds.AddRDSBody{
				Region:                    "region",
				Az:                        "az",
				InstanceID:                "d752f1a9-31c9-4b8c-bb2d-d26bc000009",
				NodeModel:                 "some-model",
				Address:                   "some.example.rds",
				Port:                      5432,
				Engine:                    pointer.ToString("DISCOVER_RDS_POSTGRESQL"),
				NodeName:                  "some-node-name-000009",
				ServiceName:               "test-add-rds-service000009",
				Environment:               "some-env",
				Cluster:                   "cluster-01",
				ReplicationSet:            "rs-01",
				Username:                  "some-username",
				Password:                  "some-password",
				AWSAccessKey:              "my-aws-access-key",
				AWSSecretKey:              "my-aws-secret-key",
				RDSExporter:               true,
				CustomLabels:              map[string]string{},
				SkipConnectionCheck:       true,
				TLS:                       false,
				TLSSkipVerify:             false,
				TablestatsGroupTableLimit: 2000,
				DisableBasicMetrics:       true,
				DisableEnhancedMetrics:    true,
				QANPostgresqlPgstatements: true,
			},
			Context: pmmapitests.Context,
		}
		addRDSOK, err := client.Default.RDS.AddRDS(params)
		require.NoError(t, err)
		require.NotNil(t, addRDSOK.Payload)

		body := addRDSOK.Payload
		assert.True(t, body.RDSExporter.BasicMetricsDisabled)
		assert.True(t, body.RDSExporter.EnhancedMetricsDisabled)

		pmmapitests.RemoveAgents(t, body.PostgresqlExporter.AgentID)
		pmmapitests.RemoveAgents(t, body.QANPostgresqlPgstatements.AgentID)
		pmmapitests.RemoveServices(t, body.Postgresql.ServiceID)

		_, err = inventoryClient.Default.Agents.GetAgent(&agents.GetAgentParams{
			Body: agents.GetAgentBody{
				AgentID: body.RDSExporter.AgentID,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, fmt.Sprintf(`Agent with ID "%s" not found.`, body.RDSExporter.AgentID))

		_, err = inventoryClient.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body: nodes.GetNodeBody{
				NodeID: body.Postgresql.NodeID,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, fmt.Sprintf(`Node with ID "%s" not found.`, body.Postgresql.NodeID))
	})
}
