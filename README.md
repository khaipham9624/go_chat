# go_chat
Go Chat

docker network create go_chat_net  
docker run --rm --name mongo --network go_chat_net -d mongo

docker build . -t go_chat  
docker run --rm --name go_chat --network go_chat_net -p 8080:8080 go_chat:latest 
