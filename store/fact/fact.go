package fact

import (
	"github.com/quailyquaily/gzk9000/core"
	"github.com/quailyquaily/gzk9000/store"
	"github.com/quailyquaily/gzk9000/store/fact/dao"
	"gorm.io/gen"
)

func init() {
	store.RegistGenerate(
		gen.Config{
			OutPath: "store/fact/dao",
		},
		func(g *gen.Generator) {
			g.ApplyInterface(func(core.FactStore) {}, core.Fact{})
		},
	)
}

func New(h *store.Handler) core.FactStore {
	var q *dao.Query
	if !dao.Q.Available() {
		dao.SetDefault(h.DB)
		q = dao.Q
	} else {
		q = dao.Use(h.DB)
	}

	v, ok := interface{}(q.Fact).(core.FactStore)
	if !ok {
		panic("dao.Fact is not core.FactStore")
	}

	return &storeImpl{
		FactStore: v,
	}
}

type storeImpl struct {
	core.FactStore
}
