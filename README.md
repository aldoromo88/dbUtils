
# Go DbUtils

Simple utility functions for handle Db operations in Go. **Not a Db Driver**

## Installing

Using get
```
go get github.com/aldoromo88/dbUtils
```

Using dep

```
dep ensure -add github.com/aldoromo88/dbUtils
```


## Usage


```
func GetAllProducts(ctx *fasthttp.RequestCtx) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, _ := pq.Open(connStr)
	defer db.Close()
	q, _ := db.Prepare(`SELECT * FROM public.product`)
	r, _ := q.Query(nil)


	js, err := dbUtils.ParseAsJson(r)

	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(js)
}
```
