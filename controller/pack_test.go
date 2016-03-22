package controller

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

func TestPack(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetPacks", func() {
		var packs model.Packs

		g.BeforeEach(func() {
			packs = model.Packs{
				&model.Pack{
					Name: "Pack 3",
				},
				&model.Pack{
					Name: "Pack 1",
				},
				&model.Pack{
					Name: "Pack 2",
				},
			}

			for _, record := range packs {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.Pack{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetPack(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetPack(ctx)

			out := model.Packs{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(packs))
			g.Assert(out[0]).Equal(packs[2])
		})
	})
}