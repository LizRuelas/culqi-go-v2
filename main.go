package main

import (
  "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/adaptors/httprouter"
  "gopkg.in/kataras/iris.v6/adaptors/view"
  "fmt"

  "github.com/culqi/culqi-go"
  "github.com/culqi/culqi-go/charge"

)



func main() {
  app := iris.New()
  app.Adapt(iris.DevLogger())
  app.Adapt(httprouter.New())

  tmpl := view.HTML("./views", ".html")
  tmpl.Layout("index.html")


  app.Adapt(tmpl)



  app.Get("/", func(ctx *iris.Context) {
    ctx.Render("index.html", iris.Map{"gzip": true})
    // Note that: you can pass "layout" : "otherLayout.html" to bypass the config's Layout property
    // or iris.NoLayout to disable layout on this render action.
    // third is an optional parameter
  })

  app.Post("/cargo", func(ctx *iris.Context) {
    // 1. Configuración
     config := &culqi.Config{
       MerchantCode:   "pk_test_Rp2uV5dXI3quFq2X",  // Llave publica
       ApiKey:         "sk_test_8GC9UJfifciOurwW", // Llave privada
     }

     // 2. Crea un nuevo cliente
     client := culqi.New(config)

     // 3. Parametros de creación de cargo
    params := &charge.ChargeParams{
      TokenId: app.Post[token],
      CurrencyCode: "PEN",
      Amount: 100,
      Email: "liz@ruelas.com",
    }

    // 4. Crear Cargo
    resp, err := charge.Create(params, client)

    if err != nil {
        panic(err.Error())
    }
  })

  app.StaticWeb("/static", "./assets")


  app.Listen(":8084")
}
