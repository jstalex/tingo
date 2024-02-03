build:
	docker build -t robot:local --file ./cmd/robot/Dockerfile .
	docker build -t migrate:local --file ./migrate/Dockerfile .
	docker build -t candles:local --file ./cmd/candles_downloader/Dockerfile .
	docker build -t backtest:local --file ./cmd/backtest/Dockerfile .
	docker-compose build
run: build
	docker-compose up -d
stop:
	docker-compose down