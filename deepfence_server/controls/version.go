package controls

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/deepfence/ThreatMapper/deepfence_utils/controls"
	"github.com/deepfence/ThreatMapper/deepfence_utils/directory"
	"github.com/deepfence/ThreatMapper/deepfence_utils/utils"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const (
	DEFAULT_AGENT_IMAGE_NAME = "deepfence.io"
	DEFAULT_AGENT_IMAGE_TAG  = "thomas"
	DEFAULT_AGENT_VERSION    = "0.0.1"
)

func ScheduleAgentUpgrade(ctx context.Context, version string, nodeIds []string, action controls.Action) error {

	client, err := directory.Neo4jClient(ctx)
	if err != nil {
		return err
	}

	session, err := client.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return err
	}
	defer tx.Close()

	action_str, err := json.Marshal(action)
	if err != nil {
		return err
	}

	_, err = tx.Run(`
		MATCH (v:AgentVersion{node_id: $version})
		MATCH (n:Node)
		WHERE n.node_id IN $node_ids
		MERGE (v) -[:SCHEDULED{status: $status, retries: 0, trigger_action: $action, updated_at: TIMESTAMP()}]-> (n)`,
		map[string]interface{}{
			"version":  version,
			"node_ids": nodeIds,
			"status":   utils.SCAN_STATUS_STARTING,
			"action":   string(action_str),
		})

	if err != nil {
		return err
	}

	return tx.Commit()

}

func GetAgentVersionTarball(ctx context.Context, version string) (string, error) {

	client, err := directory.Neo4jClient(ctx)
	if err != nil {
		return "", err
	}

	session, err := client.Session(neo4j.AccessModeRead)
	if err != nil {
		return "", err
	}
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return "", err
	}
	defer tx.Close()

	res, err := tx.Run(`
		MATCH (v:AgentVersion{node_id: $version})
		return v.url`,
		map[string]interface{}{
			"version": version,
		})

	if err != nil {
		return "", err
	}

	r, err := res.Single()

	if err != nil {
		return "", err
	}

	return r.Values[0].(string), nil
}

func GetAgentPluginVersionTarball(ctx context.Context, version, plugin_name string) (string, error) {

	client, err := directory.Neo4jClient(ctx)
	if err != nil {
		return "", err
	}

	session, err := client.Session(neo4j.AccessModeRead)
	if err != nil {
		return "", err
	}
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return "", err
	}
	defer tx.Close()

	query := fmt.Sprintf(`
		MATCH (v:AgentVersion{node_id: $version})
		return v.url_%s`, plugin_name)
	res, err := tx.Run(query,
		map[string]interface{}{
			"version": version,
		})

	if err != nil {
		return "", err
	}

	r, err := res.Single()

	if err != nil {
		return "", err
	}

	return r.Values[0].(string), nil
}

func hasPendingUpgradeOrNew(ctx context.Context, version string, nodeId string) (bool, error) {

	client, err := directory.Neo4jClient(ctx)
	if err != nil {
		return false, err
	}

	session, err := client.Session(neo4j.AccessModeRead)
	if err != nil {
		return false, err
	}
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return false, err
	}
	defer tx.Close()

	res, err := tx.Run(`
		MATCH (n:Node{node_id:$node_id})
		MATCH (v:AgentVersion{node_id:$version})
		OPTIONAL MATCH (v) -[rs:SCHEDULED]-> (n)
		OPTIONAL MATCH (n) -[rv:VERSIONED]-> (v)
		RETURN rs IS NOT NULL OR rv IS NULL`,
		map[string]interface{}{
			"node_id": nodeId,
			"version": version,
		})
	if err != nil {
		return false, err
	}

	r, err := res.Single()
	if err != nil {
		// No results means new
		return true, nil
	}
	return r.Values[0].(bool), nil
}

func CompleteAgentUpgrade(ctx context.Context, version string, nodeId string) error {

	has, err := hasPendingUpgradeOrNew(ctx, version, nodeId)

	if err != nil {
		return err
	}

	if !has {
		return nil
	}

	client, err := directory.Neo4jClient(ctx)
	if err != nil {
		return err
	}

	session, err := client.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return err
	}
	defer tx.Close()

	_, err = tx.Run(`
		OPTIONAL MATCH (n:Node{node_id:$node_id}) -[old:VERSIONED]-> (v)
		DELETE old`,
		map[string]interface{}{
			"node_id": nodeId,
		})
	if err != nil {
		return err
	}

	_, err = tx.Run(`
		MERGE (n:Node{node_id:$node_id})
		MERGE (v:AgentVersion{node_id:$version})
		MERGE (n) -[r:VERSIONED]-> (v)
		WITH n, v
		OPTIONAL MATCH (v) -[r:SCHEDULED]-> (n)
		DELETE r`,
		map[string]interface{}{
			"version": version,
			"node_id": nodeId,
		})

	if err != nil {
		return err
	}

	return tx.Commit()

}

func ScheduleAgentPluginEnable(ctx context.Context, version, plugin_name string, nodeIds []string, action controls.Action) error {

	client, err := directory.Neo4jClient(ctx)
	if err != nil {
		return err
	}

	session, err := client.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return err
	}
	defer tx.Close()

	action_str, err := json.Marshal(action)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`
		MATCH (v:%sVersion{node_id: $version})
		MATCH (n:Node)
		WHERE n.node_id IN $nonternal_req := ctl.EnableAgentPluginRequest{
		BinUrl:     url,
		Version:    agentUp.Version,
		PluginName: agentUp.PluginName,
		}
		de_ids
		MERGE (v) -[:SCHEDULED{status: $status, retries: 0, trigger_action: $action, updated_at: TIMESTAMP()}]-> (n)`, plugin_name)

	_, err = tx.Run(query,
		map[string]interface{}{
			"version":  version,
			"node_ids": nodeIds,
			"status":   utils.SCAN_STATUS_STARTING,
			"action":   string(action_str),
		})

	if err != nil {
		return err
	}

	return tx.Commit()

}

func ScheduleAgentPluginDisable(ctx context.Context, plugin_name string, nodeIds []string, action controls.Action) error {

	client, err := directory.Neo4jClient(ctx)
	if err != nil {
		return err
	}

	session, err := client.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return err
	}
	defer tx.Close()

	action_str, err := json.Marshal(action)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`
		MATCH (n:Node) -[:USES]-> (v:%sVersion)
		WHERE n.node_id IN $node_ids
		MERGE (v) -[:SCHEDULED{status: $status, retries: 0, trigger_action: $action, updated_at: TIMESTAMP()}]-> (n)
		SET n.status_%s = 'disabling'`, plugin_name, plugin_name)

	_, err = tx.Run(query,
		map[string]interface{}{
			"node_ids": nodeIds,
			"status":   utils.SCAN_STATUS_STARTING,
			"action":   string(action_str),
		})

	if err != nil {
		return err
	}

	return tx.Commit()

}
