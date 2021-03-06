## Installation

To install all of the template files, run the following script from the root of your project's directory:

```
git clone https://github.com/snbilall/keyValueGoService.git
```

----

# Key Value Go Service

**Description**:  A Basic key value project. Gets key value values from client, saves data on memory. It saves memory data to a file in periodic interval. Gin library used to build web service.

## Dependencies

gin v1.7.7

## Installation

Just run code below in root directory

```bash
go run .
```

If you want to build app rn following commands
```bash
mkdir ./bin
go build -o .\bin\
```

## Configuration

```json
{
    "interval": 5,
    "fileName": "data.json"
}
```

Interval config is the interval of saving key value data in memory into a json file in seconds.
Filename config is the file name which will save to os temp directory.

## How to test the software

Go to tests directory and run bash script below

```bash
go test
```

## Usage

### Locally

1. Clone repository:

```
git clone https://github.com/snbilall/keyValueGoService.git

2. Install project mod packages. In Terminal, from project root, run:
```bash
go run .
```

3. Install Postman app to your computer, run postman and import [go.postman_collection.json](https://github.com/snbilall/keyValueGoService/blob/master/go.postman_collection.json) file. open request with name not presence of 'heroku', send requests. enjoy!

### On Heroku

1. Install Postman app to your computeri run postman and import [go.postman_collection.json](https://github.com/snbilall/keyValueGoService/blob/master/go.postman_collection.json) file. open request with name presence of 'heroku', send requests. enjoy!

## Getting help

Contact with me.

## Future Work

1. Working with db
2. Containerizing the app
3. Send logs kibana or something else
4. Make logging infasctucture better

## Open source licensing info
1. [TERMS](TERMS.md)
2. [LICENSE](LICENSE)
3. [CFPB Source Code Policy](https://github.com/cfpb/source-code-policy/)