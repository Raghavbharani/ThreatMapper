package ingesters

import (
	"context"
	"encoding/json"

	"github.com/deepfence/ThreatMapper/deepfence_utils/directory"
	"github.com/deepfence/ThreatMapper/deepfence_utils/log"
	"github.com/deepfence/ThreatMapper/deepfence_utils/utils"
	"github.com/twmb/franz-go/pkg/kgo"
)

type SecretIngester struct{}

func NewSecretIngester() KafkaIngester[[]map[string]interface{}] {
	return &SecretIngester{}
}

func (tc *SecretIngester) Ingest(
	ctx context.Context,
	cs []map[string]interface{},
	ingestC chan *kgo.Record,
) error {
	tenantID, err := directory.ExtractNamespace(ctx)
	if err != nil {
		return err
	}

	rh := []kgo.RecordHeader{
		{Key: "tenant_id", Value: []byte(tenantID)},
	}

	for _, c := range cs {
		cb, err := json.Marshal(c)
		if err != nil {
			log.Error().Msg(err.Error())
		} else {
			ingestC <- &kgo.Record{
				Topic:   utils.SECRET_SCAN,
				Value:   cb,
				Headers: rh,
			}
		}
	}

	return nil
}
