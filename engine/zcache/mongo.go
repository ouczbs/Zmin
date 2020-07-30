package zcache

import (
	"Zmin/engine/zlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *UMongoClient
type UMongoClient struct {
	*mongo.Database
}
func init()  {
	MongoClient = NewMongoClient("mongodb://111.229.54.9:27017" , "mmo")
}
func NewMongoClient(listenAddr string, dbname string)* UMongoClient{
	client, err := mongo.Connect(ctx , options.Client().ApplyURI(listenAddr))
	if err != nil {
		return nil
	}
	return &UMongoClient{
		Database:client.Database(dbname),
	}
}
func (mongo UMongoClient)UpdateOrInsert(model IModel , query interface{})error{
	var upsert = true
	_ , err := mongo.Collection(model.Table()).UpdateOne(ctx , query, bson.M{
		"$set":model.M(),
	} , &options.UpdateOptions{Upsert: &upsert})
	return err
}
func (mongo * UMongoClient)InsertOne(model IModel)error{
	_ , err := mongo.Collection(model.Table()).InsertOne(ctx , model.M())
	return err
}
func (mongo * UMongoClient)ClearTable(model IModel)error{
	opts := options.Delete().SetCollation(&options.Collation{     Locale:    "en_US",     Strength:  1,     CaseLevel: false})
	res , err := mongo.Collection(model.Table()).DeleteMany(ctx , bson.D{},opts)
	zlog.Debug(res.DeletedCount)
	return err
}
func (mongo * UMongoClient)Find(model IModel , query interface{} ,results interface{})error{
	res , err := mongo.Collection(model.Table()).Find(ctx , query)
	if err != nil{
		return err
	}
	return res.All(ctx , results)
}
func (mongo * UMongoClient)FindOne(model IModel , query interface{} ,result interface{})error{
	err := mongo.Collection(model.Table()).FindOne(ctx , query).Decode(result)
	return err
}