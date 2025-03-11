# Golist Documentation
Golist is a whitelist made just in **Go** that uses the library crypt to secure your keys
> [!CAUTION]
> Is recommended using a host to substitute the folder Data to avoid expose the hashs directly inside the file [Data/Keys.json](https://github.com/RIOTLaF/Sentinel/blob/master/Data/Keys.json)

## Tree
```rust
├─── whitelist // root of the project
│   go.mod
│   go.sum
│   LICENSE
│   main.go // Init and repeat
│   README.md
│   whitelist.go // configure the whitelist
│
├───Data
│       config.json // General configuration
│       Keys.json // Database of the keys
│       // Format: {key: hwid}
├───gen
│       Salty.go // Salty generator
│
├───Handlers
│       identifier.go // Handle almost everything
│       json.go // Read and import json files in Data
│
└───phrase
        dictionary.go // Wordlist
        phrase.go // Word generator
        README // Warning
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
![Go](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white) <a href="https://github.com/lukewhrit/phrase" target="_blank">
    <img src="https://img.shields.io/badge/phrase-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Botão" width="100">
</a>

