# Data Structures in Go Language
Data Structures in Go to PhD discipline in Federal University of Ceará (UFC)
## Author
Wellington Wagner F. Sarmento

## Using VSCode

1. Install Go Extention (Go for Visual Studio Code)
2.  Use F1 e write "Go:Install/Update Tools". Select all tools
3. Create a task.json (Terminal -> Configure Tasks...)
``` json

{
    // See https://go.microsoft.com/fwlink/?LinkId=733558
    // for the documentation about the tasks.json format
    
    "version": "2.0.0",
    "tasks": [
        {
            "label": "Build and run",
            "type": "shell",
            "command": "rm -f  ${fileBasenameNoExtension} && go build ${fileBasenameNoExtension}.go && ./${fileBasenameNoExtension}",
            "group": {
                "kind": "build",
                "isDefault": true
            }
        }
    ]
}
```
4. Create a launch.json (Run -> Add Configuration)
``` json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387

    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch file",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${file}"
        },
        

    ]
}
```

