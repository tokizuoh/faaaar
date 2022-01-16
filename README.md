# faaaar
Learn GraphQL with THE IDOLM@STER SHINY COLORS.

## Getting Started
The following is a simple example which get information about 20-year-old idols.
  
```bash
$ make build
...
Starting postgres-syani  ... done
Recreating faaaar-server ... done

$ make run
docker-compose exec server go mod download
docker-compose exec server go run main.go
2022/01/16 13:08:58 {"data":{"idols":[{"id":16,"name":"有栖川 夏葉"},{"id":26,"name":"斑鳩 ルカ"}]}} 

#
#query := `
#       {
#           idols(age: 20) {
#               id
#               name
#           }
#       }
#   `
#

```