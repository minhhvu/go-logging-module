#Go Logging Module

##Current version: 1.0

##Introduction
This Go lang module was developed as a standard logging module in the go project. The log will be outputted as the same
as the api-logging-framework (in Java). 

For the consistent logging system in Lambda CloudWatch, developers should use this log module to log at key points. It 
is important to provide the log with the background information (e.g project name, environment, ...) and the level of 
log. The **log level** can be one of 6 types: info, debug, trace, warn, error, and fatal.

By default, the **output format** of log is set up as JSON type. However, it can be configured to FLAT type.

##How to Use
To import this module, developers must first install the module by configuring on the go.mod file

Next, you import the module in your Go file and name the module to user it later, using this statement:

```
import logger "go-logging-module"
```

After that, you can now use the log module. To get an idea on how to fully utilize this framework, you can
take a look at the sample go module in this repository called **logger-example**.


