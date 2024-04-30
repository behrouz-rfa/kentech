package utils

import (
	"testing"
	"time"

	"github.com/behrouz-rfa/kentech/pkg/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestToMap(t *testing.T) {
	t.Run("primitives", func(t *testing.T) {
		item := struct {
			SimpleValue string
		}{
			SimpleValue: "test",
		}

		result := ToMap(item)

		assert.Contains(t, result, "_id")
		assert.NotEmpty(t, result["_id"], "result['_id'] is empty")
		assert.Contains(t, result, "createdAt")
		assert.NotEmpty(t, result["createdAt"], "result['createdAt'] is empty")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
		assert.Contains(t, result, "simpleValue")
		assert.Equal(t, "test", result["simpleValue"])
	})

	t.Run("bson primitives", func(t *testing.T) {
		item := bson.M{"simpleValue": "test"}

		result := ToMap(item)

		assert.Contains(t, result, "simpleValue")
		assert.Equal(t, "test", result["simpleValue"])
	})

	t.Run("nested", func(t *testing.T) {
		item := struct {
			SimpleValue    string
			EmbeddedObject struct {
				EmbeddedValue string
			}
		}{
			SimpleValue: "test",
			EmbeddedObject: struct {
				EmbeddedValue string
			}{
				EmbeddedValue: "embedded",
			},
		}

		result := ToMap(item)

		assert.Contains(t, result, "_id")
		assert.NotEmpty(t, result["_id"], "result['_id'] is empty")
		assert.Contains(t, result, "createdAt")
		assert.NotEmpty(t, result["createdAt"], "result['createdAt'] is empty")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
		assert.Contains(t, result, "simpleValue")
		assert.Contains(t, result, "embeddedObject")
		assert.Equal(t, "test", result["simpleValue"])
		assert.IsType(t, map[string]interface{}{}, result["embeddedObject"])
		assert.Equal(t, "embedded", result["embeddedObject"].(map[string]interface{})["embeddedValue"])
	})

	t.Run("nested with empty pointer", func(t *testing.T) {
		item := struct {
			SimpleValue    string
			EmbeddedObject *struct {
				EmbeddedValue string
			}
		}{
			SimpleValue: "test",
		}

		result := ToMap(item)

		assert.Contains(t, result, "_id")
		assert.NotEmpty(t, result["_id"], "result['_id'] is empty")
		assert.Contains(t, result, "createdAt")
		assert.NotEmpty(t, result["createdAt"], "result['createdAt'] is empty")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
		assert.Contains(t, result, "simpleValue")
		assert.Equal(t, "test", result["simpleValue"])
	})

	t.Run("basic pointer", func(t *testing.T) {
		item := struct {
			SimpleValue  string
			PointerValue *string
		}{
			SimpleValue:  "test",
			PointerValue: utils.ToValue("pointer"),
		}

		result := ToMap(item)

		assert.Contains(t, result, "_id")
		assert.NotEmpty(t, result["_id"], "result['_id'] is empty")
		assert.Contains(t, result, "createdAt")
		assert.NotEmpty(t, result["createdAt"], "result['createdAt'] is empty")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
		assert.Contains(t, result, "simpleValue")
		assert.Equal(t, "test", result["simpleValue"])
		assert.Contains(t, result, "pointerValue")
		assert.Equal(t, utils.ToValue("pointer"), result["pointerValue"])
	})

	t.Run("nested with pointer", func(t *testing.T) {
		item := struct {
			SimpleValue    string
			EmbeddedObject *struct {
				EmbeddedValue string
			}
		}{
			SimpleValue: "test",
			EmbeddedObject: &struct {
				EmbeddedValue string
			}{
				EmbeddedValue: "embedded",
			},
		}

		result := ToMap(item)

		assert.Contains(t, result, "_id")
		assert.NotEmpty(t, result["_id"], "result['_id'] is empty")
		assert.Contains(t, result, "createdAt")
		assert.NotEmpty(t, result["createdAt"], "result['createdAt'] is empty")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
		assert.Contains(t, result, "simpleValue")
		assert.Contains(t, result, "embeddedObject")
		assert.Equal(t, "test", result["simpleValue"])
		assert.IsType(t, map[string]interface{}{}, result["embeddedObject"])
		assert.Equal(t, "embedded", result["embeddedObject"].(map[string]interface{})["embeddedValue"])
	})

	t.Run("time test", func(t *testing.T) {
		now := time.Now()
		item := struct {
			SimpleValue string
			TimeVal     time.Time
		}{
			SimpleValue: "test",
			TimeVal:     now,
		}

		result := ToMap(item)

		assert.Contains(t, result, "_id")
		assert.NotEmpty(t, result["_id"], "result['_id'] is empty")
		assert.Contains(t, result, "createdAt")
		assert.NotEmpty(t, result["createdAt"], "result['createdAt'] is empty")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
		assert.Contains(t, result, "simpleValue")
		assert.Equal(t, "test", result["simpleValue"])
		assert.Contains(t, result, "timeVal")
		assert.Equal(t, now, result["timeVal"])
	})

	t.Run("update", func(t *testing.T) {
		item := struct {
			SimpleValue string
		}{
			SimpleValue: "test",
		}

		result := ToMap(item, MethodUpdate)

		assert.Contains(t, result, "simpleValue")
		assert.Equal(t, "test", result["simpleValue"])
		assert.NotContains(t, result, "_id")
		assert.NotContains(t, result, "createdAt")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
	})

	t.Run("update nested with pointer", func(t *testing.T) {
		item := struct {
			SimpleValue    string
			EmbeddedObject *struct {
				EmbeddedValue string
			}
		}{
			SimpleValue: "test",
			EmbeddedObject: &struct {
				EmbeddedValue string
			}{
				EmbeddedValue: "embedded",
			},
		}

		result := ToMap(item, MethodUpdate)

		assert.NotContains(t, result, "_id")
		assert.NotContains(t, result, "createdAt")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
		assert.Contains(t, result, "simpleValue")
		assert.Contains(t, result, "embeddedObject.embeddedValue")
		assert.Equal(t, "test", result["simpleValue"])
		assert.IsType(t, "string", result["embeddedObject.embeddedValue"])
		assert.Equal(t, "embedded", result["embeddedObject.embeddedValue"])
	})

	t.Run("update nested with zero pointer", func(t *testing.T) {
		item := struct {
			SimpleValue    string
			EmbeddedObject *struct {
				EmbeddedValue string
			}
		}{
			SimpleValue:    "test",
			EmbeddedObject: nil,
		}

		result := ToMap(item, MethodUpdate)

		assert.NotContains(t, result, "_id")
		assert.NotContains(t, result, "createdAt")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
		assert.Contains(t, result, "simpleValue")
		assert.NotContains(t, result, "embeddedObject")
	})

	t.Run("update basic pointer", func(t *testing.T) {
		item := struct {
			SimpleValue  string
			PointerValue *string
		}{
			SimpleValue:  "test",
			PointerValue: utils.ToValue("pointer"),
		}

		result := ToMap(item, MethodUpdate)

		assert.NotContains(t, result, "_id")
		assert.NotContains(t, result, "createdAt")
		assert.Contains(t, result, "updatedAt")
		assert.NotEmpty(t, result["updatedAt"], "result['updatedAt'] is empty")
		assert.Contains(t, result, "simpleValue")
		assert.Equal(t, "test", result["simpleValue"])
		assert.Contains(t, result, "pointerValue")
		assert.Equal(t, utils.ToValue("pointer"), result["pointerValue"])
	})

	t.Run("filter map", func(t *testing.T) {
		item := struct {
			ID          *string
			SimpleValue *string
			EmptyValue  *string
		}{
			ID:          utils.ToValue("id"),
			SimpleValue: utils.ToValue("simple"),
			EmptyValue:  nil,
		}

		result := ToMap(item, MethodFilter)

		assert.Contains(t, result, "_id")
		assert.Equal(t, utils.ToValue("id"), result["_id"])
		assert.NotContains(t, result, "id")
		assert.Contains(t, result, "simpleValue")
		assert.Equal(t, utils.ToValue("simple"), result["simpleValue"])
		assert.NotContains(t, result, "emptyValue")
	})
}
