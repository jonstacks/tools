# JIRA

A JIRA command line utility for easily opening your issues from the terminal.

- [Usage](#usage)
  - [Config](#config)
- [Basic](#basic)
- [Print Description](#print-description)
## Usage

### Config
Create the following file ~/.jira/config.json:

  ```json
  {
    "host": "https://myjira.com",
    "user_email": "my.email@domain.com",
    "api_token": "abc123"
  }
  ```

## Basic

Now, if you name your branch after the current issue you are working on, you
can simply do:

```sh
$ jira
```

to open up the JIRA issue you are currently working on. If you would like to
open up a different issue, you can also do that by specifying it like so:

```sh
$ jira ABC-1234
```

## Print Description

You may also wish to print the description in markdown format for easily including
in a GH issue or PR. This sub command requires a user email and API token in your
config.

```sh
$ jira description
```

or

```sh
$ jira description ABC-1234
```