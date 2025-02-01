package memslice

import (
	"github.com/quailyquaily/gzk9000/core"
	"github.com/quailyquaily/gzk9000/store"
	"github.com/quailyquaily/gzk9000/store/memslice/dao"
	"gorm.io/gen"
)

func init() {
	store.RegistGenerate(
		gen.Config{
			OutPath: "store/memslice/dao",
		},
		func(g *gen.Generator) {
			g.ApplyInterface(func(core.MemsliceStore) {}, core.Memslice{})
		},
	)
}

func New(h *store.Handler) core.MemsliceStore {
	var q *dao.Query
	if !dao.Q.Available() {
		dao.SetDefault(h.DB)
		q = dao.Q
	} else {
		q = dao.Use(h.DB)
	}

	v, ok := interface{}(q.Memslice).(core.MemsliceStore)
	if !ok {
		panic("dao.Memslice is not core.MemsliceStore")
	}

	return &storeImpl{
		MemsliceStore: v,
	}
}

type storeImpl struct {
	core.MemsliceStore
}
