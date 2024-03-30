# testovoe_20_03_24
##### Проект предлагает выдачу информации по заказам.

### Для удобства весь проект можно запустить за несколько шагов
#### Для этого нужен Docker, golang, migrate, sqlc

1. **Установка sqlc**
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

2. **Установка migrate**
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

3. **Запуск программы**
```bash
make docker
```
```bash
make migrateup
```
```bash
make sqlc
```
#### Также скопируйте запросы с файла db/insert_base_data.sql и добавте в бд

```bash
make run
```

## Иерархия

- **cmd**
  - **main.go** (Entry point)
- **db**
  - **migration** (Миграции)
  - **query** (SQL который будет в дальнейшем сгенерирован в golang)
  - **sqlc** (Сгенерированные sql запросы, а также вспомогательные функции для более сложных запросов в бд)
- **internal**
  - **app** (Основное логика)
  - **util** (Дополнительные модули)

### Доп. ссылки
Схема БД: https://shorturl.at/iqrsV

### Что люди говорят об этом проекте:
> Невероятно крутая работа, возьмите его за 300К/НС. © Мой друг Рома (Первая личность)<br>

> Невероятно отвратительная работа, не берите его и за покушатц. © Мой друг Рома (Вторая личность)<br>