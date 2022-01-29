# faaaar
Learn GraphQL with THE IDOLM@STER SHINY COLORS.

## Setup

```bash
$ make build
$ make run
```

## Usage

```bash
# get idolList (by age)
$ curl -H 'Content-Type:application/json' -X POST -d '{ idols(age:20) { id age name height birth_place birth_day blood_type } }' 'http://localhost:8080/graphql'

# get unitList (by idolId)
$ curl -H 'Content-Type:application/json' -X POST -d '{ units(idolId: 2) { id name } }' 'http://localhost:8080/graphql'
```
  
## Development
  
### Add issue number to the commit message
  
1. Write the content of `./git-hooks/commit-msg` in `.git/hooks/commit-msg`.
  
2. Grant permission to shell-script.  
  
```bash
$ chmod +x .git/hooks/commit-msg
```
   