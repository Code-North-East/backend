# LogHelper 

v1.0.0

## Concept

A common utlity library which can be used in all the present modules of the project for standardizing logger.

## Components

Every log should have the following components:

- Level of the log.
- Code occurence (where exactly the log is originating from).
- Operation code (this is a short code which makes it easier to query).
- Descriptive messge of the log. 


## Structure 

### Basic INFO level using zap

```
{"level":"info","ts":1779819441.7318513,"caller":"logger/demo_logger.go:14","msg":"injecting an info level log"}```
```

