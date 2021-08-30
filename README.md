### Make Your Whatsapp Messages Schedule

<a href="https://github.com/snowron/gowp-scheduler/commits/master">
  <img
    src="https://img.shields.io/github/last-commit/snowron/gowp-scheduler.svg?style=flat-square&logo=github&logoColor=white"
    alt="GitHub last commit"
  />
</a>

<a href="https://github.com/snowron/gowp-scheduler/issues">
  <img
    src="https://img.shields.io/github/issues-raw/snowron/gowp-scheduler.svg?style=flat-square&logo=github&logoColor=white"
    alt="GitHub issues"
  />
</a>

<a href="https://github.com/snowron/gowp-scheduler/pulls">
  <img
    src="https://img.shields.io/github/issues-pr-raw/snowron/gowp-scheduler.svg?style=flat-square&logo=github&logoColor=white"
    alt="GitHub pull requests"
  />
</a>

# About

Package has two functional actions.

* Instant
* Schedule

# Usage

The whatsapp qr code will be on your terminal. So easily can scan and then start use

## Instant Mode

``
go run . -type instant -contactsPath ./service/test-data/contacts.csv
``

The command runs the package with instant mode and next step is typing message.

````
Enter Message: Test message from cli
Test message from cli
Are You Sure? [y/N] :
````

Your csv file should be like mentioned below. Actually package doesn't use name field but sometimes saves us from typo
or mismatch problems

```
murat,90539xxxxxxx
turan,90539xxxxxxx
```

## Schedule Mode

``
go run . -type schedule -ordersPath ./service/test-data/orders.csv
``

Your csv file should be like mentioned below. The timer set every 10 seconds.

```
90539,hello from csv,2022-10-13T12:12:00Z
90539,hello from go,2022-10-13T12:12:00Z
```