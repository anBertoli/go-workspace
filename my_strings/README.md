### Module `my_strings`
The module is NOT listed in the `go.work` file. It has some published versions (e.g. 
my_strings/v0.0.1). Inside the workspace other modules doesn't import this local copy 
(because my_strings is NOT present in the workspace).

To test, build or run a module inside the workspace directory, but not part of the workspace
(not listed in the `go.work` file) we must disable the workspace mode via:

```bash
GOWORK=off go <command> 
```