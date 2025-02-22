# Steps

1. Crete the **app.go** and add the above code

```
package  main

import  "fmt"

import  "rsc.io/quote"

func  main() {

fmt.Println(quote.Go())

}
```

2. Generate the module file by ` go mod init example.com/first-app  `
3. Add module requirements and sums  ` go mod tidy ` 
4. Install externaml package ` go get rsc.io/quote ` 
5. Run the above code ` go run app.go`