# JIRA

A JIRA command line utility for easily opening your issues from the terminal.

## Usage

Create the following file ~/.jira/config.json:

  ```json
  {
    "host": "https://myjira.com"
  }
  ```

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
