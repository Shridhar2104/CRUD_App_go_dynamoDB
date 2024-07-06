package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/entityresolution"
)

type Database struct{
	dynamodb *dynamodb.DynamoDB
	logMode bool
}

type Interface interface{

}
func NewAdapter() Interface{
	
}

func (db *Database) Health() bool{

	_, err:= db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil


}
func (db *Database) FindAll(){
	

}
func (db *Database) FindOne(condition map[string]interface{}, tableName string) (res *dynamodb.GetItemOutput, err error){
	conditionParsed, err:= dynamodbattribute.MarshalMap(condition)
	input:= &dynamodb.GetItemInput{
		TableName: aws.String("tableName"),
		Key:conditionParsed,

	}

	return db.connection.GetItem(input)


}
func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (res *dynamodb.PutItemOutput, err error){
	entityParsed, err:=dynamodbattribute.MarshalMap(entity)
	if err!=nil{
		return err, nil
	}

	input:=&dynamodb.PutItemInput{
		Item: entityParsed,
		TableName: aws.String("tableName"),
	}
	return db.connection.PutItemInput(input)
}

func (db *Database) Delete(condition map[string]interface{}, tableName string) (res *dynamodb.DeleteItemOutput, err error){
	conditionParsed, err:= dynamodbattribute.MarshalMap(condition)
	if err!=nil{
		return err, nil
	}
	input:= &dynamodb.DeleteItemInput{
		TableName: aws.String("tableName"),
		Key:conditionParsed,
	}
		return db.connection.DeleteItem(input)

	

}