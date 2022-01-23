# faaaar
Learn GraphQL with THE IDOLM@STER SHINY COLORS.

## Getting Started
The following is a simple example which get information about 20-year-old idols.
  
### Query

```bash
query := `
	{
		idols {
			id
			name
			age
			height
			birth_place
			birth_day
			blood_type
		}
	}
`
```

### Build & Run

```bash
$ make build
...
Starting postgres-syani  ... done
Recreating faaaar-server ... done

$ make run
docker-compose exec server go mod download
docker-compose exec server go run main.go
2022/01/23 06:08:34 {
        "data": {
                "idols": [
                        {
                                "age": 16,
                                "birth_day": "4/25",
                                "birth_place": "東京都",
                                "blood_type": "A",
                                "height": 155,
                                "id": 1,
                                "name": "櫻木 真乃"
                        },
                        {
                                "age": 16,
                                "birth_day": "7/22",
                                "birth_place": "アメリカ マサチューセッツ州",
                                "blood_type": "O",
                                "height": 157,
                                "id": 2,
                                "name": "八宮 めぐる"
                        },
                        ...
                        {
                                "age": 20,
                                "birth_day": "1/31",
                                "birth_place": "神奈川県",
                                "blood_type": "A",
                                "height": 161,
                                "id": 26,
                                "name": "斑鳩 ルカ"
                        }
                ]
        }
} 
```
  
## Development
  
### Add issue number to the commit message
  
1. Write the content of `./git-hooks/commit-msg` in `.git/hooks/commit-msg`.
  
2. Grant permission to shell-script.  
  
```bash
chmod +x .git/hooks/commit-msg
```
   