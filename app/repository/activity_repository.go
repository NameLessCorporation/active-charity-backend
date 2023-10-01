package repository

import (
	"context"
	"database/sql"

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
