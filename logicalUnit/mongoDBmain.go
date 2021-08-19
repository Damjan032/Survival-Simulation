package main

import (
	"./services"
	"fmt"
)

func main6() {
	/*	document := services.ReadLogs()
		services.WriteInMongoDB(document)*/
	/*foodSourcesDto := new(dto.FoodSourcesDto)
	foodSourcesDto.FoodSources = globals.FoodSources
	foodSourcesDto.CurrentEpoch = globals.CurrentEpoch
	foodSourcesDto.FinalEpoch = globals.FinalEpoch
	btResult, _ := json.Marshal(foodSourcesDto)
	fmt.Println(string(btResult))
	foodSourcesDto2 := new(dto.FoodSourcesDto)
	json.Unmarshal(btResult, foodSourcesDto2)
	btResult2, _ := json.Marshal(foodSourcesDto2)
	fmt.Println(string(btResult2))

	client, err := mongo.NewClient(options.Client().
		ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")*/

	//	episodesCollection := quickstartDatabase.Collection("episodes2")

	/*podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{"title", "The Polyglot Developer Podcast"},
		{"author", "Nic Raboy"},
		{"tags", bson.A{"development", "programming", "coding"}},
	})

	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "GraphQL for API Development"},
			{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Progressive Web Application Development"},
			{"description", "Learn about PWA development with Tara Manicsic."},
			{"duration", 32},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v documents into episode collection!\n", len(episodeResult.InsertedIDs))*/
	var mainParamsList = services.ReadAllMainParams()
	var fullDataList = services.ReadFullDataDocument()
	fmt.Println(len(mainParamsList))
	fmt.Println(len(fullDataList))
	fmt.Println(services.FindAllDataById("611e6bd6bc14806d31da0ac8"))
	//initRouter()
}
