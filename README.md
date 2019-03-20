<div align="center">
  <img src="logo.png" />
</div>

</br>

## osctrl
Fast and efficient operative system management.

## Dependencies

*(Presuming that you aready have Golang >= 1.11 installed in your system)*

The project uses [Go modules](https://github.com/golang/go/wiki/Modules) to manage dependencies. Make sure you have it installed, otherwise he easiest way to install the latest release on Mac or Linux is with the following script:

## Service configuration

The most basic `tls.json` configuration file will have the following format:

```json
{
  "tls": {
    "listener": "127.0.0.1",
    "port": "_TLS_PORT",
    "host": "_TLS_HOST",
    "auth": "none"
  },
  "db": {
    "host": "_DB_HOST",
    "port": "_DB_PORT",
    "name": "_DB_NAME",
    "username": "_DB_USERNAME",
    "password": "_DB_PASSWORD"
  },
  "logging": {
    "graylog": false,
    "graylogcfg": {
      "url": ""
    },
    "splunk": false,
    "splunkcfg": {
      "url": "",
      "token": "",
      "search": "results_for_{{NAME}}"
    },
    "stdout": false,
    "postgres": true
  }
}
```

And for `admin.json` it will look like this:

```json
{
  "admin": {
    "listener": "127.0.0.1",
    "port": "_ADMIN_PORT",
    "host": "_ADMIN_HOST",
    "auth": "local"
  },
  "db": {
    "host": "_DB_HOST",
    "port": "_DB_PORT",
    "name": "_DB_NAME",
    "username": "_DB_USERNAME",
    "password": "_DB_PASSWORD"
  },
  "logging": {
    "graylog": false,
    "graylogcfg": {
      "url": ""
    },
    "splunk": false,
    "splunkcfg": {
      "url": "",
      "token": "",
      "search": "results_for_{{NAME}}"
    },
    "stdout": false,
    "postgres": true
  }
}
```

## Using docker

## Using vagrant

Vagrant machines can be used for `osctrl` local development. Execute `vagrant up ubuntu` or `vagrant up centos` to create a local virtual machine running Ubuntu 18.04 or CentOS 7 respectively. Once it has finished deploying, `osctrl` will be ready to be used and you can access it following the instructions in the terminal.
