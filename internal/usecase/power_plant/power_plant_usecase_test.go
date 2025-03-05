package powerplantusecase_test

import (
	"context"
	"tensor-graphql/internal/model"
	"tensor-graphql/internal/test"
	powerplantusecase "tensor-graphql/internal/usecase/power_plant"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type params struct {
	PowerPlant *model.PowerPlant
}

func TestCreatePowerPlant(t *testing.T) {
	mc := test.InitMockComponent(t)
	ctx := context.Background()
	testUsecase := powerplantusecase.NewPowerPlantUsecase(mc.PowerPlantRepository)

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(err error)
	}{
		{
			caseName: "CreatePowerPlant_Success",
			params: params{
				&model.PowerPlant{
					ID:        "1",
					Name:      "test_name",
					Latitude:  1.0,
					Longitude: 1.0,
				},
			},
			expectations: func(params params) {
				mc.PowerPlantRepository.On("CreatePowerPlant", mock.Anything, mock.Anything, params.PowerPlant).
					Return(nil)
			},
			results: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			caseName: "CreatePowerPlant_Error",
			params: params{
				&model.PowerPlant{
					ID: "2",
				},
			},
			expectations: func(params params) {
				mc.PowerPlantRepository.On("CreatePowerPlant", mock.Anything, mock.Anything, params.PowerPlant).
					Return(assert.AnError)
			},
			results: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			err := testUsecase.CreatePowerPlant(ctx, testCase.params.PowerPlant)
			testCase.results(err)
		})
	}
}

func TestGetPowerPlantByID(t *testing.T) {
	mc := test.InitMockComponent(t)
	ctx := context.Background()
	testUsecase := powerplantusecase.NewPowerPlantUsecase(mc.PowerPlantRepository)

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(powerPlant *model.PowerPlant, err error)
	}{
		{
			caseName: "GetPowerPlantByID_Success",
			params: params{
				&model.PowerPlant{
					ID:        "1",
					Name:      "test_name",
					Latitude:  1.0,
					Longitude: 1.0,
				},
			},
			expectations: func(params params) {
				mc.PowerPlantRepository.On("GetPowerPlantByID", mock.Anything, params.PowerPlant.ID).
					Return(params.PowerPlant, nil)
			},
			results: func(powerPlant *model.PowerPlant, err error) {
				assert.NotNil(t, powerPlant)
				assert.NoError(t, err)
			},
		},
		{
			caseName: "GetPowerPlantByID_Error",
			params: params{
				&model.PowerPlant{
					ID: "2",
				},
			},
			expectations: func(params params) {
				mc.PowerPlantRepository.On("GetPowerPlantByID", mock.Anything, params.PowerPlant.ID).
					Return(nil, assert.AnError)
			},
			results: func(powerPlant *model.PowerPlant, err error) {
				assert.Nil(t, powerPlant)
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			powerPlant, err := testUsecase.GetPowerPlantByID(ctx, testCase.params.PowerPlant.ID)
			testCase.results(powerPlant, err)
		})
	}
}

func TestUpdatePowerPlant(t *testing.T) {
	mc := test.InitMockComponent(t)
	ctx := context.Background()
	testUsecase := powerplantusecase.NewPowerPlantUsecase(mc.PowerPlantRepository)

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(err error)
	}{
		{
			caseName: "UpdatePowerPlant_Success",
			params: params{
				&model.PowerPlant{
					ID:        "1",
					Name:      "test_name",
					Latitude:  1.0,
					Longitude: 1.0,
				},
			},
			expectations: func(params params) {
				mc.PowerPlantRepository.On("UpdatePowerPlant", mock.Anything, mock.Anything, params.PowerPlant).
					Return(nil)
			},
			results: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			caseName: "UpdatePowerPlant_Error",
			params: params{
				&model.PowerPlant{
					ID: "2",
				},
			},
			expectations: func(params params) {
				mc.PowerPlantRepository.On("UpdatePowerPlant", mock.Anything, mock.Anything, params.PowerPlant).
					Return(assert.AnError)
			},
			results: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			err := testUsecase.UpdatePowerPlant(ctx, testCase.params.PowerPlant)
			testCase.results(err)
		})
	}
}

func TestGetPowerPlants(t *testing.T) {
	mc := test.InitMockComponent(t)
	ctx := context.Background()
	testUsecase := powerplantusecase.NewPowerPlantUsecase(mc.PowerPlantRepository)

	var testCases = []struct {
		caseName     string
		page, limit  int
		expectations func(page, limit int)
		results      func(powerplants []*model.PowerPlant, total int, err error)
	}{
		{
			caseName: "GetPowerPlants_Success",
			page:     1,
			limit:    10,
			expectations: func(page, limit int) {
				mc.PowerPlantRepository.On("GetPowerPlants", mock.Anything, page, limit).
					Return([]*model.PowerPlant{
						{
							ID:        "1",
							Name:      "test_name",
							Latitude:  1.0,
							Longitude: 1.0,
						},
					}, 1, nil)
			},
			results: func(powerplants []*model.PowerPlant, total int, err error) {
				assert.NotNil(t, powerplants)
				assert.Equal(t, 1, total)
				assert.NoError(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.page, testCase.limit)
			powerplants, total, err := testUsecase.GetPowerPlants(ctx, testCase.page, testCase.limit)
			testCase.results(powerplants, total, err)
		})
	}
}
