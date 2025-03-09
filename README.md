# Golist Documentation
Golist is a whitelist made just in **Go** that uses the library crypt to secure your keys
> [!WARNING]
> This is not  secure if you are gonna use delete the folder "Data" and use a online server to handle the keys storage

## Tree
```YAML
Data:
    - config.json # Defines what the whitelist will use (HWID / IP Adress)
    - Keys.json # The keys itself in format "key": "identifier"
         # identifier = ip or HWID
Handlers:
     - identifier.go # Get the info of the user
     - json.go # JSON Handler
```

## Functions
the main whitelist functions are get by the file "whitelist.go" found in the root of the project
### VerifyKey
VerifyKey checks if the key is in **Keys.json** and if it matches with the user data provided

Example:
```go
if VerifyKey("User_Key_Here") {
    fmt.Println("Sucess")
}
```
## Made using <3
![Go](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
