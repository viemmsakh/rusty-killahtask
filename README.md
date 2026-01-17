# killahtask
A simple TODO CLI tool for managing your task. This The cobra package was used to create the CLI interface. Task are written to a CSV file in following header structure `task_id, description, created, completed`. 

This CLI tool performs simple CRUD operations:
```bash
dilly@dilly:$ killahtask add "my really cool task"
dill@dilly:$ killahtask list
dilly@dilly:$ killahtask complete <task_id>
```

## Add
Running `killahtask add "my task"` will add a task to your to a CSV file in your home directory (not OS restricted). 


## List

## Completed
