package repository

import (
	"context"

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

func (a *Activity) GetCurrentPeriodId(ctx context.Context, userId uint64) (uint64, error) {
	var periodId uint64

	err := a.db.GetContext(
		ctx,
		&periodId,
		`select id from steps_history 
                where created_at > CURRENT_DATE and 
                      created_at < CURRENT_DATE + interval '1 day' and
                      user_id = $1`,
		userId)
	if err != nil {
		return 0, err
	}

	return periodId, nil
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
		"update steps_history set steps = $1 and updated_at = now() where id = $2",
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

	err := a.db.SelectContext(ctx, &activities, "select id, name from activities")
	if err != nil {
		return nil, err
	}

	return activities, nil
}
