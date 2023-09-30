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

func (a *Activity) TrackSteps(ctx context.Context, steps uint32, activityId uint64, userId uint64) error {
	_, err := a.db.ExecContext(
		ctx,
		"insert into steps_history(user_id, steps, activity_id) values($1, $2, $3)",
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
