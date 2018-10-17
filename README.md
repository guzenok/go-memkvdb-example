# key-value in-memory database

тестовое задание


### Задание

Необходимо реализовать hot key-value хранилище в памяти без использования субд, а с использованием map.
Основные требования:
1. Данным объектом должен управлять 1 сервис.
2. Количество read/write-ов может быть любым, поэтому надо реализовать многопоточность.
3. Доступ должен быть потокобезопасным (map+mutex, либо sync.Map).
4. При успешном считывании по ключу надо получить значение, а сама запись должна удаляться из данного объекта.
5. Необходимо реализовать автоочистку объекта, несчитанные записи надо удалять из структуры по истечению 30 секунд.
6. Удаление должно происходить с использованием каналов, а не с помощью итерации и поиска expired записей.
7. Большим плюсом будет решение с использованием context.WithTimeout, а не time.After.


### Установка

`go get github.com/guzenok/go-memkvdb-example`


### Использование

см. [полный пример](examples/console/main.go)


```
import (
    db "github.com/guzenok/go-memkvdb-example"
)
```
```
    memcache, err := db.New(30*time.Second, db.CreateMapStore())
    if err != nil {
        panic(err)
    }

    key := []byte("index")
    val := []byte("stored data")

    err = memcache.Set(key, val)

    val, err = memcache.Get(key)
```


### Тесты

см. [Makefile](Makefile)

`make test`
`make bench`