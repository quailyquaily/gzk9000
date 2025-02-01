package studygoal

import (
	"github.com/quailyquaily/gzk9000/core"
	"github.com/quailyquaily/gzk9000/store"
	"github.com/quailyquaily/gzk9000/store/studygoal/dao"
	"gorm.io/gen"
)

func init() {
	store.RegistGenerate(
		gen.Config{
			OutPath: "store/studygoal/dao",
		},
		func(g *gen.Generator) {
			g.ApplyInterface(func(core.StudygoalStore) {}, core.Studygoal{})
		},
	)
}

func New(h *store.Handler) core.StudygoalStore {
	var q *dao.Query
	if !dao.Q.Available() {
		dao.SetDefault(h.DB)
		q = dao.Q
	} else {
		q = dao.Use(h.DB)
	}

	v, ok := interface{}(q.Studygoal).(core.StudygoalStore)
	if !ok {
		panic("dao.Studygoal is not core.StudygoalStore")
	}

	return &storeImpl{
		StudygoalStore: v,
	}
}

type storeImpl struct {
	core.StudygoalStore
}
