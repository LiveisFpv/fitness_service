package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
)

func (r *Queries) GetTrainingById(ctx context.Context, training_id int) (*models.Training, error) {
	sqlStatement := `SELECT * FROM training WHERE training_id=$1`
	training := &models.Training{}
	err := r.pool.QueryRow(ctx, sqlStatement, training_id).Scan(
		&training.Training_id,
		&training.Training_title,
		&training.Training_desc,
		&training.Training_user_level,
	)
	if err != nil {
		return nil, fmt.Errorf("couldn`t find training: %w", err)
	}
	return training, nil
}

func (r *Queries) AddTraining(ctx context.Context, training *models.Training) (*models.Training, error) {
	sqlStatement := `INSERT INTO training (training_title, training_desc, training_user_level) VALUES ($1, $2, $3) RETURNING training_id`

	training_id := 0
	err := r.pool.QueryRow(ctx, sqlStatement, training.Training_title, training.Training_desc, training.Training_user_level).Scan(&training_id)
	if err != nil {
		return nil, fmt.Errorf("can`t create training: %w", err)
	}

	training, err = r.GetTrainingById(ctx, training_id)
	if err != nil {
		return nil, err
	}

	return training, nil
}

func (r *Queries) UpdateTraining(ctx context.Context, training *models.Training) (*models.Training, error) {
	sqlStatement := `UPDATE training SET training_title=$2, training_desc=$3, training_user_level=$4 WHERE training_id=$1`
	_, err := r.pool.Exec(
		ctx, sqlStatement,
		training.Training_id,
		training.Training_title,
		training.Training_desc,
		training.Training_user_level,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t update training: %w", err)
	}

	new_training, err := r.GetTrainingById(ctx, training.Training_id)
	if err != nil {
		return nil, err
	}

	return new_training, nil
}
func (r *Queries) DeleteTraining(ctx context.Context, training_id int) (*models.Training, error) {
	sqlStatement := `DELETE FROM training WHERE training_id=$1 RETURNING training_id, training_title, training_desc, training_user_level`
	training := &models.Training{}
	err := r.pool.QueryRow(ctx, sqlStatement, training_id).Scan(
		&training.Training_id,
		&training.Training_title,
		&training.Training_desc,
		&training.Training_user_level,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t delete training: %w", err)
	}
	return training, nil
}
func (r *Queries) AddTrainingPlan(ctx context.Context, training_plan *models.TrainingPlan) (*models.TrainingPlan, error) {
	sqlStatement := `INSERT INTO train_plan (training_id, user_id, date) VALUES ($1, $2, $3) RETURNING training_id, user_id, date`
	new_training_plan := &models.TrainingPlan{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		training_plan.Training_id,
		training_plan.User_id,
		training_plan.Date,
	).Scan(
		&new_training_plan.Training_id,
		&new_training_plan.User_id,
		&new_training_plan.Date,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t add new training plan: %w", err)
	}
	return new_training_plan, nil
}
func (r *Queries) UpdateTrainingPlan(ctx context.Context, training_plan *models.TrainingPlan) (*models.TrainingPlan, error) {
	sqlStatement := `UPDATE train_plan SET date=$3 WHERE training_id=$1, user_id=$2 RETURNING training_id, user_id, date`
	new_training_plan := &models.TrainingPlan{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		training_plan.Training_id,
		training_plan.User_id,
		training_plan.Date,
	).Scan(
		&new_training_plan.Training_id,
		&new_training_plan.User_id,
		&new_training_plan.Date,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t update training plan: %w", err)
	}
	return new_training_plan, nil
}
func (r *Queries) DeleteTrainingPlan(ctx context.Context, training_id, user_id int) (*models.TrainingPlan, error) {
	sqlStatement := `DELETE FROM train_plan WHERE training_id=$1, user_id=$2 RETURNING training_id, user_id, date`
	old_training_plan := &models.TrainingPlan{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		training_id,
		user_id,
	).Scan(
		&old_training_plan.Training_id,
		&old_training_plan.User_id,
		&old_training_plan.Date,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t delete training plan: %w", err)
	}
	return old_training_plan, nil
}
func (r *Queries) GetTrainingInsrtuctionsList(ctx context.Context, training_id int) ([]*models.TrainingInstructions, error) {
	sqlStatement := `SELECT * from training_instr where training_id=$1`

	rows, err := r.pool.Query(ctx, sqlStatement, training_id)
	if err != nil {
		return nil, fmt.Errorf("can`t find instructions list: %w", err)
	}

	instructions := []*models.TrainingInstructions{}
	for rows.Next() {
		instruction := &models.TrainingInstructions{}
		err := rows.Scan(
			&instruction.Training_id,
			&instruction.Training_order,
			&instruction.Training_instr,
			&instruction.Training_img,
		)
		if err != nil {
			return nil, fmt.Errorf("can`t process query result: %w", err)
		}
		instructions = append(instructions, instruction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return instructions, nil
}
func (r *Queries) AddTrainingInstruction(ctx context.Context, instruction *models.TrainingInstructions) (*models.TrainingInstructions, error) {
	sqlStatement := `INSERT INTO training_instr (training_id, training_order, training_instr, training_img) VALUES ($1, $2, $3, $4) RETURNING training_id, training_order, training_instr, training_img`
	new_instructions := &models.TrainingInstructions{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		instruction.Training_id,
		instruction.Training_order,
		instruction.Training_instr,
		instruction.Training_img,
	).Scan(
		&new_instructions.Training_id,
		&new_instructions.Training_order,
		&new_instructions.Training_instr,
		&new_instructions.Training_img,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t add new training instruction: %w", err)
	}
	return new_instructions, nil
}
func (r *Queries) UpdateTrainingInstruction(ctx context.Context, instruction *models.TrainingInstructions) (*models.TrainingInstructions, error) {
	sqlStatement := `UPDATE training_instr SET training_instr=$3, training_img=$4 WHERE training_id=$1, training_order=$2 RETURNING training_id, training_order, training_instr, training_img`
	new_instructions := &models.TrainingInstructions{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		instruction.Training_id,
		instruction.Training_order,
		instruction.Training_instr,
		instruction.Training_img,
	).Scan(
		&new_instructions.Training_id,
		&new_instructions.Training_order,
		&new_instructions.Training_instr,
		&new_instructions.Training_img,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t update training instruction: %w", err)
	}
	return new_instructions, nil
}
func (r *Queries) DeleteTrainingInstruction(ctx context.Context, training_id, training_order int) (*models.TrainingInstructions, error) {
	sqlStatement := `DELETE FROM training_instr WHERE training_id=$1, training_order=$2 RETURNING training_id, training_order, training_instr, training_img`
	old_instructions := &models.TrainingInstructions{}
	err := r.pool.QueryRow(ctx, sqlStatement, training_id, training_order).Scan(
		&old_instructions.Training_id,
		&old_instructions.Training_order,
		&old_instructions.Training_instr,
		&old_instructions.Training_img,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t delete training instruction: %w", err)
	}
	return old_instructions, nil
}
