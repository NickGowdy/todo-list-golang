# todo-list-golang

To GET todos, run: 
```
curl localhost:8080/todos
```

Or to GET a single todo, run: 
```
curl localhost:8080/todos/1
```

To POST a todo, run: 

```
curl http://localhost:8080/todos \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"value": "A new TODO", "isComplete": false}'
```

To PUT a todo, run: 

```
curl http://localhost:8080/todos/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"id": 1, "value": "A new TODO 1111","isComplete": true}'
```

To delete a todo, run:

```
curl http://localhost:8080/todos/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "DELETE" \'
```
