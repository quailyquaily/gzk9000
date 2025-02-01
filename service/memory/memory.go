package memory

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
) *service {
	return &service{
		cfg:    cfg,
		aiInst: aiInst,
		qd:     qd,
	}
}

type (
	Config struct {
		CollectionFact string `json:"collection_fact"`
	}
	service struct {
		cfg Config

		aiInst *ai.Instant
		qd     *qdrant.QdrantClient
	}
)

func (s *service) CreateFact(ctx context.Context, ft *core.Fact) error {
	vec, err := s.aiInst.CreateEmbeddingAzureOpenAI(ctx, []string{ft.Content})
	if err != nil {
		return err
	}

	params := qdrant.UpsertPointsParams{}
	params.CollectionName = s.cfg.CollectionFact
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

func (s *service) FindSimilarFactsByAgentID(ctx context.Context, vec []float32, agentID uint64, topK uint64) ([]any, error) {
	result, err := s.qd.SearchPointsWithFilter(ctx, qdrant.SearchPointsParams{
		CollectionName: s.cfg.CollectionFact,
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

	// // get post details
	// pts, err := s.facts.GetPostsByIDs(ctx, topSimilarPostIDs)
	// if err != nil {
	// 	return nil, err
	// }

	// similarPostItems := make([]any, 0)
	// for _, sim := range result {
	// 	for _, item := range pts {
	// 		if item.PublishedAt == nil {
	// 			continue
	// 		}
	// 		if item.ID == uint64(sim.ID) {
	// 			similarPostItems = append(similarPostItems, core.SimilarPostItem{
	// 				PostID:        item.ID,
	// 				Slug:          item.Slug,
	// 				Title:         item.Title,
	// 				Summary:       item.Summary,
	// 				CoverImageURL: item.CoverImageURL,
	// 				Similarity:    float64(sim.Score),
	// 			})
	// 			break
	// 		}
	// 	}
	// }
	return nil, nil
}
