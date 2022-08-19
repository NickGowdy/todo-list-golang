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
    --data '{"id": 4,"name": "A new TODO","description": "Some more details about TODO","isComplete": false}'
```

To PUT a todo, run: 

```
curl http://localhost:8080/todos/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"name": "A new TODO","description": "Some more details about TODO","isComplete": true}'
```

To delete a todo, run:

```
curl http://localhost:8080/todos/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "DELETE" \'
```
