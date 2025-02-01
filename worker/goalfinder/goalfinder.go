package goalfinder

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/lyricat/goutils/ai"
	"github.com/quailyquaily/gzk9000/core"
	"gorm.io/gorm"
)

const (
	personalityPrompt = `
	You act as a highschool girl, you're in age 17.
	Here is your personality passion: a strong sense of purpose and a drive to pursue activities that are meaningful and enjoyable.
	`
)

type Worker struct {
	cfg       Config
	aiInst    *ai.Instant
	memslicez core.MemsliceService

	studygoals core.StudygoalStore
	agents     core.AgentStore
}

type Config struct {
	Name    string
	AIDebug bool
}

func New(
	cfg Config,
	aiInst *ai.Instant,
	memslicez core.MemsliceService,

	studygoals core.StudygoalStore,
	agents core.AgentStore,
) *Worker {
	return &Worker{
		cfg:        cfg,
		aiInst:     aiInst,
		memslicez:  memslicez,
		studygoals: studygoals,
		agents:     agents,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	// slog.Info("[goalfinder] worker started")
	// dur := time.Millisecond
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		return ctx.Err()
	// 	case <-time.After(dur):
	// 		if err := w.run(ctx); err != nil {
	// 			if !errors.Is(err, gorm.ErrRecordNotFound) {
	// 				slog.Error("[goalfinder] failed to run", "error", err)
	// 				dur = time.Second * 180
	// 			} else {
	// 				dur = time.Second * 180
	// 			}
	// 		} else {
	// 			// wait for 60s
	// 			dur = time.Second * 60
	// 		}
	// 	}
	// }
	return nil
}

func (w *Worker) run(ctx context.Context) error {
	// get all agents
	// for each agent, get all memslicez in the last 24 hours
	//   for each memslice, find all related facts
	//   create a prompt with all facts, ask AI to find the topics today
	//   get the topics
	//   compare the topics with the goals, find the new goals
	//   create the new goals
	ags, err := w.agents.GetAllAgents(ctx)
	if err != nil {
		return err
	}
	for _, ag := range ags {
		now := time.Now()
		startTime := now.Add(-24 * time.Hour)

		mmslcs, err := w.memslicez.GetMemslicesByRange(ctx, ag.ID, &startTime, &now)
		if err != nil {
			return err
		}

		prompt := w.buildExtractGoalsPrompt(mmslcs)
		rp := &ai.SusanoParams{
			Format: "json",
			Conditions: ai.SusanoParamsConditions{
				PreferredModel: "o1-mini",
			},
		}
		ret, err := w.aiInst.OneTimeRequestWithParams(ctx, prompt, rp.ToMap())
		if err != nil {
			return err
		}

		if val, ok := ret.Json["goals"]; !ok || len(val.([]any)) == 0 {
			return nil
		}

		possibleGoalList := []string{}
		for _, v := range ret.Json["goals"].([]any) {
			possibleGoalList = append(possibleGoalList, v.(string))
		}
		fmt.Printf("possibleGoalList: %+v\n", possibleGoalList)

		existingGoals, err := w.studygoals.GetActiveStudygoalsByAgentID(ctx, ag.ID)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		filteredGoalList := []string{}
		if len(existingGoals) != 0 {
			slog.Info("[goalfinder] compare goals.")
			prompt = w.buildCompareGoalsPrompt(existingGoals, possibleGoalList)
			rp = &ai.SusanoParams{
				Format: "json",
				Conditions: ai.SusanoParamsConditions{
					PreferredModel: "o1-mini",
				},
			}
			ret, err = w.aiInst.OneTimeRequestWithParams(ctx, prompt, rp.ToMap())
			if err != nil {
				return err
			}
			if val, ok := ret.Json["goals"]; ok {
				for _, v := range val.([]any) {
					filteredGoalList = append(filteredGoalList, v.(string))
				}
			}
		} else {
			filteredGoalList = possibleGoalList
		}

		fmt.Printf("filteredGoalList: %+v\n", filteredGoalList)

		if err := w.insertNewGoals(ctx, ag.ID, filteredGoalList); err != nil {
			slog.Error("[goalfinder] failed to insert new goals", "error", err)
			return err
		}
	}

	return nil
}

func (w *Worker) insertNewGoals(ctx context.Context, agID uint64, goals []string) error {
	for _, goal := range goals {
		sg := &core.Studygoal{
			AgentID: agID,
			Content: goal,
			Status:  core.StudygoalStatusActive,
		}
		if err := w.studygoals.CreateStudygoal(ctx, sg); err != nil {
			return err
		}
	}
	return nil
}

func (w *Worker) buildCompareGoalsPrompt(existingGoals []*core.Studygoal, possibleGoalList []string) string {
	promptTpl := personalityPrompt + `
	Compare the goals you have with the possible goals you may have.
	You need to remove the duplicated goals, similar goals and irrelevant goals from the possible goals.
	You output the filtered possible goals as JSON.
	Here is the example of the output:
	{
		"goals": [ "goal1", "goal2", "goal3" ...]
	}

	Here is the existing goals you have:
	%s

	And here are the possible goals you may have:
	%s
	`
	egRows := []string{}
	for i, eg := range existingGoals {
		egRows = append(egRows, fmt.Sprintf("existing goal %d: %s", i, eg.Content))
	}

	pgRows := []string{}
	for i, pg := range possibleGoalList {
		pgRows = append(pgRows, fmt.Sprintf("possible goal %d: %s", i, pg))
	}

	return fmt.Sprintf(promptTpl, strings.Join(egRows, "\n"), strings.Join(pgRows, "\n"))
}

func (w *Worker) buildExtractGoalsPrompt(mmslcs []*core.Memslice) string {
	promptTpl := `
	You act as a highschool girl, you're in age 17.
	Here is your personality passion: a strong sense of purpose and a drive to pursue activities that are meaningful and enjoyable.

	Here are memory pieces recalled from your memory in past 24 hours.
	Your memory may be varied.

	You try to find some goals worth to further research later from the memory pieces.
	The goals are something you want to achieve or something you want to know in the future, not now.
	The goals need to worth to continue to research or overthink.
	The goals shouldn't be too easy to achieve.
	The goals should decribed in your way and follow your personality.
	No duplicates, similar and irrelevant goals.
	In most cases, each memory piece may only have zero goal or one goal.
	You must try to simplify the goal expression.
	You must evaluate the goal based on your personality and value.

	Output the goals as JSON.
	If there are not goal worth to mention, you simply output a empty list.

	Here are the examples:

	some goals are found:
	{
		"goals":	[ "goal1", "goal2", "goal3" ...]
  }

	only one goal is found:
	{
		"goals": [ "goal" ]
	}

	no goal are found:
	{
		"goals": [ ]
  }

	memorys pieces:
	%s
	`
	rows := []string{}
	for i, mmslc := range mmslcs {
		rows = append(rows, fmt.Sprintf("memory piece %d: %s", i, mmslc.Content))
		for j, fact := range mmslc.IncludedFacts {
			// build the prompt
			rows = append(rows, fmt.Sprintf("- fact %d: %s", j, fact.Content))
		}
	}
	rowStr := strings.Join(rows, "\n")
	return fmt.Sprintf(promptTpl, rowStr)
}
