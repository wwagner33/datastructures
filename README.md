# datastructures
Data Structures in Go to PhD discipline in Federal University of Ceará (UFC)

## Using VSCode

a) Install Go Extention (Go for Visual Studio Code)
b)  Use F1 e write "Go:Install/Update Tools". Select all tools
c) Create a task.json (Terminal -> Configure Tasks...)
``` json
    // See https://go.microsoft.com/fwlink/?LinkId=733558
    // for the documentation about the tasks.json format
{
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
d) Create a launch.json (Run -> Add Configuration)
``` json
    // See https://go.microsoft.com/fwlink/?LinkId=733558
    // for the documentation about the tasks.json format
{
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

