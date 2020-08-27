# Fx Medium Example

Full blog post can be found here https://medium.com/@erez.levi/using-uber-fx-to-simplify-dependency-injection-875363245c4c

## To Install
    go get github.com/ErezLevip/fx-medium-example

## To Run

    go run main.go
    
## TL:DR
The main looks like this:

    func main() {
	  ctx, cancel := context.WithCancel(context.Background())
	  kill := make(chan os.Signal, 1)
	  signal.Notify(kill)

	  go func() {
		  <-kill
		  cancel()
	  }()

	  app := fx.New(
		  fx.Provide(newZapLogger),
		  fx.Provide(newRedisClient),
		  fx.Provide(cache.NewMeaningOfLifeCacheRedis),
		  fx.Provide(handlers.NewMeaningOfLifeHandler),
		  fx.Invoke(runHttpServer),
	  )
	  if err := app.Start(ctx); err != nil {
		  fmt.Println(err)
	  }
    }

    func runHttpServer(lifecycle fx.Lifecycle, molHandler *handlers.MeaningOfLife) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		r := fasthttprouter.New()
		r.Handle(http.MethodGet, "/what-is-the-meaning-of-life", molHandler.Handle)
		return fasthttp.ListenAndServe("localhost:8080", r.Handler)
	}})}
