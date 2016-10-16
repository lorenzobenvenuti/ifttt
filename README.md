# ifttt
A golang library to trigger IFTTT events.

Usage:

```go

import "ifttt"

// ...

iftttClient := ifttt.NewIftttClient(apiKey)
values := []string{"firstValue", "secondValue"}
iftttClient.Trigger(event, values)
```
