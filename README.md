<h1 align="center"> WB Tech: level # 0 (Golang)</h1>

<p align="center">
  <img alt="Golang" src="https://img.shields.io/badge/Golang-74.6 %25-blue?style=for-the-badge&logo=appveyor">
  <img alt="Golang" src="https://img.shields.io/badge/HTML-14.3 %25-red?style=for-the-badge&logo=appveyor">
  <img alt="Golang" src="https://img.shields.io/badge/Makefile-11.1 %25-green?style=for-the-badge&logo=appveyor">
</p>
  

## **Тестовое задание**

**В БД**:
* Развернуть локально postgresql
* Создать свою бд
* Настроить своего пользователя.
* Создать таблицы для хранения полученных данных.

**В сервисе**:
* Подключение и подписка на канал в nats-streaming
* Полученные данные писать в Postgres
* Так же полученные данные сохранить in memory в сервисе (Кеш)
* В случае падения сервиса восстанавливать Кеш из Postgres
* Поднять http сервер и выдавать данные по id из кеша
* Сделать простейший интерфейс отображения полученных данных, для их запроса по id
***
## **Как работать с проектом:**

* `make`: make server_run

* `make server_run`: запустить натс стриминг, паблиш и сервис
* `make kill`: убить процессы, запущенные `make server_run`
* `make postgre`: создать базу данных wildberries и создать в ней таблицу **order_id**
* `make del_db`: удалить базу данных wildberries и таблицу **order_id**
* `make del_table`: удалить таблицу **order_id**
* `make table`: создать таблицу **order_id**
* `make show_table`: вывести содержимое таблицы **order_id**
* `make reset_table`: обнулить таблицу **order_id**
