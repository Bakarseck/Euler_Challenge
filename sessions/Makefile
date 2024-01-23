# Codes de couleur ANSI
GREEN := \033[32m
RED := \033[31m
RESET := \033[0m
EXE := forum
DB_PATH := ./db/sessions.db
SQL_FILE := ./db/sessions.sql
EXEC_FILE := ./db/exec.sql
POST_FILE := ./db/posts.sql
CATEGORY_FILE := ./db/categories.sql

start:
	go run .

# Target to Build the forum image and start a container
up:
	@echo "\n\n\t\t\t\t$(GREEN)Build The Docker Image...$(RESET)\n\n"
	docker build -t $(EXE) .
	docker run -it --name $(EXE) -p 8080:8080 $(EXE)
	@echo "Docker images started!"

clean:
	@echo "\n\n\t\t\t\t$(RED)Stopping and Removing Containers...$(RESET)\n\n"
	docker stop $(EXE)
	docker rm $(EXE)
	@echo "Containers and artifacts removed!"

rmproper: clean
	@echo "\n\n\t\t\t\t$(RED)Stopping and Removing Containers...$(RESET)\n\n"
	docker rmi $(EXE)
	@echo "Containers and artifacts removed!"

create-db:
	@sqlite3 $(DB_PATH) < $(SQL_FILE)

exec:
	@sqlite3 $(DB_PATH) < $(EXEC_FILE)

post:
	@sqlite3 $(DB_PATH) < $(POST_FILE)

category:
	@sqlite3 $(DB_PATH) < $(CATEGORY_FILE)
