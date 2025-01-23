Basic TODO app to learn how to build a rest api

- get item 
  - For all items using Invoke-WebRequest `curl http://localhost:t8080/task/all`
  - For specific items using Invoke-WebRequest `curl http://localhost:t8080/task/{itemId}`
  
- add item 
  - Using Invoke-WebRequest `curl -Method Post -Body '{"Detail":"1nd task"}' http://localhost:8080/task/`
- delete item

