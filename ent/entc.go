//go:build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)


func main() {
	spec := new(ogen.Spec)

	//oas, err := entoas.NewExtension(entoas.Spec(spec))
	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
			spec.AddPathItem("/users/{id}/card/start", ogen.NewPathItem().
			SetDescription("Start an draw as done").
			SetPatch(ogen.NewOperation().
			SetOperationID("drawStart").
			SetSummary("Draws a card item as done.").
			AddTags("Card").
			AddResponse("204", ogen.NewResponse().SetDescription("Item marked as done")),
		).
		AddParameters(ogen.NewParameter().
		InPath().
		SetName("id").
		SetRequired(true).
		SetSchema(ogen.Int()),
	),
)
return nil
								}),

//								entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
//									spec.AddPathItem(
//										"/syui",
//										ogen.NewPathItem().
//										SetGet(
//											ogen.NewOperation().
//											SetOperationID("customReadUser").
//
//			AddResponse("204", ogen.NewResponse().SetDescription("Item marked as done")),
//										),
//									)
//						return nil
//					}),
								entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
									spec.AddPathItem("/cards/{id}/d", ogen.NewPathItem().
									SetDescription("Start an draw as done").
									SetPut(ogen.NewOperation().
									SetOperationID("drawDone").
									SetSummary("Draws a card item as done.").
									AddTags("Card").
									AddResponse("204", ogen.NewResponse().SetDescription("Item marked as done")),
									//AddResponse("204", ogen.NewResponse().SetDescription("Item marked as done").SetSchema("test")),
								).
								AddParameters(ogen.NewParameter().
								InPath().
								SetName("id").
								SetRequired(true).
								SetSchema(ogen.Int()),
							),
						)
						return nil
					}),
				)

				if err != nil {
					log.Fatalf("creating entoas extension: %v", err)
				}
				ogent, err := ogent.NewExtension(spec)
				if err != nil {
					log.Fatalf("creating ogent extension: %v", err)
				}
				err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
				if err != nil {
					log.Fatalf("running ent codegen: %v", err)
				}
			}
