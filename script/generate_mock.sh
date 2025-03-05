#!/bin/sh

# Generate mocks for repository interfaces
mockery --name=PowerPlantRepository --dir=internal/repository/power_plant --output=internal/test/mockrepository --outpkg=mockrepository

# Generate mocks for usecase interfaces
mockery --name=PowerPlantUsecase --dir=internal/usecase/power_plant --output=internal/test/mockusecase --outpkg=mockusecase