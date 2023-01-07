package Database

import (
"context"
"fmt"
"go.mongodb.org/mongo-driver/bson"
"log"
"time"

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)




	//updatePass(ctx, client, "AaronB", "cheese")

	//updateRole(ctx, client, "AaronB", "Senior Golang Engineer")

	//createRecord("AaronB", "pass2", "Golang Engineer", ctx, client)

	//deleteRecord(ctx, client, "AaronX")



}

/*func getRole(ctx context.Context, CL *mongo.Client, u string){
	ourDatabase = CL.Database("ourDatabase")
	usersCollection = ourDatabase.Collection("users")
	result := usersCollection.FindOne(
		ctx,
		bson.M{"user_name": u},
			)

	fmt.Print()
}

*/


// The updatePass function updates the password, taking the context, client, username and new password as parameters
func updatePass(ctx context.Context, CL *mongo.Client, u string,p string) {
	ourDatabase := CL.Database("ourDatabase")
	usersCollection := ourDatabase.Collection("users")
	result, err :=

		usersCollection.UpdateOne(
			ctx,
			bson.M{"user_name": u},
			bson.M{	"$set": bson.M{"password": p}},
		)
	if err != nil {
		fmt.Print(err)
	} else {

		fmt.Println("Updated ", result.ModifiedCount, " records")
	}
}

// The updatePass function updates the password, taking the context, client, username and new password as parameters
func updateRole(ctx context.Context, CL *mongo.Client, u string,r string) {
	ourDatabase := CL.Database("ourDatabase")
	usersCollection := ourDatabase.Collection("users")
	result, err :=

		usersCollection.UpdateOne(
			ctx,
			bson.M{"user_name": u},
			bson.M{	"$set": bson.M{"Role": r}},
		)
	if err != nil {
		fmt.Print(err)
	} else {

		fmt.Println("Updated ", result.ModifiedCount, " records")
	}
}

// This function will delete a user from the database
func deleteRecord(ctx context.Context, CL *mongo.Client, u string){
	ourDatabase := CL.Database("ourDatabase")
	usersCollection := ourDatabase.Collection("users")
	result, err := usersCollection.DeleteOne(ctx, bson.M{"user_name": u})
	if err != nil {
		fmt.Println("There is an error")
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)", result.DeletedCount)
}




// createRecord adds a record to our database
func createRecord(u string, p string, r string, c context.Context, CL *mongo.Client){
	fmt.Println("adding record")
	ourDatabase := CL.Database("ourDatabase")
	usersCollection := ourDatabase.Collection("users")
	userResult, err := usersCollection.InsertOne(c, bson.D{
		{Key: "user_name", Value: u},
		{Key: "password", Value: p},
		{Key: "Role", Value: r},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userResult.InsertedID)

}
