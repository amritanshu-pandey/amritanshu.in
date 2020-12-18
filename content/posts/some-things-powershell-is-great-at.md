---
title: "Some things Powershell is great at"
date: 2020-12-17T22:30:00+11:00
draft: true
categories: ['powershell']
tags: ['powershell', 'json', 'cli']
---

Powershell is an Object based shell and scripting language. It is different from majority of other shells as instead of accepting and returning text, commands in Powershell accept and return objects. Each object has a textual representation that you see as output by default, but behind the scene, each output is an object that contains plethora of information (properties) and possible actions (methods). When two Powershell commands are piped, its not the text that is passed from first command to second, instead it is objects that are passed, hence the second command has a lot of properties and methods available to them.

Owing to the Object oriented nature of Powershell, it is similar to Python and its REPL, but being a shell, Powershell has tight integration with OS including the filesystem, executables, services etc.

In this article I just want to hightlight a few scenarios which are very easiely accomplished in powershell as compared to other shells.

## 1. Save output in Json, CSV or HTML format

Each command in powershell (cmdlet) returns an object. Let's see that in action.

If I type command `ls` in powershell, it returns the following output. 

```powershell
> ls

    Directory: C:\Users\amritanshu\Projects\amritanshu.in    

Mode                 LastWriteTime         Length Name       
----                 -------------         ------ ----       
d----          16/12/2020  9:12 PM                archetypes 
d----          16/12/2020  9:12 PM                content    
d----          16/12/2020  9:21 PM                public     
d----          16/12/2020  9:21 PM                resources  
d----          16/12/2020  9:12 PM                themes     
-a---          16/12/2020  9:12 PM              6 .gitignore 
-a---          16/12/2020  9:12 PM            112 .gitmodules
-a---          16/12/2020  9:12 PM            691 config.yaml
-a---          16/12/2020  9:12 PM            710 README.md  
```

`ls` is an alias to the powershell command `Get-ChildItem`. what we see on screen is just a textual representation, but behind the scenes it is an object. We can see the properties and method pertaining to the object using the `Get-Members` cmdlet (`Get-ChildItem | Get-Member`)

