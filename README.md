# go-blockchain
![golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
<br /><br /><br />
Basic implementation of blockchain using golang.

### Initialize Blockchain
The application will automatically creates genesis block if the local blocks data are not available. All blocks data will be stored in `tmp/blocks`

### Database
This blockchain use `BadgerDB` which is a key-value store

### Basic usage

### Printing Blockchain
```
go run main.go print
```

### Add Block
```
go run main.go add -block <data: as a string>
```
