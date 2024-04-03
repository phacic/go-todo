SHELL=/bin/bash

%:
	@:

conn_str="postgres://nana-beng:vssta@0.0.0.0:5432/tester1?search_path=public&sslmode=disable"

# capture position arguments
args=$(filter-out $@,$(MAKECMDGOALS))

# generate entity (table)
entity:
	go run entgo.io/ent/cmd/ent new $(args)

# generate code needed
schema:
	go run entgo.io/ent/cmd/ent generate ./ent/schema

# inspect schema created with atlas
schema_inspect:
	atlas schema inspect -u $(conn_str) -w

# run this
# create main under ent/migrate
versioned-migration-schema:
	go run entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./ent/schema

versioned-migration:
	go run ./ent/migrate/main.go $(args)

migration:
	atlas migrate diff $(args) \
		--dir "file://ent/migrate/migrations" \
		--to "ent://ent/schema" \
		--dev-url $(conn_str)

migrate:
	atlas migrate apply \
		--dir "file://ent/migrate/migrations" \
		--url $(conn_str)

migrate-status:
	atlas migrate status \
		--url $(conn_str) \
		--dir "file://ent/migrate/migrations"

css_build:
	pnpm dlx tailwindcss -i ./static/css/input.css -o ./static/css/main.css

css_watch:
	pnpm dlx tailwindcss -i ./static/css/input.css -o ./static/css/main.css --watch

templ_generate:
	templ generate

run:
	air