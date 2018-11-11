package app

import (
	"flag"
	"github.com/RichardKnop/uuid"
	"github.com/d7561985/opt_stream/models"
	"github.com/icrowley/fake"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/rs/zerolog/log"
	"io"
	"math/rand"
	"time"
)

var (
	app  *iris.Application
	stor = [10]uuid.UUID{}
)

func init() {
	for i := range stor {
		stor[i] = uuid.NewRandom()
	}
}

// Initialize application
func Initialize() {

	app = iris.New()

	// CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://bo.1pt.gcsd.me", "https://localhost:3000"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		// should contain all supported
		AllowedMethods:     []string{"GET", "DELETE", "POST", "PUT"},
		OptionsPassthrough: true,
		// Debug:true,
	})

	app.Use(crs)
	app.Logger().SetLevel("debug")

	subRouter := app.Party("/", crs, func(context iris.Context) {
		// hack for OPTIONS, no need handle options method.
		if context.Request().Method != "OPTIONS" {
			context.Next()
			return
		}
		context.StatusCode(iris.StatusNoContent)
	}).AllowMethods(iris.MethodOptions)

	subRouter.Get("/", func(context iris.Context) {
		for {
			context.Header("Transfer-Encoding", "chunked")
			context.ContentType("application/json")
			context.StreamWriter(func(w io.Writer) bool {

				count := rand.Intn(10)
				res := make([]models.Work, count)
				for i := range res {
					res[i] = models.Work{
						ID:   stor[rand.Intn(10)],
						Name: fake.FirstName(),
						Hum:  fake.Latitude(),
						Temp: rand.Float32() * 100.0,
					}
				}

				context.JSON(&models.Request{Data: res})
				context.ResponseWriter().Flush()

				t := time.After(time.Second * 5)
				<-t
				return true
			})
		}
	})

	// recover usage enable
	app.Use(recover.New())
}

// Prepare custom tasks at start app
// not use in on test as this is integration part
func Prepare() {

}

// Run application
func Run() {
	if app == nil {
		log.Panic().Msg("no initialization")
	}

	P := flag.String("port", ":8081", ":8081")
	flag.Parse()

	app.Run(
		iris.Addr(*P),
		iris.WithOptimizations,
	)
}

// GetApp take iris.Application instance
func GetApp() *iris.Application {
	return app
}
