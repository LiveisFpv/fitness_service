package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
	"time"
)

// TODO
func (r *Queries) GetWeightHistoryList(ctx context.Context, user_id int, date time.Time) ([]*models.WeightHistory, error) {
	sqlStatement := `SELECT weight, date from weight_hist where user_id=$1 and (TO_DATE(date, 'YYYY-MM-DD) between $2::date and ($2::date + INTERVAL '1 week')) ORDER BY date ASC`
	rows, err := r.pool.Query(ctx, sqlStatement, user_id, date)
	if err != nil {
		return nil, fmt.Errorf("can`t consturct weight history: %w", err)
	}
	response := []*models.WeightHistory{}
	for rows.Next() {
		resp := &models.WeightHistory{}
		var dateString string
		err := rows.Scan(
			&resp.Weight,
			&dateString,
		)
		if err != nil {
			return nil, fmt.Errorf("can`t process query result: %w", err)
		}
		parsedDate, err := time.Parse("2006-01-02T15:04:05", dateString+"T00:00:00")
		if err != nil {
			return nil, fmt.Errorf("can't parse date %s: %w", dateString, err)
		}
		resp.Date = parsedDate

		response = append(response, resp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return response, nil
}
func (r *Queries) AddWeightHistory(ctx context.Context, weight *models.WeightHistory) (*models.WeightHistory, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) UpdateWeightHistory(ctx context.Context, weight *models.WeightHistory) (*models.WeightHistory, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) DeleteWightHistory(ctx context.Context, user_id int, date time.Time) (*models.WeightHistory, error) {
	return nil, fmt.Errorf("not implemented")
}
