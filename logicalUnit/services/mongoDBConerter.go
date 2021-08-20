package services

import (
	"../const"
	"../dto"
	"../globals"
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func initMongoDocumentDto(firstLine string) *dto.MongoDocumentDto {
	mongoDocument := new(dto.MongoDocumentDto)
	splitForDateAndTag := strings.Split(strings.Split(firstLine, "]")[0], "[")
	date := strings.Trim(splitForDateAndTag[0], "\t \n")
	info := splitForDateAndTag[1]
	if info != "INFO" {
		return nil
	}
	finalEpochString := strings.Split(regexp.MustCompile(`FinalEpochs: .[0-9]*`).FindString(firstLine), ":")[1]
	finalEpoch, _ := strconv.Atoi(strings.Trim(finalEpochString, "\t \n"))

	numberOfFoodSourcesString := strings.Split(regexp.MustCompile(`NumberOfFoodSources: .[0-9]*`).FindString(firstLine), ":")[1]
	numberOfFoodSources, _ := strconv.Atoi(strings.Trim(numberOfFoodSourcesString, "\t \n"))

	lvlString := strings.Split(regexp.MustCompile(`Level: .*`).FindString(firstLine), ":")[1]
	lvl := strings.Trim(lvlString, "\t \n")
	var id = primitive.NewObjectID()
	mongoDocument.Id = id
	mongoDocument.MainParams.Id = id
	mongoDocument.MainParams.FinalEpoch = finalEpoch
	mongoDocument.MainParams.NumberOfFoodSources = numberOfFoodSources
	mongoDocument.MainParams.StartDate = date
	mongoDocument.MainParams.Level = lvl
	readEpoch(firstLine, mongoDocument)
	//finalEpoch, _ := strconv.ParseInt(finalEpochString, 10, 64)
	fmt.Println(finalEpoch)
	//fmt.Println(int(strings.Trim(finalEpohString, "\t \n")))

	fmt.Println(info)
	return mongoDocument
}

func readEpoch(line string, documentDto *dto.MongoDocumentDto) {
	epoch := new(dto.Epoch)
	numberOfGoodsString := strings.Split(regexp.MustCompile(`NumberOfGoods: .[0-9]*`).FindString(line), ":")[1]
	numberOfGoods, _ := strconv.Atoi(strings.Trim(numberOfGoodsString, "\t \n"))

	numberOfBadsString := strings.Split(regexp.MustCompile(`NumberOfBads: .[0-9]*`).FindString(line), ":")[1]
	numberOfBads, _ := strconv.Atoi(strings.Trim(numberOfBadsString, "\t \n"))

	currentEpochString := strings.Split(regexp.MustCompile(`CurrentEpoch: .[0-9]*`).FindString(line), ":")[1]
	currentEpoch, _ := strconv.Atoi(strings.Trim(currentEpochString, "\t \n"))

	epoch.CurrentEpoch = currentEpoch
	epoch.NumberOfGood = numberOfGoods
	epoch.NumberOfBad = numberOfBads

	documentDto.Epoches = append(documentDto.Epoches, epoch)
}

func ReadLogs() *dto.MongoDocumentDto {
	file, err := os.Open(_const.LogFilePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	mongoDocument := initMongoDocumentDto(scanner.Text())
	fmt.Println(scanner.Text())
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		readEpoch(scanner.Text(), mongoDocument)
	}
	return mongoDocument
}

func connectDB() {
	if globals.MongoDBConnection == nil {
		client, clientError := mongo.NewClient(options.Client().
			ApplyURI(_const.MongoDataBaseUrl))
		if clientError != nil {
			log.Fatal(clientError)
		}
		ctx, cancelFun := context.WithTimeout(context.Background(), 10*time.Second)
		clientError = client.Connect(ctx)

		go func() {
			if clientError != nil {
				cancelFun()
				log.Fatal(clientError)

			}
		}()

		globals.MongoDBConnection = client.Database(_const.DBName)
		fmt.Println("connected to mongodb")

	}
}

func WriteInMongoDB(document *dto.MongoDocumentDto) {
	connectDB()
	ctx, cancelFun := context.WithTimeout(context.Background(), 10*time.Second)
	fullDataCollection := globals.MongoDBConnection.Collection(_const.FullDataCollection)
	mainParamsCollection := globals.MongoDBConnection.Collection(_const.MainCollectionName)

	_, fullDataCollectionErr := fullDataCollection.InsertOne(ctx, document)
	go func() {
		if fullDataCollectionErr != nil {
			cancelFun()
			log.Fatal("fullDataCollectionErr")
		}
	}()

	_, mainParamsCollectionErr := mainParamsCollection.InsertOne(ctx, document.MainParams)

	if mainParamsCollectionErr != nil {
		log.Fatal("mainParamsCollectionErr")
	}
}

func ReadAllMainParams() []dto.MainParams {
	connectDB()
	var mainParamsList []dto.MainParams
	ctx, cancelFun := context.WithTimeout(context.Background(), 10*time.Second)

	mainParamsCollection := globals.MongoDBConnection.Collection(_const.MainCollectionName)
	mainCollectionCursor, mainCollectionCursorErr := mainParamsCollection.Find(ctx, bson.M{})
	go func() {
		if mainCollectionCursorErr != nil {
			cancelFun()
			log.Fatal("mainCollectionCursorErr")
		}
	}()

	for mainCollectionCursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem dto.MainParams
		err := mainCollectionCursor.Decode(&elem)
		fmt.Println("ID", elem.Id)
		if err != nil {
			log.Fatal(err)
		}
		mainParamsList = append(mainParamsList, elem)
	}
	return mainParamsList
}

func ReadFullDataDocument() []dto.MongoDocumentDto {
	connectDB()
	var fullDataDocument []dto.MongoDocumentDto
	ctx, cancelFun := context.WithTimeout(context.Background(), 10*time.Second)

	fullDataCollection := globals.MongoDBConnection.Collection(_const.FullDataCollection)
	fullDataCursor, fullDataCollectionErr := fullDataCollection.Find(ctx, bson.M{})
	go func() {
		if fullDataCollectionErr != nil {
			cancelFun()
			log.Fatal("fullDataCollectionErr")
		}
	}()

	for fullDataCursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem dto.MongoDocumentDto
		err := fullDataCursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		fullDataDocument = append(fullDataDocument, elem)
	}
	return fullDataDocument
}

func ReadFullFirstData() dto.MongoDocumentDto {
	connectDB()
	ctx, cancelFun := context.WithTimeout(context.Background(), 10*time.Second)

	fullDataCollection := globals.MongoDBConnection.Collection(_const.FullDataCollection)
	fullDataCursor, fullDataCollectionErr := fullDataCollection.Find(ctx, bson.M{})
	go func() {
		if fullDataCollectionErr != nil {
			cancelFun()
			log.Fatal("fullDataCollectionErr")
		}
	}()
	fullDataCursor.Next(context.TODO())
	var elem dto.MongoDocumentDto
	err := fullDataCursor.Decode(&elem)
	if err != nil {
		log.Fatal(err)
	}
	return elem

}

func FindAllDataById(id string) dto.MongoDocumentDto {
	connectDB()

	fullDataCollection := globals.MongoDBConnection.Collection(_const.FullDataCollection)
	objectId, objectIDFromHexError := primitive.ObjectIDFromHex(id)
	if objectIDFromHexError != nil {
		log.Fatal("objectIDFromHexError")
	}
	result := fullDataCollection.FindOne(context.Background(), bson.M{"_id": objectId})
	var fullData dto.MongoDocumentDto
	fmt.Println(result)
	err := result.Decode(&fullData)
	if err != nil {
		return dto.MongoDocumentDto{}
	}
	return fullData

}
