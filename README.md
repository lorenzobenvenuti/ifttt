# ifttt
A golang library to trigger IFTTT events.

Usage:

```g

import "ifttt"

// ...

iftttClient := ifttt.NewIftttClient(apiKey)
values := []string{"firstValue", "secondValue"}
iftttClient.Trigger(event, values)
```
