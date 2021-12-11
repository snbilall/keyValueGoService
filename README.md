#### Key Value project with GoLang

1. [Clone the repo](#installation)
2. [Copy these files into the new project](#installation)
3. Update the README, replacing the contents below as prescribed.
4. Add any libraries, assets, or hard dependencies whose source code will be included
   in the project's repository to the _Exceptions_ section in the [TERMS](TERMS.md).
  - If no exceptions are needed, remove that section from TERMS.
5. If working with an existing code base, answer the questions on the [open source checklist](opensource-checklist.md)
6. Delete these instructions and everything up to the _Project Title_ from the README.
7. Write some great software and tell people about it.

> Keep the README fresh! It's the first thing people see and will make the initial impression.

## Installation

To install all of the template files, run the following script from the root of your project's directory:

```
git clone https://github.com/snbilall/keyValueGoService.git
```

----

# Project Title

**Description**:  A Basic key value project. Gets key value values from client, saves data on memory. It saves memory data to a file in periodic interval. Gin library used to build web service.

## Dependencies

gin v1.7.7

## Installation

Just run code below in root directory

```bash
go run .
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

## Getting help

Contact with me.

**Example**

----

## Open source licensing info
1. [TERMS](TERMS.md)
2. [LICENSE](LICENSE)
3. [CFPB Source Code Policy](https://github.com/cfpb/source-code-policy/)