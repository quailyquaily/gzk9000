package fact

import (
	"context"
	"sort"

	"github.com/lyricat/goutils/ai"
	"github.com/lyricat/goutils/qdrant"
	"github.com/quailyquaily/gzk9000/core"
)

func New(
	cfg Config,
	aiInst *ai.Instant, qd *qdrant.QdrantClient,
	facts core.FactStore,
) *service {
	return &service{
		cfg:    cfg,
		aiInst: aiInst,
		qd:     qd,
		facts:  facts,
	}
}

type (
	Config struct {
		CollectionName string
	}
	service struct {
		cfg Config

		aiInst *ai.Instant
		qd     *qdrant.QdrantClient

		facts core.FactStore
	}
)

func (s *service) CreateFact(ctx context.Context, ft *core.Fact) error {
	vec, err := s.aiInst.CreateEmbeddingAzureOpenAI(ctx, []string{ft.Content})
	if err != nil {
		return err
	}

	params := qdrant.UpsertPointsParams{}
	params.CollectionName = s.cfg.CollectionName
	params.PointID = ft.ID
	params.Vector = vec
	params.Payload = make(map[string]qdrant.UpsertPointPayloadItem)
	params.Payload["fact_id"] = qdrant.UpsertPointPayloadItem{Type: "int", Value: int64(ft.ID)}
	params.Payload["agent_id"] = qdrant.UpsertPointPayloadItem{Type: "int", Value: int64(ft.AgentID)}
	if err := s.qd.UpsertPoints(ctx, params); err != nil {
		return err
	}

	return nil
}

func (s *service) FindSimilarFactsByAgentID(ctx context.Context, vec []float32, agentID uint64, topK uint64) ([]*core.Fact, error) {
	result, err := s.qd.SearchPointsWithFilter(ctx, qdrant.SearchPointsParams{
		CollectionName: s.cfg.CollectionName,
		Vector:         vec,
		TopK:           topK,
		Key:            "agent_id",
		Value:          int64(agentID),
		Offset:         0,
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})

	topFactIDs := make([]uint64, 0)
	for _, item := range result {
		topFactIDs = append(topFactIDs, uint64(item.Payload["fact_id"].GetIntegerValue()))
	}

	// get post details
	fts, err := s.facts.GetFactsByIDs(ctx, topFactIDs)
	if err != nil {
		return nil, err
	}

	return fts, nil
}
