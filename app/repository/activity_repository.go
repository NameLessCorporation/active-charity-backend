package repository

import (
	"context"
	"database/sql"
	"sort"

	"github.com/jmoiron/sqlx"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

type Activity struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) *Activity {
	return &Activity{
		db: db,
	}
}

func (a *Activity) GetStepsGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	var graph []*models.Graph

	err := a.db.SelectContext(ctx, &graph, "select created_at, steps as value from steps_history where user_id = $1 order by id", userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Activity) GetPushUpGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	var graph []*models.Graph

	err := a.db.SelectContext(ctx, &graph, "select created_at, repeats as value from pull_ups_history where user_id = $1 order by id", userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Activity) GetCrunchesGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	var graph []*models.Graph

	err := a.db.SelectContext(ctx, &graph, "select created_at, repeats as value from crunches_history where user_id = $1 order by id", userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Activity) GetCycleGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	var graph []*models.Graph

	err := a.db.SelectContext(ctx, &graph, "select created_at, metres as value from cycling_history where user_id = $1 order by id", userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Activity) GetPullUpGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	var graph []*models.Graph

	err := a.db.SelectContext(ctx, &graph, "select created_at, repeats as value from pull_ups_history where user_id = $1 order by id", userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Activity) GetBenchPressGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	var graph []*models.Graph

	err := a.db.SelectContext(ctx, &graph, "select created_at, repeats as value from bench_press_history where user_id = $1 order by id", userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Activity) GetUserActivityAnalytics(ctx context.Context, userID uint64) ([]models.ActivityAnalytics, error) {
	activities := make([]models.ActivityAnalytics, 6)

	err := a.db.QueryRowxContext(ctx, `
		select a.id, max(sh.steps), min(sh.steps), avg(sh.steps) from activities as a
		inner join steps_history sh on a.id = sh.activity_id
		where user_id = $1
		group by a.id`,
		userID,
	).Scan(&activities[0].Id, &activities[0].Max, &activities[0].Min, &activities[0].Avg)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		activities[0] = models.ActivityAnalytics{}
	}

	err = a.db.QueryRowxContext(ctx, `
		select a.id, max(sh.repeats), min(sh.repeats), avg(sh.repeats) from activities as a
		inner join push_ups_history sh on a.id = sh.activity_id
		where user_id = $1
		group by a.id`,
		userID,
	).Scan(&activities[1].Id, &activities[1].Max, &activities[1].Min, &activities[1].Avg)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		activities[1] = models.ActivityAnalytics{}
	}

	err = a.db.QueryRowxContext(ctx, `
		select a.id, max(sh.repeats), min(sh.repeats), avg(sh.repeats) from activities as a
		inner join pull_ups_history sh on a.id = sh.activity_id
		where user_id = $1
		group by a.id`,
		userID,
	).Scan(&activities[2].Id, &activities[2].Max, &activities[2].Min, &activities[2].Avg)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		activities[2] = models.ActivityAnalytics{}
	}

	err = a.db.QueryRowxContext(ctx, `
		select a.id, max(sh.metres), min(sh.metres), avg(sh.metres) from activities as a
		inner join cycling_history sh on a.id = sh.activity_id
		where user_id = $1
		group by a.id`,
		userID,
	).Scan(&activities[3].Id, &activities[3].Max, &activities[3].Min, &activities[3].Avg)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		activities[3] = models.ActivityAnalytics{}
	}

	err = a.db.QueryRowxContext(ctx, `
		select a.id, max(sh.repeats), min(sh.repeats), avg(sh.repeats) from activities as a
		inner join crunches_history sh on a.id = sh.activity_id
		where user_id = $1
		group by a.id`,
		userID,
	).Scan(&activities[4].Id, &activities[4].Max, &activities[4].Min, &activities[4].Avg)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		activities[4] = models.ActivityAnalytics{}
	}

	err = a.db.QueryRowxContext(ctx, `
		select a.id, max(sh.repeats), min(sh.repeats), avg(sh.repeats) from activities as a
		inner join bench_press_history sh on a.id = sh.activity_id
		where user_id = $1
		group by a.id`,
		userID,
	).Scan(&activities[5].Id, &activities[5].Max, &activities[5].Min, &activities[5].Avg)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		activities[5] = models.ActivityAnalytics{}
	}
	return activities, nil
}

func (a *Activity) GetUserFavouriteActivity(ctx context.Context, userID uint64) (string, error) {
	var (
		pushUps  models.FavouriteActivity
		pullUps  models.FavouriteActivity
		cycling  models.FavouriteActivity
		crunches models.FavouriteActivity
		bench    models.FavouriteActivity
	)

	err := a.db.GetContext(ctx, &pullUps,
		`select count(*), a.name from pull_ups_history
				inner join activities a on a.id = pull_ups_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		pullUps = models.FavouriteActivity{
			Count: 0,
			Name:  "",
		}
	}

	err = a.db.GetContext(ctx, &pushUps,
		`select count(*), a.name from push_ups_history
				inner join activities a on a.id = push_ups_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		pushUps = models.FavouriteActivity{
			Count: 0,
			Name:  "",
		}
	}

	err = a.db.GetContext(ctx, &cycling,
		`select count(*), a.name from cycling_history
				inner join activities a on a.id = cycling_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		cycling = models.FavouriteActivity{
			Count: 0,
			Name:  "",
		}
	}

	err = a.db.GetContext(ctx, &crunches,
		`select count(*), a.name from crunches_history
				inner join activities a on a.id = crunches_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		crunches = models.FavouriteActivity{
			Count: 0,
			Name:  "",
		}
	}

	err = a.db.GetContext(ctx, &bench,
		`select count(*), a.name from bench_press_history
				inner join activities a on a.id = bench_press_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		bench = models.FavouriteActivity{
			Count: 0,
			Name:  "",
		}
	}

	arr := []models.FavouriteActivity{pushUps, pullUps, cycling, crunches, bench}
	sort.Slice(arr, func(i, j int) bool {
		return arr[j].Count < arr[i].Count
	})

	return arr[0].Name, nil
}

func (a *Activity) GetUserMostEarnedActivity(ctx context.Context, userID uint64) (string, error) {
	var (
		pushUps  models.MostEarned
		pullUps  models.MostEarned
		cycling  models.MostEarned
		crunches models.MostEarned
		bench    models.MostEarned
	)

	err := a.db.GetContext(ctx, &pullUps,
		`select sum(repeats), a.name from pull_ups_history
				inner join activities a on a.id = pull_ups_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		pullUps = models.MostEarned{
			Sum:  0,
			Name: "",
		}
	}

	err = a.db.GetContext(ctx, &pushUps,
		`select sum(repeats), a.name from push_ups_history
				inner join activities a on a.id = push_ups_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		pushUps = models.MostEarned{
			Sum:  0,
			Name: "",
		}
	}

	err = a.db.GetContext(ctx, &cycling,
		`select sum(metres), a.name from cycling_history
				inner join activities a on a.id = cycling_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		cycling = models.MostEarned{
			Sum:  0,
			Name: "",
		}
	}

	err = a.db.GetContext(ctx, &crunches,
		`select sum(repeats), a.name from crunches_history
				inner join activities a on a.id = crunches_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		crunches = models.MostEarned{
			Sum:  0,
			Name: "",
		}
	}

	err = a.db.GetContext(ctx, &bench,
		`select sum(repeats), a.name from bench_press_history
				inner join activities a on a.id = bench_press_history.activity_id
				where user_id = $1
				group by a.name;`,
		userID,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}

		bench = models.MostEarned{
			Sum:  0,
			Name: "",
		}
	}

	arr := []models.MostEarned{pushUps, pullUps, cycling, crunches, bench}
	sort.Slice(arr, func(i, j int) bool {
		return arr[j].Sum < arr[i].Sum
	})

	return arr[0].Name, nil
}

func (a *Activity) TrackPushUps(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {
	_, err := a.db.ExecContext(
		ctx,
		"insert into push_ups_history(user_id, repeats, activity_id, created_at, updated_at) values($1, $2, $3, now(), now())",
		userId,
		repeats,
		activityId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *Activity) TrackBenchPress(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {
	_, err := a.db.ExecContext(
		ctx,
		"insert into bench_press_history(user_id, repeats, activity_id, created_at, updated_at) values($1, $2, $3, now(), now())",
		userId,
		repeats,
		activityId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *Activity) TrackCycling(ctx context.Context, metres uint32, activityId uint64, userId uint64) error {
	_, err := a.db.ExecContext(
		ctx,
		"insert into cycling_history(user_id, metres, activity_id, created_at, updated_at) values($1, $2, $3, now(), now())",
		userId,
		metres,
		activityId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *Activity) TrackCrunches(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {
	_, err := a.db.ExecContext(
		ctx,
		"insert into crunches_history(user_id, repeats, activity_id, created_at, updated_at) values($1, $2, $3, now(), now())",
		userId,
		repeats,
		activityId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *Activity) TrackPullUps(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {
	_, err := a.db.ExecContext(
		ctx,
		"insert into pull_ups_history(user_id, repeats, activity_id, created_at, updated_at) values($1, $2, $3, now(), now())",
		userId,
		repeats,
		activityId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *Activity) GetCurrentPeriodId(ctx context.Context, userId uint64) (uint64, uint32, error) {
	var (
		periodId uint64
		value    uint32
	)

	err := a.db.QueryRowxContext(
		ctx,
		`select id, steps from steps_history 
                where created_at > CURRENT_DATE and 
                      created_at < CURRENT_DATE + interval '1 day' and
                      user_id = $1`,
		userId).Scan(&periodId, &value)
	if err != nil {
		return 0, 0, err
	}

	return periodId, value, nil
}

func (a *Activity) IsActiveStepsPeriod(ctx context.Context, userId uint64) (bool, error) {
	var count int32

	err := a.db.GetContext(
		ctx,
		&count,
		`select count(*) from steps_history 
                where created_at > CURRENT_DATE and 
                      created_at < CURRENT_DATE + interval '1 day' and
                      user_id = $1`, userId)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (a *Activity) TrackCurrentPeriodSteps(ctx context.Context, steps uint32, periodId uint64) error {
	_, err := a.db.ExecContext(
		ctx,
		"update steps_history set steps = $1, updated_at = now() where id = $2",
		steps,
		periodId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *Activity) TrackNewPeriodSteps(ctx context.Context, steps uint32, activityId uint64, userId uint64) error {
	_, err := a.db.ExecContext(
		ctx,
		"insert into steps_history(user_id, steps, activity_id, created_at, updated_at) values($1, $2, $3, now(), now())",
		userId,
		steps,
		activityId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *Activity) GetActivityList(ctx context.Context) ([]*models.Activity, error) {
	var activities []*models.Activity

	err := a.db.SelectContext(ctx, &activities, "select id, name, unit from activities")
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (a *Activity) GetStepsValue(ctx context.Context, userID uint64) (sql.NullInt32, error) {
	var value sql.NullInt32

	err := a.db.GetContext(ctx, &value, "select sum(steps) as value from steps_history where user_id = $1", userID)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return value, nil
}

func (a *Activity) GetBenchPressValue(ctx context.Context, userID uint64) (sql.NullInt32, error) {
	var value sql.NullInt32

	err := a.db.GetContext(ctx, &value, "select sum(repeats) from bench_press_history where user_id = $1", userID)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return value, nil
}

func (a *Activity) GetCrunchesValue(ctx context.Context, userID uint64) (sql.NullInt32, error) {
	var value sql.NullInt32

	err := a.db.GetContext(ctx, &value, "select sum(repeats) from crunches_history where user_id = $1", userID)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return value, nil
}

func (a *Activity) GetCyclingValue(ctx context.Context, userID uint64) (sql.NullInt32, error) {
	var value sql.NullInt32

	err := a.db.GetContext(ctx, &value, "select sum(metres) as value from cycling_history where user_id = $1", userID)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return value, nil
}

func (a *Activity) GetPullUpValue(ctx context.Context, userID uint64) (sql.NullInt32, error) {
	var value sql.NullInt32

	err := a.db.GetContext(ctx, &value, "select sum(repeats) from pull_ups_history where user_id = $1", userID)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return value, nil
}

func (a *Activity) GetPushUpValue(ctx context.Context, userID uint64) (sql.NullInt32, error) {
	var value sql.NullInt32

	err := a.db.GetContext(ctx, &value, "select sum(repeats) from push_ups_history where user_id = $1", userID)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return value, nil
}
