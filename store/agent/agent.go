package agent

import (
	"github.com/quailyquaily/gzk9000/core"
	"github.com/quailyquaily/gzk9000/store"
	"github.com/quailyquaily/gzk9000/store/agent/dao"
	"gorm.io/gen"
)

func init() {
	store.RegistGenerate(
		gen.Config{
			OutPath: "store/agent/dao",
		},
		func(g *gen.Generator) {
			g.ApplyInterface(func(core.AgentStore) {}, core.Agent{})
		},
	)
}

func New(h *store.Handler) core.AgentStore {
	var q *dao.Query
	if !dao.Q.Available() {
		dao.SetDefault(h.DB)
		q = dao.Q
	} else {
		q = dao.Use(h.DB)
	}

	v, ok := interface{}(q.Agent).(core.AgentStore)
	if !ok {
		panic("dao.Agent is not core.AgentStore")
	}

	return &storeImpl{
		AgentStore: v,
	}
}

type storeImpl struct {
	core.AgentStore
}
