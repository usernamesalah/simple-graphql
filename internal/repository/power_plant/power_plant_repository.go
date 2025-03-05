package powerPlantrepository

import (
	"context"
	"database/sql"
	"tensor-graphql/internal/model"
	repository "tensor-graphql/internal/repository/common"
	"tensor-graphql/pkg/derrors"
)

type (
	powerPlantRepository struct {
		repository.Repository
	}

	PowerPlantRepository interface {
		repository.Repository
		CreatePowerPlant(ctx context.Context, tx *sql.Tx, powerPlant *model.PowerPlant) (err error)
		GetPowerPlantByID(ctx context.Context, id string) (powerPlant *model.PowerPlant, err error)
		GetPowerPlants(ctx context.Context, page, limit int) (powerPlants []*model.PowerPlant, total int, err error)
		UpdatePowerPlant(ctx context.Context, tx *sql.Tx, powerPlant *model.PowerPlant) (err error)
		DeletePowerPlant(ctx context.Context, tx *sql.Tx, id string) (err error)
	}
)

func NewPowerPlantRepository(store repository.Repository) PowerPlantRepository {
	return &powerPlantRepository{
		Repository: store,
	}
}

func (r *powerPlantRepository) CreatePowerPlant(ctx context.Context, tx *sql.Tx, powerPlant *model.PowerPlant) (err error) {
	defer derrors.Wrap(&err, "CreatePowerPlant(%q)", powerPlant.ID)

	query := `INSERT INTO power_plant (name, latitude, longitude) VALUES (?, ?, ?)`
	args := []interface{}{
		powerPlant.Name,
		powerPlant.Latitude,
		powerPlant.Longitude,
	}

	_, err = r.Exec(ctx, tx, query, args)
	if err != nil {
		return derrors.WrapStack(err, derrors.Unknown, "r.Exec")
	}

	return
}

func (r *powerPlantRepository) GetPowerPlantByID(ctx context.Context, id string) (powerPlant *model.PowerPlant, err error) {
	defer derrors.Wrap(&err, "GetPowerPlantByID(%q)", id)

	query := `SELECT id, name, latitude, longitude FROM power_plant WHERE id = ?`
	powerPlant = &model.PowerPlant{}
	args := []any{
		id,
	}

	err = r.Query(ctx, query, r.getDest(powerPlant), args)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, derrors.HandleSQLError(err, "r.Query")
	}

	return powerPlant, nil
}

func (r *powerPlantRepository) UpdatePowerPlant(ctx context.Context, tx *sql.Tx, powerPlant *model.PowerPlant) (err error) {
	defer derrors.Wrap(&err, "UpdatePowerPlant(%q)", powerPlant.ID)

	query := `UPDATE power_plant SET name = ?, latitude = ?, longitude = ? WHERE id = ?`
	args := []interface{}{
		powerPlant.Name,
		powerPlant.Latitude,
		powerPlant.Longitude,
		powerPlant.ID,
	}

	_, err = r.Exec(ctx, tx, query, args)
	if err != nil {
		return derrors.WrapStack(err, derrors.Unknown, "r.Exec")
	}

	return nil
}

func (r *powerPlantRepository) DeletePowerPlant(ctx context.Context, tx *sql.Tx, id string) (err error) {
	defer derrors.Wrap(&err, "DeletePowerPlant(%q)", id)

	query := `DELETE FROM power_plant WHERE id = ?`
	args := []interface{}{
		id,
	}

	_, err = r.Exec(ctx, tx, query, args)
	if err != nil {
		return derrors.WrapStack(err, derrors.Unknown, "r.Exec")
	}

	return nil
}

func (r *powerPlantRepository) getDest(powerPlant *model.PowerPlant) []interface{} {
	return []interface{}{
		&powerPlant.ID,
		&powerPlant.Name,
		&powerPlant.Latitude,
		&powerPlant.Longitude,
	}
}

func (r *powerPlantRepository) GetPowerPlants(ctx context.Context, page, limit int) (powerPlants []*model.PowerPlant, total int, err error) {
	defer derrors.Wrap(&err, "GetPowerPlants")
	args := []interface{}{}

	query := `SELECT id , name, latitude, longitude FROM power_plant LIMIT ?,?`

	args = append(args, r.GetOffset(page, limit), limit)

	powerPlants = make([]*model.PowerPlant, 0)

	rows, err := r.Slave().QueryContext(ctx, query, args...)
	if err != nil {
		err = derrors.HandleSQLError(err, "QueryContext")
		return
	}

	for rows.Next() {
		wc := &model.PowerPlant{}
		err = rows.Scan(r.getDest(wc)...)
		if err != nil {
			return
		}

		powerPlants = append(powerPlants, wc)
	}

	totalCountQuery := `SELECT COUNT(*) FROM power_plant`
	// Get total count
	totalCountRow, err := r.Slave().QueryContext(ctx, totalCountQuery, []interface{}{}...)
	if err != nil {
		err = derrors.HandleSQLError(err, "QueryRowContext(%s)", totalCountQuery)
		return
	}

	var totalCount int
	for totalCountRow.Next() {
		err = totalCountRow.Scan(&totalCount)
		if err != nil {
			return
		}
	}

	return powerPlants, totalCount, nil
}
