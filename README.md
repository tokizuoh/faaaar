# faaaar
Learn GraphQL with THE IDOLM@STER SHINY COLORS.

## Development Environment

```bash
> docker --version
Docker version 20.10.8, build 3967b7d

> docker-compose --version
docker-compose version 1.29.2, build 5becea4c
```

## Setup

```bash
# install task
# cf: https://taskfile.dev/#/installation
> sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

# build
> task build
```

## Usage

### Get idolList (by age)

```bash
> curl -H 'Content-Type:application/json' -X POST -d '{ idols(age:20) { id age name height birth_place birth_day blood_type } }' 'http://localhost:8080/graphql'

# 2022/01/29 15:31:55 {
#         "data": {
#                 "idols": [
#                         {
#                                 "age": 20,
#                                 "birth_day": "8/16",
#                                 "birth_place": "愛知県",
#                                 "blood_type": "B",
#                                 "height": 168,
#                                 "id": 16,
#                                 "name": "有栖川 夏葉"
#                         },
#                         {
#                                 "age": 20,
#                                 "birth_day": "1/31",
#                                 "birth_place": "神奈川県",
#                                 "blood_type": "A",
#                                 "height": 161,
#                                 "id": 26,
#                                 "name": "斑鳩 ルカ"
#                         }
#                 ]
#         }
# } 
```

### Get unitList (by idolId)

```bash
> curl -H 'Content-Type:application/json' -X POST -d '{ units(idolId: 2) { id name } }' 'http://localhost:8080/graphql'

# {
# 	"data": {
# 		"units": [
# 			{
# 				"id": "2",
# 				"name": "イルミネーションスターズ"
# 			},
# 			{
# 				"id": "2",
# 				"name": "Sol"
# 			}
# 		]
# 	}
# }
```
  
## Playground
 
Please visit the page `http://localhost:8081/playground`.
  
![](https://user-images.githubusercontent.com/37968814/154619250-3afdda0d-8610-496a-b66a-e45dc2465c1c.png)
  
## Development
  
### Add issue number to the commit message
  
1. Write the content of `./.git-hooks/commit-msg` in `.git/hooks/commit-msg`.
  
2. Grant permission to shell-script.  
  
```bash
$ chmod +x .git/hooks/commit-msg
```
   
