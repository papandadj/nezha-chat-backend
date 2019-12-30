.PHONY: buildBinary buildImage restart

buildBinary:
	go build -o auth-srv/cmd ../auth-srv/cmd
	go build -o common-srv/cmd ../common-srv/cmd
	go build -o common-web/cmd ../common-web/cmd
	go build -o friend-srv/cmd ../friend-srv/cmd
	go build -o friend-web/cmd ../friend-web/cmd
	go build -o user-srv/cmd ../user-srv/cmd
	go build -o user-web/cmd ../user-web/cmd
	go build -o chat-srv/cmd ../chat-srv/cmd
	go build -o chat-web/cmd ../chat-web/cmd

buildImage:
	docker build -f ./dockerfile --tag nezha-chat/auth-srv:v1 --target auth-srv .
	docker build -f ./dockerfile --tag nezha-chat/common-srv:v1 --target common-srv .
	docker build -f ./dockerfile --tag nezha-chat/common-web:v1 --target common-web .
	docker build -f ./dockerfile --tag nezha-chat/friend-srv:v1 --target friend-srv .
	docker build -f ./dockerfile --tag nezha-chat/friend-web:v1 --target friend-web .
	docker build -f ./dockerfile --tag nezha-chat/user-srv:v1 --target user-srv .
	docker build -f ./dockerfile --tag nezha-chat/user-web:v1 --target user-web .
	docker build -f ./dockerfile --tag nezha-chat/chat-srv:v1 --target chat-srv .
	docker build -f ./dockerfile --tag nezha-chat/chat-web:v1 --target chat-web .


restart:
	go build -o chat-srv/cmd ../chat-srv/cmd
	docker build -f ./dockerfile --tag nezha-chat/chat-srv:v1 --target chat-srv .
	docker-compose up -d chat-srv
