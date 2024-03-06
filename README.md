### BACKEND-сервис survey предварительного опроса кандидатов

Сервис survey публикует grpc-сервер с функциями:
1. Регистрация кандидата в БД (AddCandidate)
2. Удаление кандидата из БД (DeleteCandidate)
3. Старт опроса (StartSurvey)
4. Запись ответа в БД (SetAnswer)
5. Завершение опроса (SetFinishCandidate)
6. Получение результатов опроса кардидата (GetSurveyForCandidate)

Также, сервис survey публикует rest-сервер для frontend с функциями:
1. Старт опроса (StartSurvey)
2. Запись ответа в БД (SetAnswer)

Схема БД задана в sql-файлах миграции в директории migrations пароекта.

Сонфигурация сервиса задается в файле config.json (в директории cmd/survey/ проекта)

Документация по [REST API](./docs/api/swagger.md).

