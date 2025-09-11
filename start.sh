chmod +x /home/ubuntu/server/app
GIN_MODE=release ENV=production  EMO_URL='emo_guang:MyXinLog!1995@tcp(127.0.0.1:3306)/emo?charset=utf8mb4&parseTime=True&loc=Local' nohup  ./app &