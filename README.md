# todo-list-golang

To get todos, run: 
```
curl localhost:8080/todos
```

To post a todo, run: 

```
curl http://localhost:8080/todos \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": 4,"name": "A new TODO","description": "Some more details about TODO","isComplete": false}'
```
