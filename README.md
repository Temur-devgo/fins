# fins-api
Api service for fins

1) копируем файл env.example и сохраняем под новым именем .env
2) прописываем в нём свои локальные данные(такие же как и в другом проекте)
3) запускаем команду docker build -t fins-api-services .
4) запускаем команду docker run -p 8086:8080 fins-api-service 
