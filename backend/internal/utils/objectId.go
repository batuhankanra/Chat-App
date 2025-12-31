package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ToObjectId(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

func ToObjectIdSlice(ids []string) ([]primitive.ObjectID, error) {
	objectIds := make([]primitive.ObjectID, 0, len(ids))
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIds = append(objectIds, objId)
	}
	return objectIds, nil
}
